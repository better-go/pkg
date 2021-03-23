package convert

import "testing"

func TestStructToMapSlice(t *testing.T) {
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

	resp, err := StructToStringsDict(in)
	t.Logf("convert to map: %+v, err: %v", resp, err)
	t.Logf("convert to url string: %+v", resp.Encode())
}

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

	resp := StructToDict(in)
	t.Logf("convert to map: %+v", resp)
}

func TestStructToStringDict(t *testing.T) {
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

	resp, _ := StructToStringDict(in)
	t.Logf("convert to map: %+v", resp)
	t.Logf("convert to url string: %v", resp.Encode())
}
