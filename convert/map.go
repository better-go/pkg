package convert

import (
	"net/url"

	"github.com/better-go/pkg/container/dict"
	"github.com/better-go/pkg/log"
)

// map[string][]string -> map[string]interface{}
func StringsDictToDict(in url.Values) (out dict.Dict) {
	out = make(dict.Dict)

	for k, v := range in {
		if len(v) > 1 {
			log.Infof("unsafe convert this field: key=%v, values=%+v", k, v)
		}
		out[k] = v[0]
	}
	return out
}

// 注意: 这个方法不安全, 要确定 值, 只有一个才可以使用, 否则丢数据
func StringsDictToStringDict(in url.Values) (out dict.StringDict) {
	out = make(dict.StringDict)

	for k, v := range in {
		if len(v) > 1 {
			log.Infof("unsafe convert this field: key=%v, values=%+v", k, v)
		}
		out[k] = v[0]
	}
	return out
}
