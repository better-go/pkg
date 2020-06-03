package time

import (
	"strconv"
	"time"
)

// 10bit:
func Gen10BitTimestamp() string {
	ts := Gen10BitTs()
	return strconv.FormatInt(ts, 10)
}

// 19bit:
func Gen19BitTimestamp() string {
	ts := Gen19BitTs()
	return strconv.FormatInt(ts, 10)
}

// 13bit: (19-6)
func Gen13BitTimestamp() string {
	ts := Gen13BitTs()
	return strconv.FormatInt(ts, 10)
}

//////////////////////////////////////////////////////

// 10bit: 单位: 秒 s
func Gen10BitTs() int64 {
	return time.Now().Unix()
}

// 19bit: 单位: 纳秒(十亿分之一秒)
func Gen19BitTs() int64 {
	return time.Now().UnixNano()
}

// 13bit: (19-6) 单位: 毫秒
func Gen13BitTs() int64 {
	return Gen19BitTs() / 1000000
}

// 时间戳是否过期:
func Is13BitTimestampValid(ts string, hour int64, minute int64, second int64) bool {
	inTs, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return false
	}

	// ok:
	return Is13BitTsValid(inTs, hour, minute, second)
}

// 时间戳是否过期:
func Is13BitTsValid(ts int64, hour int64, minute int64, second int64) bool {
	// invalid:
	if ts < 0 {
		return false
	}

	// diff:
	now := Gen13BitTs()
	duration := Duration13Bit(hour, minute, second)
	alive := float64(ts) + duration

	diff := alive - float64(now)
	// 已经过期:
	if diff < 0 {
		return false
	}

	// 未过期:
	return true
}

////////////////////////////////////////////////////////////////////

// DurationWithSecond 别名
func Duration10Bit(hour int64, minute int64, second int64) float64 {
	return DurationWithSecond(hour, minute, second)
}

// DurationWithSecond 别名
func Duration13Bit(hour int64, minute int64, second int64) float64 {
	return DurationWithMillisecond(hour, minute, second)
}

////////////////////////////////////////////////////////////////////
// 时间段: 单位 秒(s)
func DurationWithSecond(hour int64, minute int64, second int64) float64 {
	totalSecond := float64(0)

	if hour > 0 {
		totalSecond += time.Hour.Seconds() * float64(hour)
	}
	if minute > 0 {
		totalSecond += time.Minute.Seconds() * float64(minute)
	}
	if second > 0 {
		totalSecond += time.Second.Seconds() * float64(second)
	}
	return totalSecond
}

// 时间段: 单位 毫秒 (千分之一秒)
func DurationWithMillisecond(hour int64, minute int64, second int64) float64 {
	return DurationWithSecond(hour, minute, second) * 1000
}
