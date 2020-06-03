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
	id := Gen36BitUUID4()
	return strings.ReplaceAll(id, "-", "")
}

// 20bit: 推荐: 安全, 无冲突(详见 Benchmark test)
func Gen20BitUUID() string {
	return strings.ToUpper(Gen20BitUUIDLower())
}

// 默认全小写:
func Gen20BitUUIDLower() string {
	guid := xid.New()
	return guid.String()
}

//////////////////////////////////////////////////////

// Random (Version 4) UUID.
func GenUUID4() uuid.UUID {
	return uuid.New()
}

// 此方法不要用, 冲突率非常高, 参看 Benchmark test
func Gen20BitDigit() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%05d", rnd.Int63n(99999)))
	b.WriteString(fmt.Sprintf("%03d", time.Now().UnixNano()/1e6%1000))
	b.WriteString(time.Now().Format("060102150405"))
	return b.String()
}
