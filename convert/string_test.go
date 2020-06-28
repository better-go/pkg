package convert

import "testing"

func TestInt64ToString(t *testing.T) {
	var i1 int64
	var i2 int
	var s1, s2 string

	i1 = 60
	i2 = 30

	s1 = Int64ToString(i1)
	s2 = IntToString(i2)

	t.Log("s1=", s1)
	t.Log("s2=", s2)

}

func TestStringToInt64(t *testing.T) {
	var i1 int64
	var i2 int

	s1 := "2233"
	s2 := "3456"

	i1, _ = StringToInt64(s1)
	i2, _ = StringToInt(s2)

	t.Log("i1=", i1)
	t.Log("i2=", i2)
}
