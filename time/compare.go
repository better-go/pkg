package time

import (
	"time"
)

/*

参考: 标准库 strings.Compare()


*/

//
func Compare(a time.Time, b time.Time) int32 {
	tsA := a.UnixNano()
	tsB := b.UnixNano()

	if tsA == tsB { // 	strings.Compare()
		return 0
	}

	if tsA < tsB {
		return -1
	}

	return +1
}

// a<b
func IsBefore(a time.Time, b time.Time) bool {
	return Compare(a, b) < 0
}

// a>b
func IsAfter(a time.Time, b time.Time) bool {
	return Compare(a, b) > 0
}

// a=b:
func IsEqual(a time.Time, b time.Time) bool {
	return Compare(a, b) == 0
}
