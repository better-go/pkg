package mysql

import (
	"github.com/better-go/pkg/database/orm"

	// database driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/sqlite"
	//  _ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jinzhu/gorm"
)

type Client struct {
	// 后续支持集成多个 orm 包
	db *gorm.DB
}

func NewClient(opts *orm.Options) *Client {
	return &Client{
		db: NewMySQL(opts),
	}
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(opts *orm.Options) *gorm.DB {
	opt := orm.NewOptions(
		orm.Dialect(orm.MySQL),
		orm.DSN(opts.DSN),
		orm.ConnParams(opts.ActiveNum, opts.IdleNum, opts.IdleTimeout),
		orm.TableFields(opts.CreatedTsName, opts.UpdatedTsName, opts.DeletedTsName, opts.IsDeletedName),
	)

	// conn:
	conn := opt.DBConn()
	// db 默认表名复数: https://gorm.io/zh_CN/docs/conventions.html
	conn.SingularTable(opts.IsSingularTable)
	return conn
}

func (m *Client) DB() *gorm.DB {
	return m.db
}

func (m *Client) Close() error {
	return m.db.Close()
}
