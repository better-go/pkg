package time

import (
	"time"
)

/*

说明:
	1. 对 github.com/jinzhu/now 时间包的扩展.
	2. 扩展原因: 原包不支持对月度等时间的偏移量计算. (只支持当前月时间计算)
	3. 支持偏移量计算: 对于跨年的月份, 要小心. 年和月的处理. (取模计算)

使用方式:
	1. 直接调用如下方法.
	2. 单元测试包含使用示例.

*/

// 月初时刻: 当前月+月偏移量(支持 +, - 偏移)
func BeginningOfMonthWithMonthOffset(offset int) time.Time {
	return New(time.Now()).BeginningOfMonthWithMonthOffset(offset)
}

// 月末时刻: 当前月+月偏移量(支持 +, - 偏移)
func EndOfMonthWithMonthOffset(offset int) time.Time {
	return New(time.Now()).EndOfMonthWithMonthOffset(offset)
}

// 日初时刻: 当天+天偏移量(支持 +, - 偏移)
func BeginningOfDayWithDayOffset(offset int) time.Time {
	return New(time.Now()).BeginningOfDayWithDayOffset(offset)
}

// 日末时刻: 当天+天偏移量(支持 +, - 偏移)
func EndOfDayWithDayOffset(offset int) time.Time {
	return New(time.Now()).EndOfDayWithDayOffset(offset)
}

/////////////////////////////////////////////////////////////////////////////

// 月初时刻: 上个月
func BeginningOfMonthAgo() time.Time {
	return New(time.Now()).BeginningOfMonthAgo()
}

// 月末时刻: 上个月
func EndOfMonthAgo() time.Time {
	return New(time.Now()).EndOfMonthAgo()
}

/////////////////////////////////////////////////////////////////////////////

// 月初时刻: 上个月
func BeginningOfMonthLater() time.Time {
	return New(time.Now()).BeginningOfMonthLater()
}

// 月末时刻: 上个月
func EndOfMonthLater() time.Time {
	return New(time.Now()).EndOfMonthLater()
}

/////////////////////////////////////////////////////////////////////////////

// 基于 <月初1号> 偏移 N 天(+, -)
func BeginOfMonthWithDayOffset(monthOffset int, dayOffset int) time.Time {
	monthStart := BeginningOfMonthWithMonthOffset(monthOffset)

	// N day
	ts := time.Duration(dayOffset) * 24 * time.Hour
	day := monthStart.Add(ts)
	return day
}
