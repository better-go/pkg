package time

import (
	"testing"
	"time"
)

func TestCompare(t *testing.T) {
	in := []struct {
		A time.Time
		B time.Time
	}{
		{
			A: BeginningOfMonthWithMonthOffset(0),
			B: BeginningOfMonthWithMonthOffset(0),
		},
		{
			A: BeginningOfMonthWithMonthOffset(0),
			B: BeginningOfMonthWithMonthOffset(-1),
		},
		{
			A: BeginningOfMonthWithMonthOffset(0),
			B: BeginningOfMonthWithMonthOffset(1),
		},
	}

	for _, item := range in {
		t.Log("a=", item.A, "b=", item.B)
		t.Log("\ta vs b:", Compare(item.A, item.B))
		t.Log("\tIsEqual:", IsEqual(item.A, item.B))
		t.Log("\tIsBefore:", IsBefore(item.A, item.B))
		t.Log("\tIsAfter:", IsAfter(item.A, item.B))
		t.Log("\tstd.before:", item.A.Before(item.B))
		t.Log("\tstd.after:", item.A.After(item.B))
	}

}
