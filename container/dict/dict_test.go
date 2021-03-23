package dict

import (
	"testing"
)

func TestDict_Encode(t *testing.T) {
	in := []Dict{
		{
			"name":     "jim",
			"age":      "20",
			"address":  "shanghai",
			"children": []string{"bob", "jackie"}, // TODO: 无法正常处理, 注意
			"dog":      "frank",
		},
		{
			"name":     "jim",
			"age":      "20",
			"address":  "shanghai-pudong",
			"children": "bob",
			"dog":      "frank",
		},

		{
			"name":     "jim",
			"age":      "20",
			"address":  "shanghai-pudong",
			"children": []byte("bob"), // TODO: 无法正常处理, 注意
			"dog":      "frank",
		},
	}

	for _, item := range in {
		resp := item.Encode()
		t.Logf("in: %+v", item)
		t.Logf("convert to url string: %+v", resp)
	}
}
