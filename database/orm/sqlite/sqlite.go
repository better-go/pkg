package sqlite

import (
	"github.com/better-go/pkg/database/orm"

	"github.com/jinzhu/gorm"
	// database driver
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
ref:
	- https://gorm.io/zh_CN/docs/connecting_to_the_database.html#Sqlite3

*/

// NewSQLite new db and retry connection when has error.
func NewSQLite(opts *orm.Options) *gorm.DB {
	opt := orm.NewOptions(
		orm.Dialect(orm.SQLite),
		orm.DSN(opts.DSN),
		orm.ConnParams(opts.ActiveNum, opts.IdleNum, opts.IdleTimeout),
		orm.TableFields(opts.CreatedTsName, opts.UpdatedTsName, opts.DeletedTsName, opts.IsDeletedName),
	)
	return opt.DBConn()
}
