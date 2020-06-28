package time

import (
	"fmt"
	"testing"
	"time"
)

func TestGenTimestamp10bit(t *testing.T) {
	in := []string{
		Gen10BitTimestamp(),
		Gen13BitTimestamp(),
		Gen19BitTimestamp(),
	}

	for i, item := range in {
		t.Log(i, "ts:", len(item), item)
	}
}

func TestGenTs10bit(t *testing.T) {
	in := []int64{
		Gen10BitTs(),
		Gen13BitTs(),
		Gen19BitTs(),
	}

	for i, item := range in {
		t.Log(i, "ts:", item)
	}

	ts := fmt.Sprintf("%d", time.Now().Unix()*1000)
	t.Log("ts=", ts)

}

func TestDuration13Bit(t *testing.T) {
	in := []interface{}{
		DurationWithSecond(0, 1, 0),
		Duration10Bit(0, 1, 0),
		DurationWithMillisecond(0, 1, 0),
		Duration13Bit(0, 1, 0),
	}

	for i, item := range in {
		t.Log(i, "duration:", item)
	}
}

func TestIsTimestampValid(t *testing.T) {
	in := []string{
		// 5s 之前
		fmt.Sprintf("%d", Gen13BitTs()-int64(Duration13Bit(0, 0, 5))),
		// 10s 之前
		fmt.Sprintf("%d", Gen13BitTs()-int64(Duration13Bit(0, 0, 10))),
	}

	t.Log("now:", Gen10BitTs())

	for i, item := range in {
		ok := Is13BitTimestampValid(item, 0, 0, 20) // 20s 有效
		t.Log(i, "ok:", item, ok)
		ok2 := Is13BitTimestampValid(item, 0, 0, 6) // 6s 有效
		t.Log(i, "ok2:", item, ok2)
	}

}
