package convert

import "testing"

func TestInt64ToString(t *testing.T) {
	var i int64
	var s string
	i = 60

	s = Int64ToString(i)
	t.Log("s=", s)

}

func TestStringToInt64(t *testing.T) {
	var i int64
	s := "2233"

	i, _ = StringToInt64(s)
	t.Log("i=", i)
}
