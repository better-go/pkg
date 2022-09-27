package random

import (
	"fmt"
	"testing"
)

func TestGenUUID4(t *testing.T) {
	t.Log(GenUUID4())

	// string:
	in := []string{
		Gen36BitUUID4(),
		Gen32BitUUID4(),
		Gen20BitUUID(),
		Gen20BitDigit(),
		GenDigit(4),
		GenDigit(6),
		GenDigit(20),
		GenDigit(21),
	}

	for i, item := range in {
		t.Log(i, "uuid:", len(item), item)
	}

}

var (
	set = map[string]struct{}{} // 冲突检测
)

// 冲突检测:
func conflictDetect(id string) {
	_, ok := set[id]
	if !ok {
		set[id] = struct{}{}
	} else {
		fmt.Printf("id conflict: %v, count=%v\n", id, len(set))
	}
	return
}

// 此方法不要用, 冲突太高
func BenchmarkGen20BitDigit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conflictDetect(Gen20BitDigit())
	}
}

// 推荐: 安全, 无冲突
func BenchmarkGen20BitUUID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uid := Gen20BitUUID()
		fmt.Println("gen 20 bit uuid: ", uid)
		conflictDetect(uid)
	}
}

// 推荐: 安全, 无冲突
func BenchmarkGen32BitUUID4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uid := Gen32BitUUID4()
		fmt.Println("gen 32 bit uuid: ", uid)
		conflictDetect(uid)
	}
}

func TestGen20BitUUIDUpper(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("uuid:", Gen20BitUUID())
	}
}

func BenchmarkGenUIDPair(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uid20, uid32 := GenUIDPair()
		fmt.Println("gen uid pair: ", uid20, uid32)
		conflictDetect(uid20)
		conflictDetect(uid32)
	}
}
