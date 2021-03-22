package convert

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

/*
go struct to map[string][]string

tips:
	- struct fields need capitalize the first letter
*/
func StructToMap(v interface{}) (url.Values, error) {
	return query.Values(v)
}
