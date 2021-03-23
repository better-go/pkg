package dict

import "testing"

func TestStringDict_Encode(t *testing.T) {
	in := []StringDict{
		{
			"name":     "jim",
			"age":      "20",
			"address":  "shanghai",
			"children": "bob, kate, tom",
			"dog":      "frank",
		},
		{
			"name":     "jim",
			"age":      "20",
			"address":  "shanghai-pudong",
			"children": "bob",
			"dog":      "frank",
		},
	}

	for _, item := range in {
		resp := item.Encode()
		t.Logf("in: %+v", item)
		t.Logf("convert to url string: %+v", resp)
	}
}
