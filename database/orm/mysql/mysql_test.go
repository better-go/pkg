package mysql

import (
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/better-go/pkg/database/orm"
)

const (
	testDsn   = "dev:dev@tcp(127.0.0.1:13306)/dev?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4"
	testTable = "order_exchange"
)

// 测试表结构:
type testModel struct {
	orm.TxModel

	Mid          int64
	OrderNo      string
	ProductTitle string
}

// 表名:
func (m *testModel) TableName() string {
	return testTable
}

func TestNewMySQL(t *testing.T) {
	var (
		obj1, obj2 testModel
		//v   interface{}
	)
	objList1 := make([]*testModel, 0, 0)
	objList2 := make([]*testModel, 0, 0)

	db := NewMySQL(&orm.Options{
		Dialect:     "",
		DSN:         testDsn,
		IsDebugMode: false, // show raw log
	})

	t.Logf("stats: %+v", db.DB().Stats())
	//t.Log(v1.LogMode(true))

	//
	// orm query:
	//

	// first:
	db.First(&obj1)
	// show raw sql 2:
	db.Debug().First(&obj1)
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
		item := new(testModel)
		rows.Scan(&item.TxModel.ID, &item.Mid, &item.OrderNo, &item.ProductTitle)
		objList1 = append(objList1, item)
	}
	for _, item := range objList1 {
		t.Logf("raw sql query many way2: %+v", item)
	}
}

func TestZeroTimestamp(t *testing.T) {
	ts, _ := time.Parse("1/2/2006 15:04:05", "01/01/0001 00:00:00")
	t.Logf("zero time: %v", ts)
}

func v2Client() *gorm.DB {
	client := NewClient(&orm.Options{
		Dialect:     "",
		DSN:         testDsn,
		IsDebugMode: true, // show raw log
	})

	// use v2 v1:
	db := client.DB()
	db.Debug()
	return db
}

func TestCreateTable(t *testing.T) {
	db := v2Client()

	req := &testModel{

		Mid:          33,
		OrderNo:      "214",
		ProductTitle: "test create",
	}
	t.Logf("%v", req)

	//
	db.Migrator().DropTable(&testModel{})
	db.Migrator().AutoMigrate(&testModel{})

	// create table:
	db.Create(req)
}

func TestInsert(t *testing.T) {
	db := v2Client()

	in := []*testModel{
		{
			Mid:          11,
			OrderNo:      "order11",
			ProductTitle: "product 1",
		},
		{
			Mid:          12,
			OrderNo:      "order12",
			ProductTitle: "product 1",
		},
		{
			Mid:          13,
			OrderNo:      "order13",
			ProductTitle: "product 1",
		},
	}

	// insert:
	for _, item := range in {
		db.Create(item)
	}

}

func TestQuery(t *testing.T) {
	resp := new(testModel)
	db := v2Client()

	// query 1:
	err := db.Model(&testModel{}).Where("mid=?", 33).First(&resp).Error

	// query:
	//err := db.Where(&testModel{Mid: 33,}).Find(&resp).Error

	t.Logf("resp: %+v, err:%v", resp, err)

	batchQuery := []*testModel{
		{
			Mid: 11,
		},
		{
			Mid: 12,
		},
		{
			Mid: 13,
		},
	}

	for _, req := range batchQuery {
		resp := new(testModel)

		// query2:
		err := db.Where(&req).Find(&resp).Error
		t.Logf("query item: %+v, err: %v", resp, err)
	}

}
