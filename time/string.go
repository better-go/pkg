package time

// 月份+偏移量(支持 +, - 偏移):  [格式: 2019-08]
func MonthWithOffset(offset int) string {
	return BeginningOfMonthWithMonthOffset(offset).Format(fmtMonth)
}

// 今天+偏移量(支持 +, - 偏移):  [格式: 2019-08-23]
func TodayWithOffset(offset int) string {
	return BeginningOfDayWithDayOffset(offset).Format(fmtDay)
}

/////////////////////////////////////////////////////////////////////////////

// 本月:
func CurrentMonth() string {
	return MonthWithOffset(0)
}

// 上个月:
func LastMonth() string {
	return MonthWithOffset(-1)
}

// 下个月:
func NextMonth() string {
	return MonthWithOffset(1)
}

/////////////////////////////////////////////////////////////////////////////

// 今日:
func Today() string {
	return TodayWithOffset(0)
}

// 昨日:
func Yesterday() string {
	return TodayWithOffset(-1)
}

// 明日:
func Tomorrow() string {
	return TodayWithOffset(1)
}
