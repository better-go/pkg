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
type StringDict map[string]string

func (m StringDict) Get(key string) string {
	return m[key]
}

func (m StringDict) Set(key string, value string) {
	m[key] = value
}

// Del deletes the values associated with key.
func (m StringDict) Del(key string) {
	delete(m, key)
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (m StringDict) Encode() string {
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
		valueEscaped := url.QueryEscape(v)

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
