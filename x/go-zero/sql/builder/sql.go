package builder

import (
	"fmt"
	"reflect"

	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

/*
go-zero sql builder:
	- 参考 builderx.FieldNames() 实现
	- 修复2个问题:
		- 支持 匿名嵌套字段解析
		- 支持 忽略不含 `db: "xxx"` tag 的字段类型

*/

const (
	dbTag = "db"
)

// ToMap converts interface into map
func ToMap(in interface{}) map[string]interface{} {
	return builderx.ToMap(in)
}

// FieldNames converts golang struct field into slice string
func FieldNames(in interface{}) (out []string) {
	v := reflect.ValueOf(in)
	return nestParseField(v)
}

func nestParseField(v reflect.Value) (out []string) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		ft := fi.Type

		// recursion: 匿名+结构体类型+has tag
		if fi.Anonymous && ft.Kind() == reflect.Struct {
			out = append(out, nestParseField(v.Field(i))...)
		} else {
			if tagV := fi.Tag.Get(dbTag); tagV != "" {
				out = append(out, fmt.Sprintf("`%s`", tagV))
			} else {
				//out = append(out, fmt.Sprintf(`"%s"`, fi.Name))
			}
		}
	}
	return out
}
