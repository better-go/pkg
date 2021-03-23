package convert

import "testing"

func TestMapSliceToMap(t *testing.T) {
	// 字段必须是大写:
	in := struct {
		Name     string `json:"name"`
		Age      int64
		Address  string `json:"address"`
		Meta     []byte
		Children []string
	}{
		Name:     "jim",
		Age:      22,
		Address:  "shanghai",
		Meta:     []byte("test meta"),
		Children: []string{"tom", "bob", "kate"}, // TODO: 注意不安全转换, 会丢数据!!! 确信自己的转换是安全的
	}

	resp, err := StructToStringsDict(in)
	t.Logf("convert to map slice: %+v, err: %v", resp, err)
	t.Logf("convert to url string: %+v", resp.Encode())

	out := StringsDictToStringDict(resp)
	t.Logf("convert to map: %+v", out)
	t.Logf("convert to url string: %+v", out.Encode())
}
