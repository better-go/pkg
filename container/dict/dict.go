package dict

import (
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"github.com/better-go/pkg/log"
)

/*
对 map 字段 进行 url 参数拼接
	- 参考 标准库 "net/url" 的 url.Values 实现

TODO:
	针对 encode() 特别说明:
		1. 对于 嵌套 struct 的数据类型, 会在首层迭代时, 转换成 json string, 再 encode()
		2. 这里对 interface{} 数据类型的断言, 是不完善的, 暂时不处理
*/
type Dict map[string]interface{}

func (m Dict) Get(key string) interface{} {
	return m[key]
}

func (m Dict) Set(key string, value interface{}) {
	m[key] = value
}

// Del deletes the values associated with key.
func (m Dict) Del(key string) {
	delete(m, key)
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (m Dict) Encode() string {
	if m == nil {
		return ""
	}

	var buf strings.Builder
	keys := make([]string, 0, len(m))

	// keys:
	for k := range m {
		keys = append(keys, k)
	}

	// sort keys:
	sort.Strings(keys)

	// iter keys:
	for _, k := range keys {
		v := m[k]
		keyEscaped := url.QueryEscape(k)
		var valueEscaped string

		// TODO: need fix type: string/[]bytes // 对 int 类型, 处理并不完善, 建议自主转换成 string 类型, 再传入
		// assert type: 兼容性处理
		switch item := v.(type) {
		case string:
			valueEscaped = url.QueryEscape(item)
		//case []byte:
		//	valueEscaped = url.QueryEscape(string(item))
		//case int:
		//	s := strconv.Itoa(item)
		//	valueEscaped = url.QueryEscape(s)
		default:
			s, err := json.Marshal(item) // TODO: 无法识别的对象类型, 先转换成 json, 再处理
			if err != nil {
				log.Infof("invalid type, skip convert this field, key=%v, value=%+v, err=%v", k, v, err)
				continue
			}
			valueEscaped = url.QueryEscape(string(s))
			log.Infof("object field transfer to json string, then do convert, key=%v, value=%+v, string=%v", k, v, string(s))
		}

		// combine pairs:
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		// set one pair: k=v
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(valueEscaped)
	}
	return buf.String()
}
