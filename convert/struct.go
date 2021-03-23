package convert

import (
	"net/url"

	"github.com/better-go/pkg/container/dict"
	"github.com/fatih/structs"
	"github.com/google/go-querystring/query"
)

/*
go struct to map[string]xxx

tips:
	- struct fields need capitalize the first letter
*/

// go struct to map[string]interface{}
func StructToDict(in interface{}) (out dict.Dict) {
	return structs.Map(in)
}

// go struct to map[string]string
func StructToStringDict(in interface{}) (out dict.StringDict, err error) {
	// to map slice:
	ds, err := StructToStringsDict(in)
	if err != nil {
		return nil, err
	}

	// convert to map:
	out = StringsDictToStringDict(ds)
	return out, nil
}

// go struct to map[string][]string
func StructToStringsDict(in interface{}) (out url.Values, err error) {
	return query.Values(in)
}
