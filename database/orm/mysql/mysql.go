package mysql

import (
	"github.com/better-go/pkg/database/orm"

	// database driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/sqlite"
	//  _ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	gormV2 "gorm.io/gorm"
)

type Client struct {
	// 后续支持集成多个 orm 包
	v1 *gorm.DB
	v2 *gormV2.DB
}

func NewClient(opts *orm.Options) *Client {
	// gorm v2:
	dbV2, _ := gormV2.Open(mysql.Open(opts.DSN), nil)

	return &Client{
		v1: NewMySQL(opts),
		v2: dbV2,
	}
}

// NewMySQL new v1 and retry connection when has error.
func NewMySQL(opts *orm.Options) *gorm.DB {
	opt := orm.NewOptions(
		orm.Dialect(orm.MySQL),
		orm.DSN(opts.DSN),
		orm.ConnParams(opts.ActiveNum, opts.IdleNum, opts.IdleTimeout),
		orm.TableFields(opts.CreatedTsName, opts.UpdatedTsName, opts.DeletedTsName, opts.IsDeletedName),
		orm.SingularTable(opts.IsSingularTable), // 单数
		orm.DebugMode(opts.IsDebugMode),         // debug log
	)

	// conn:
	conn := opt.DBConn()
	return conn
}

func (m *Client) DB() *gorm.DB {
	return m.v1
}

func (m *Client) DBv2() *gormV2.DB {
	return m.v2
}

func (m *Client) Close() error {
	return m.v1.Close()
}
