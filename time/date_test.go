package time

import (
	"math"
	"testing"
	"time"
)

func TestBeginningOfMonthWithMonthOffset(t *testing.T) {
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
		t.Logf("No%2d, %4v month %6s: first = %v, end = %v\n",
			i,
			math.Abs(float64(item)),
			when,
			BeginningOfMonthWithMonthOffset(item), // start
			EndOfMonthWithMonthOffset(item),       // end
		)
	}
}

func TestBeginningOfDayWithDayOffset(t *testing.T) {
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
		t.Logf("No%2d, %4v day %6s: first = %v, end = %v\n",
			i,
			math.Abs(float64(item)),
			when,
			BeginningOfDayWithDayOffset(item), // start
			EndOfDayWithDayOffset(item),       // end
		)
	}
}

func TestBeginningOfMonthAgo(t *testing.T) {
	in := []time.Time{
		BeginningOfMonthAgo(),
		EndOfMonthAgo(),
		BeginningOfMonthLater(),
		EndOfMonthLater(),
	}

	for i, item := range in {
		t.Logf("%v, item=%v\n", i, item)
	}
}

func TestLastMonthFirstDay(t *testing.T) {
	t.Log(12 % 12)
	t.Log(24%12, 24/12, -24%12, -24/12, -2%12, -2/12)
	t.Log("1+-2=", 1+-2)
	t.Log(math.Mod(12, 12))
	t.Log(math.Mod(24, 12))
	t.Log(math.Remainder(24, 12))
}

func TestOffset(t *testing.T) {
	at := BeginningOfMonthWithMonthOffset(0) // 月首日偏移
	t.Log("at:", at)
	t.Log("at:", at.String())

	in := []int{
		0,
		1,
		2,
	}

	for _, item := range in {
		t.Log("day:", item, BeginningOfDayWithDayOffset(item)) // 当天偏移
	}
}

func Test(t *testing.T) {
	day1 := BeginningOfMonthWithMonthOffset(0)
	t.Log(day1)

	day2 := day1.Add(5 * 24 * time.Hour)
	t.Log("after:", day2)

	day3 := BeginOfMonthWithDayOffset(0, -2)
	t.Log("before:", day3)
}
