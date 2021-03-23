package dict

import (
	"net/url"
	"sort"
	"strings"
)

/*
对 map 字段 进行 url 参数拼接
	- 参考 标准库 "net/url" 的 url.Values 实现
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

		// TODO: need fix type: string/[]bytes
		valueEscaped := url.QueryEscape(v.(string))

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
