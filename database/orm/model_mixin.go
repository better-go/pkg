package orm

import (
	"strings"

	"github.com/better-go/pkg/x/go-zero/sql/builder"
)

// gen sql:
type SqlBuilderMixin struct {
	_sqlSelectFields string // not db table fields
}

// model: 传数据实体的类型
func (m *SqlBuilderMixin) ToSelectFields(model interface{}) string {
	if m._sqlSelectFields == "" {
		fields := builder.FieldNames(model)
		//fields = builderx.FieldNames(m)
		m._sqlSelectFields = strings.Join(fields, ",")
	}
	return m._sqlSelectFields
}
