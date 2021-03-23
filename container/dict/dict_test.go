package dict

import (
	"testing"
)

func TestDict_Encode(t *testing.T) {
	// 字段必须是大写:
	in := Dict{

		"name":     "jim",
		"age":      20,
		"address":  "shanghai",
		"children": []string{"bob", "jackie"},
		"dog":      "frank",
	}

	resp := in.Encode()
	t.Logf("in: %+v", in)
	t.Logf("convert to url string: %+v", resp)
}
