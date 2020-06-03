package time

import (
	"time"
)

// 月份区间: 偏移量(支持 +, - 偏移)
func MonthDurationWithMonthOffset(offset int) (time.Time, time.Time) {
	start := BeginningOfMonthWithMonthOffset(offset)
	end := EndOfMonthWithMonthOffset(offset)
	return start, end
}

// 日区间: 偏移量(支持 +, - 偏移)
func DayDurationWithDayOffset(offset int) (time.Time, time.Time) {
	start := BeginningOfDayWithDayOffset(offset)
	end := EndOfDayWithDayOffset(offset)
	return start, end
}

/////////////////////////////////////////////////////////////////////////////

// 月份区间: 上个月
func MonthAgoDuration() (time.Time, time.Time) {
	return MonthDurationWithMonthOffset(-1)
}

// 月份区间: 下个月
func MonthLaterDuration() (time.Time, time.Time) {
	return MonthDurationWithMonthOffset(1)
}

/////////////////////////////////////////////////////////////////////////////

// 月份区间时间戳: 上个月
func MonthAgoDurationTs() (int64, int64) {
	start, end := MonthAgoDuration()
	startTs, endTs := start.Unix(), end.Unix()
	return startTs, endTs
}
