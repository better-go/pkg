package orm

import "testing"

func TestSqlBuilderMixin_SqlSelectFields(t *testing.T) {
	item := new(StatusModel)
	sql := item.ToSelectFields(item)
	t.Logf("gen sql select fields: %+v", sql)
}
