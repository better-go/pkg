package random

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

var (
	rnd *rand.Rand
)

func init() {
	rnd = rand.New(rand.NewSource(time.Now().Unix()))
}

// 36bit:
func Gen36BitUUID4() string {
	id := uuid.New()
	return id.String()
}

// 32bit: 推荐: 安全, 无冲突(详见 Benchmark test)
func Gen32BitUUID4() string {
	raw := Gen36BitUUID4()
	uid := strings.ReplaceAll(raw, "-", "")
	return strings.ToUpper(uid)
}

// 20bit: 推荐: 安全, 无冲突(详见 Benchmark test)
func Gen20BitUUID() string {
	return strings.ToUpper(Gen20BitUUIDLower())
}

// Gen20BitUUIDLower 默认全小写:
func Gen20BitUUIDLower() string {
	guid := xid.New()
	return guid.String()
}

// GenUIDPair 生成一对ID(publicKey, secretKey)
func GenUIDPair() (string, string) {
	return Gen20BitUUID(), Gen32BitUUID4()
}

//////////////////////////////////////////////////////

// GenUUID4 Random (Version 4) UUID.
func GenUUID4() uuid.UUID {
	return uuid.New()
}

// Gen20BitDigit TODO X: 此方法不要用, 冲突率非常高, 参看 Benchmark test
func Gen20BitDigit() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%05d", rnd.Int63n(99999)))
	b.WriteString(fmt.Sprintf("%03d", time.Now().UnixNano()/1e6%1000))
	b.WriteString(time.Now().Format("060102150405"))
	return b.String()
}

// GenDigit for verify code <= 6, 适合生成的短信验证码场景
func GenDigit(length int32) string {
	// invalid:
	if length > 20 {
		return ""
	}

	// max 20 bit:
	code := Gen20BitDigit()

	// slice:
	return code[:length]
}
