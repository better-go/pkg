package signature

import (
	"strings"
	"sync"

	"github.com/better-go/pkg/log"
)

/*
API 参数增加签名验证机制:
	- 应用场景:
		- http/rpc api: 参数签名, 防止中间人攻击
			- 区别 token auth, token auth 解决用户登录认证
			- sign: 解决 合法数据被劫持/伪造/重放(中间人攻击)
		- mq 队列消息: 参数签名, 防止伪造消息
	- 字段:
		- data: 原始数据 map
		- timestamp: 过期保护, 防止重放攻击(目前未做 timeout 保护)
		- nonce: 防止重复
		- sign: 签名结果
		- sign_type: 签名算法类型
		- app_key/public_key: 公钥
		- app_secret/private_key: 私钥

	- 签名规则:
		- 待签名数据部分: data 字段 + 拼接 nonce + timestamp + public_key + sign_type, 一起排序
			- 示例: address=shanghai&age=22&key1=a1&key2=b1&name=kate&nonce=705935809601077250&public_key=this-is-public-key&sex=girl&sign_type=useMD5&ts=1616053176
			- sign 值: 95d483f5eae66707965dfecf312c6cdf
			- privateKey: "use-this-do-sign"
		- 签名出 sign 值( sign 值 签名部分, 不含 app_key 和 app_secret)

	- 签名算法:
		- md5
		- sha256
	- signer / verifier
*/

const (
	// 签名公共字段:
	SignPublicKeyName  = "public_key"  // 字段名: 公钥
	SignPrivateKeyName = "private_key" // 字段名: 私钥
	SignNonceName      = "nonce"       // 字段名: 随机串
	SignTimestampName  = "ts"          // 字段名: 时间戳
	SignTypeName       = "sign_type"   // 字段名: 签名算法类型
	SignResultName     = "sign"        // 字段名: 签名结果

	// 签名算法类型:
	SignTypeMD5    = "md5"    // md5
	SignTypeSHA256 = "sha256" // sha256
	SignTypeSHA512 = "sha512" // sha512
)

type (
	Signer struct {
		lock sync.RWMutex      // 锁
		keys map[string]string // cached key map // TODO: 主动过期, 避免 key pair 泄露问题

		// sign common fields:
		publicKeyName  string // 字段名: 私钥 默认 SignPublicKeyName
		privateKeyName string // 字段名: 私钥
		nonceName      string
		timestampName  string
		signTypeName   string
		signResultName string
	}

	// get appSecret/privateKey by appKey/publicKey
	PrivateKeyFunc func(publicKey string) (key string, err error)

	// 签名算法: useMD5/sha256/sha512
	SignAlgorithmFunc func(data string, privateKey string) (digest string)

	// interface:
	DictType interface {
		Encode() string
		Get(key string) interface{}
		Set(key string, value interface{})
		Del(key string)
	}
)

func New(publicKeyName string, nonceName string, timestampName string) *Signer {
	// use default name:
	s := &Signer{
		keys:           make(map[string]string),
		publicKeyName:  SignPublicKeyName,
		privateKeyName: SignPrivateKeyName,
		nonceName:      SignNonceName,
		timestampName:  SignTimestampName,
		signTypeName:   SignTypeName,
		signResultName: SignResultName,
	}

	// change:
	if publicKeyName != "" {
		s.publicKeyName = publicKeyName
	}
	if nonceName != "" {
		s.nonceName = nonceName
	}
	if timestampName != "" {
		s.timestampName = timestampName
	}
	return s
}

// 签名生成:
func (m *Signer) Sign(
	payload DictType, // 数据 body
	publicKey string, // 公钥
	nonce string, // 随机串
	timestamp string, // 时间戳
	signType string, // 签名算法
	toLower bool, // 是否转换小写
	keyFn PrivateKeyFunc, // 私钥获取方法
	signFn SignAlgorithmFunc, // 签名算法
) (sign string) {
	// pack args:
	pack := m.pack(payload, publicKey, nonce, timestamp, signType)

	// do convert:
	data := pack.Encode()
	if strings.IndexByte(data, '+') > -1 {
		data = strings.Replace(data, "+", "%20", -1)
	}

	// 转换全小写:
	if toLower {
		data = strings.ToLower(data)
	}

	log.Debugf("sign encode: %v", data)

	// get private key:
	privateKey := m.privateKey(publicKey, keyFn)

	// do sign:
	return signFn(data, privateKey)
}

// 签名验证:
func (m *Signer) Verify(
	data DictType,
	signType string,
	toLower bool,
	keyFn PrivateKeyFunc,
	signFn SignAlgorithmFunc,
) bool {
	log.Debugf("verify: data=%+v", data)
	// raw data sign value:
	sign := data.Get(m.signResultName)
	//data.Del(m.signResultName)
	// after all done, reset it
	//defer data.Set(m.signResultName, sign)

	// common fields:
	publicKey := data.Get(m.publicKeyName)
	nonce := data.Get(m.nonceName)
	timestamp := data.Get(m.timestampName)

	// new sign:
	newSign := m.Sign(
		data,
		publicKey.(string), // 类型转换
		nonce.(string),
		timestamp.(string),
		signType,
		toLower,
		keyFn,
		signFn,
	)

	log.Debugf("verify: old sign=%v, new sign=%v", sign, newSign)
	return sign == newSign
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *Signer) SignMD5(
	payload DictType,
	publicKey string,
	nonce string,
	timestamp string,
	toLower bool,
	keyFn PrivateKeyFunc, // 私钥获取方法
) string {
	return m.Sign(
		payload,
		publicKey,
		nonce,
		timestamp,
		SignTypeMD5,
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return WithMD5(data, privateKey)
		},
	)
}

func (m *Signer) VerifyMD5(
	data DictType,
	toLower bool,
	keyFn PrivateKeyFunc,
) bool {
	return m.Verify(
		data,
		SignTypeMD5, // md5
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return WithMD5(data, privateKey)
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// sign with SHA256:
func (m *Signer) SignSHA256(
	payload DictType, // 数据
	publicKey string, // 公钥
	nonce string, // 随机串
	timestamp string, // 时间戳
	toLower bool, // 是否转小写
	keyFn PrivateKeyFunc, // 私钥获取方法
) string {
	return m.Sign(
		payload,
		publicKey,
		nonce,
		timestamp,
		SignTypeSHA256, // sha256
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return WithSHA256(data, privateKey)
		},
	)
}

func (m *Signer) VerifySHA256(
	data DictType,
	toLower bool,
	keyFn PrivateKeyFunc,
) bool {
	return m.Verify(
		data,
		SignTypeSHA256, // sha256
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return WithSHA256(data, privateKey)
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// sign with SHA512:
func (m *Signer) SignSHA512(payload DictType,
	publicKey string,
	nonce string,
	timestamp string,
	toLower bool,
	keyFn PrivateKeyFunc, // 私钥获取方法
) string {
	return m.Sign(
		payload,
		publicKey,
		nonce,
		timestamp,
		SignTypeSHA512, // sha512
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return WithSHA512(data, privateKey)
		},
	)
}

func (m *Signer) VerifySHA512(
	data DictType,
	toLower bool,
	keyFn PrivateKeyFunc,
) bool {
	return m.Verify(
		data,
		SignTypeSHA512, // sha512
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return WithSHA512(data, privateKey)
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// get privateKey:
func (m *Signer) privateKey(publicKey string, keyFn PrivateKeyFunc) (privateKey string) {
	// read:
	m.lock.RLock()
	cached, ok := m.keys[publicKey]
	m.lock.RUnlock()

	if ok {
		return cached
	}

	// not match:
	fetched, err := keyFn(publicKey)
	if err != nil {
		log.Error("query cached error, publicKey=%v, err=%v", publicKey, err)
		return ""
	}
	m.lock.Lock()
	m.keys[publicKey] = fetched
	m.lock.Unlock()

	return fetched
}

// 参数打包:
func (m *Signer) pack(payload DictType, publicKey string, nonce string, timestamp string, signType string) DictType {
	payload.Set(m.publicKeyName, publicKey)
	payload.Set(m.nonceName, nonce)
	payload.Set(m.timestampName, timestamp)
	// 签名算法类型:
	payload.Set(m.signTypeName, signType)

	//
	// safe guard: clean dirty sign
	//
	payload.Del(m.signResultName)
	return payload
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// use default field name:
func SignSHA256(
	payload DictType, // 数据
	publicKey string, // 公钥
	nonce string, // 随机串
	timestamp string, // 时间戳
	toLower bool, // 是否转小写
	keyFn PrivateKeyFunc, // 私钥获取方法

) string {
	signer := New("", "", "")
	return signer.SignSHA256(payload, publicKey, nonce, timestamp, toLower, keyFn)
}

// use default field name:
func VerifySHA256(
	data DictType, // 数据
	toLower bool, // 是否转小写
	keyFn PrivateKeyFunc, // 私钥获取方法
) bool {
	signer := New("", "", "")
	return signer.VerifySHA256(data, toLower, keyFn)
}
