package time

import "testing"

func TestCurrentMonth(t *testing.T) {
	t.Log(CurrentMonth())
	t.Log(LastMonth())
	t.Log(NextMonth())

	in := []int{
		-5,
		-13,
		-1,
		0,
		1,
		2,
		100,
	}

	for i, item := range in {
		t.Logf("No%2d, offest=%3v, month=%v\n", i, item, MonthWithOffset(item))
	}
}

func TestTodayWithOffset(t *testing.T) {
	t.Log(Today())
	t.Log(Yesterday())
	t.Log(Tomorrow())

	in := []int{
		-5,
		-13,
		-1,
		0,
		1,
		2,
		100,
	}

	for i, item := range in {
		t.Logf("No%2d, offest=%3v, day=%v\n", i, item, TodayWithOffset(item))
	}
}
