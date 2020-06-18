package mysql

import (
	"github.com/better-go/pkg/database/orm"
	"testing"
)

const (
	testDsn   = "dev:dev@tcp(127.0.0.1:13306)/dev?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4"
	testTable = "order_exchange"
)

// 测试表结构:
type testSchema struct {
	Id           int64
	Mid          int64
	OrderNo      string
	ProductTitle string
}

// 表名:
func (m *testSchema) TableName() string {
	return testTable
}

func TestNewMySQL(t *testing.T) {
	var (
		obj1, obj2 testSchema
		//v   interface{}
	)
	objList1 := make([]*testSchema, 0, 0)
	objList2 := make([]*testSchema, 0, 0)

	db := NewMySQL(&orm.Options{
		Dialect:     "",
		DSN:         testDsn,
		IsDebugMode: true, // show raw log
	})

	t.Logf("stats: %+v", db.DB().Stats())
	//t.Log(db.LogMode(true))

	//
	// orm query:
	//

	// first:
	db.First(&obj1)
	t.Logf("orm query: %+v", obj1)

	//
	// raw sql: https://gorm.io/zh_CN/docs/sql_builder.html
	//

	// query one:
	db.Raw("select id, mid, order_no, product_title from order_exchange where id= 1").Scan(&obj2)
	t.Logf("raw sql query one: %+v", obj2)

	// query many way1:
	db.Raw("select id, mid, order_no, product_title from order_exchange").Scan(&objList2)
	for _, item := range objList2 {
		t.Logf("raw sql query many way1: %+v", item)
	}

	// query many way2:
	rows, _ := db.Raw("select id, mid, order_no, product_title from order_exchange").Rows()
	defer rows.Close()
	for rows.Next() {
		item := new(testSchema)
		rows.Scan(&item.Id, &item.Mid, &item.OrderNo, &item.ProductTitle)
		objList1 = append(objList1, item)
	}
	for _, item := range objList1 {
		t.Logf("raw sql query many way2: %+v", item)
	}
}
