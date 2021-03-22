package convert

import "testing"

func TestStructToMap(t *testing.T) {
	// 字段必须是大写:
	in := struct {
		Name    string
		Age     int64
		Address string
		Meta    []byte
	}{
		Name:    "jim",
		Age:     22,
		Address: "shanghai",
		Meta:    []byte("test meta"),
	}

	resp, err := StructToMap(in)
	t.Logf("convert to map: %+v, err: %v", resp, err)
	t.Logf("convert to url string: %+v", resp.Encode())
}
