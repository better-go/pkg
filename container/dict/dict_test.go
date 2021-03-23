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
			"age":      20,
			"score":    -20.22,
			"address":  "shanghai-pudong",
			"children": []byte("bob"), // TODO: 做了兼容处理, 转换成 string, 注意要小心! 不一定符合预期
			"dog":      "frank",
		},

		// 嵌套 map:
		{
			"name":     "jim",
			"age":      20,
			"score":    -20.22,
			"address":  "shanghai-pudong",
			"children": []byte("bob"), // TODO: 做了兼容处理, 转换成 string, 注意要小心! 不一定符合预期
			"dog":      "frank",
			"course": map[string]interface{}{
				"math": 20,
				"cs":   100,
				"pe":   12,
			},
		},
	}

	for _, item := range in {
		resp := item.Encode()
		t.Logf("in: %+v", item)
		t.Logf("convert to url string: %+v", resp)
	}
}
