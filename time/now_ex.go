package time

import (
	"time"

	"github.com/jinzhu/now"
)

/*

说明:
	1. 对 github.com/jinzhu/now 时间包的扩展.
	2. 扩展原因: 原包不支持对月度等时间的偏移量计算. (只支持当前月时间计算)
	3. 支持偏移量计算: 对于跨年的月份, 要小心. 年和月的处理. (取模计算)

*/

const (
	monthUnit = 12 // 单位: 1 year = 12 month
	dayUnit   = 24 // 单位: 1 day = 24hour
)

// 扩展:
type NowEx struct {
	*now.Now
}

// New initialize Now with time
func New(t time.Time) *NowEx {
	return &NowEx{
		now.New(t),
	}
}

/////////////////////////////////////////////////////////////////////////////

//月初时刻: 当前月+偏移量(支持 +, - 偏移)
func (now *NowEx) BeginningOfMonthWithMonthOffset(offset int) time.Time {
	// 当前年+当前月
	y, m, _ := now.Date()

	// 月偏移量: 取模
	diffMonth := offset % monthUnit
	// 年偏移量: 取整
	diffYear := offset / monthUnit

	// 最终年 + 最终月:
	distYear, distMonth := y+diffYear, m+time.Month(diffMonth)
	return time.Date(distYear, distMonth, 1, 0, 0, 0, 0, now.Location())
}

// 月末时刻: 当前月+偏移量(支持 +, - 偏移)
func (now *NowEx) EndOfMonthWithMonthOffset(offset int) time.Time {
	return now.BeginningOfMonthWithMonthOffset(offset).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// 日初时刻: 当天+偏移量(支持 +, - 偏移)
func (now *NowEx) BeginningOfDayWithDayOffset(offset int) time.Time {
	// 当前年+当前月+当前日
	y, m, d := now.Date()
	return time.Date(y, m, d+offset, 0, 0, 0, 0, now.Time.Location())
}

// 日末时刻: 当天+偏移量(支持 +, - 偏移)
func (now *NowEx) EndOfDayWithDayOffset(offset int) time.Time {
	// 当前年+当前月+当前日
	y, m, d := now.Date()
	return time.Date(y, m, d+offset, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

/////////////////////////////////////////////////////////////////////////////

// 月初时刻: 上个月
func (now *NowEx) BeginningOfMonthAgo() time.Time {
	return now.BeginningOfMonthWithMonthOffset(-1)
}

// 月末时刻: 上个月
func (now *NowEx) EndOfMonthAgo() time.Time {
	return now.EndOfMonthWithMonthOffset(-1)
}

/////////////////////////////////////////////////////////////////////////////

// 月初时刻: 下个月
func (now *NowEx) BeginningOfMonthLater() time.Time {
	return now.BeginningOfMonthWithMonthOffset(1)
}

// 月末时刻: 下个月
func (now *NowEx) EndOfMonthLater() time.Time {
	return now.EndOfMonthWithMonthOffset(1)
}
