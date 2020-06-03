package time

import (
	"math"
	"testing"
)

func TestMonthDurationWithMonthOffset(t *testing.T) {
	in := []int{
		-1,
		-2,
		-3,
		-4,
		-5,
		-6,
		-12,
		-13,
		-14,
		-15,
		-25,
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		13,
		14,
		15,
		25,
	}

	for i, item := range in {
		when := ""
		if item >= 0 {
			when = "later"
		} else {
			when = "ago"
		}
		start, end := MonthDurationWithMonthOffset(item)

		t.Logf("No%2d, %4v month %6s: first = %v, end = %v\n",
			i,
			math.Abs(float64(item)),
			when,
			start, // start
			end,   // end
		)
	}
}

func TestDayDurationWithDayOffset(t *testing.T) {
	in := []int{
		-1,
		-2,
		-3,
		-20,
		-25,
		-50,
		-60,
		-70,
		0,
		1,
		2,
		3,
		28,
		31,
		55,
		60,
		80,
		120,
		240,
		366,
	}

	for i, item := range in {
		when := ""
		if item >= 0 {
			when = "later"
		} else {
			when = "ago"
		}
		start, end := DayDurationWithDayOffset(item)

		t.Logf("No%2d, %4v day %6s: first = %v, end = %v\n",
			i,
			math.Abs(float64(item)),
			when,
			start, // start
			end,   // end
		)
	}
}

func TestMonthAgoDuration(t *testing.T) {
	agoStart, agoEnd := MonthAgoDuration()
	t.Logf("1 month ago:   start=%v, end=%v\n", agoStart, agoEnd)
	t.Logf("1 month ago:   start=%v, end=%v\n", agoStart.String(), agoEnd.String())

	laterStart, laterEnd := MonthLaterDuration()
	t.Logf("1 month later: start=%v, end=%v\n", laterStart, laterEnd)
}

func TestMonthAgoDurationTs(t *testing.T) {
	agoStart, agoEnd := MonthAgoDurationTs()
	t.Logf("1 month ago ts:   start=%v, end=%v\n", agoStart, agoEnd)
}
