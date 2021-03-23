package convert

import (
	"net/url"

	"github.com/fatih/structs"
	"github.com/google/go-querystring/query"
)

// 更通用版本:
func StructToMap(v interface{}) map[string]interface{} {
	return structs.Map(v)
}

/*
go struct to map[string][]string

tips:
	- struct fields need capitalize the first letter
*/
func StructToMapSlice(v interface{}) (url.Values, error) {
	return query.Values(v)
}
