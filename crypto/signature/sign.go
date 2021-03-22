package signature

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"net/url"
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
	SignMD5    = "md5"    // md5
	SignSHA256 = "sha256" // sha256
	SignSHA512 = "sha512" // sha512
)

type (
	Signer struct {
		lock sync.RWMutex      // 锁
		keys map[string]string // cached key map // TODO: 主动过期, 避免 key pair 泄露问题

		// sign common fields:
		publicKeyName  string
		privateKeyName string
		nonceName      string
		timestampName  string
		signTypeName   string
		signResultName string
	}

	// get appSecret/privateKey by appKey/publicKey
	PrivateKeyFunc func(publicKey string) (privateKey string, err error)

	// 签名算法: useMD5/sha256/sha512
	SignAlgorithmFunc func(data string, privateKey string) (digest string)
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
	payload url.Values,
	publicKey string,
	nonce string,
	timestamp string,
	toLower bool,
	keyFn PrivateKeyFunc, // 私钥获取方法
	signFn SignAlgorithmFunc, // 签名算法
) (sign string) {
	// pack args:
	pack := m.pack(payload, publicKey, nonce, timestamp, SignMD5)

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
	data url.Values,
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
	newSign := m.Sign(data, publicKey, nonce, timestamp, toLower, keyFn, signFn)
	log.Debugf("verify: old sign=%v, new sign=%v", sign, newSign)
	return sign == newSign
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *Signer) SignMD5(payload url.Values,
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
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return m.WithMD5(data, privateKey)
		},
	)
}

func (m *Signer) VerifyMD5(
	data url.Values,
	toLower bool,
	keyFn PrivateKeyFunc,
) bool {
	return m.Verify(
		data,
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return m.WithMD5(data, privateKey)
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// sign with SHA256:
func (m *Signer) SignSHA256(payload url.Values,
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
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return m.WithSHA256(data, privateKey)
		},
	)
}

func (m *Signer) VerifySHA256(
	data url.Values,
	toLower bool,
	keyFn PrivateKeyFunc,
) bool {
	return m.Verify(
		data,
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return m.WithSHA256(data, privateKey)
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// sign with SHA512:
func (m *Signer) SignSHA512(payload url.Values,
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
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return m.WithSHA512(data, privateKey)
		},
	)
}

func (m *Signer) VerifySHA512(
	data url.Values,
	toLower bool,
	keyFn PrivateKeyFunc,
) bool {
	return m.Verify(
		data,
		toLower,
		keyFn,
		func(data string, privateKey string) (digest string) {
			return m.WithSHA512(data, privateKey)
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *Signer) WithMD5(data string, privateKey string) string {
	// do sign:
	digest := md5.Sum([]byte(data + privateKey))
	return hex.EncodeToString(digest[:])
}

func (m *Signer) WithSHA256(data string, privateKey string) string {
	input := []byte(data + privateKey)

	h := sha256.New()
	h.Write(input)

	// do sign:
	digest := h.Sum(nil)
	return hex.EncodeToString(digest[:])
}

func (m *Signer) WithSHA512(data string, privateKey string) string {
	input := []byte(data + privateKey)

	h := sha512.New()
	h.Write(input)

	// do sign:
	digest := h.Sum(nil)
	return hex.EncodeToString(digest[:])
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
func (m *Signer) pack(payload url.Values, publicKey string, nonce string, timestamp string, signType string) url.Values {
	if publicKey != "" {
		payload.Set(m.publicKeyName, publicKey)
	}
	if nonce != "" {
		payload.Set(m.nonceName, nonce)
	}
	if timestamp != "" {
		payload.Set(m.timestampName, timestamp)
	}
	// 签名算法类型:
	if signType != "" {
		payload.Set(m.signTypeName, signType)
	}

	// clean:
	payload.Del(m.signResultName)
	return payload
}
