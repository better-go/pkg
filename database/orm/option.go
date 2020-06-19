package orm

import (
	"time"

	"github.com/better-go/pkg/log"
	timeEx "github.com/better-go/pkg/time"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	gormV2 "gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*

ref:
	- https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	- gorm 默认字段命名同此: https://gorm.io/zh_CN/docs/conventions.html
	- gorm 2.0 特性:
		- https://www.youtube.com/watch?v=NCZHe6zb2Sg
		- https://github.com/talk-go/night/issues/511
		- v2 usage:
			- https://github.com/go-gorm/gorm/blob/v0.2.7/tests/tests_test.go
*/

const (
	//
	// db type:
	//
	MySQL  = "mysql"
	SQLite = "sqlite3"
	TiDB   = "mysql"

	//
	// default options:
	//
	DefaultDialect     = MySQL
	DefaultActiveNum   = 20   // 连接数 max
	DefaultIdleNum     = 10   // 连接数 max
	DefaultIdleExpire  = "4h" // 超时 max
	DefaultQueryExpire = "5s" // 查询超时
	DefaultExecExpire  = "5s" // 写入超时
	DefaultTxExpire    = "5s" // tx 事务超时

	//
	// default table column fields name:
	//
	createdAt = "created_at"
	updatedAt = "updated_at"
	deletedAt = "deleted_at"
	isDeleted = "is_deleted"
)

// db conn option:
type Options struct {
	Dialect     string
	DSN         string          // data source name.
	ActiveNum   int             // pool
	IdleNum     int             // pool
	IdleTimeout timeEx.Duration // connect max life time.

	// option for table:
	IsSingularTable bool // db 表名: 单数/复数
	IsDebugMode     bool // log switch: show raw sql

	// option item:
	CreatedTsName string
	UpdatedTsName string
	DeletedTsName string
	IsDeletedName string
}

// fn:
type OptionFunc func(*Options)

//
func NewOptions(opts ...OptionFunc) Options {
	var expire timeEx.Duration
	expire.UnmarshalText([]byte(DefaultIdleExpire))

	// default:
	opt := Options{
		Dialect:         DefaultDialect,
		ActiveNum:       DefaultActiveNum,
		IdleNum:         DefaultIdleNum,
		IdleTimeout:     expire,
		IsSingularTable: false, // 默认复数
		IsDebugMode:     false, // 默认关闭 show raw sql
		// table fields:
		CreatedTsName: createdAt,
		UpdatedTsName: updatedAt,
		DeletedTsName: deletedAt,
		IsDeletedName: isDeleted,
	}

	// set:
	for _, fn := range opts {
		fn(&opt)
	}

	return opt

}

//
func Dialect(dialect string) OptionFunc {
	return func(options *Options) {
		options.Dialect = dialect
	}
}

//
func DSN(dsn string) OptionFunc {
	return func(options *Options) {
		options.DSN = dsn
	}
}

// 表名单数:
func SingularTable(active bool) OptionFunc {
	return func(options *Options) {
		options.IsSingularTable = active
	}
}

// log 模式:
func DebugMode(active bool) OptionFunc {
	return func(options *Options) {
		options.IsDebugMode = active
	}
}

// 连接参数定制:
func ConnParams(activeNum int, idleNum int, idleExpire timeEx.Duration) OptionFunc {
	return func(options *Options) {
		if activeNum > 0 {
			options.ActiveNum = activeNum
		}
		if idleNum > 0 {
			options.IdleNum = idleNum
		}
		if idleExpire > 0 {
			options.IdleTimeout = idleExpire
		}
	}
}

// 表默认字段:
func TableFields(createdAt string, updatedAt string, deletedAt string, isDeleted string) OptionFunc {
	return func(options *Options) {
		if createdAt != "" {
			options.CreatedTsName = createdAt
		}
		if updatedAt != "" {
			options.UpdatedTsName = updatedAt
		}
		if deletedAt != "" {
			options.DeletedTsName = deletedAt
		}
		if isDeleted != "" {
			options.IsDeletedName = isDeleted
		}
	}
}

/////////////////////////////////////////////////////////////////////////////////////

func (m *Options) DBConn() *gorm.DB {
	conn, err := gorm.Open(m.Dialect, m.DSN)
	if err != nil {
		log.Errorf("db dsn(%s) open error: %v", m.DSN, err)
		panic(err)
	}

	// option:
	conn.DB().SetMaxIdleConns(m.IdleNum)
	conn.DB().SetMaxOpenConns(m.ActiveNum)
	conn.DB().SetConnMaxLifetime(time.Duration(m.IdleTimeout))
	// debug mode: log show raw sql
	conn.LogMode(m.IsDebugMode)
	// db 表名: 使用单数形式 https://gorm.io/zh_CN/docs/conventions.html
	conn.SingularTable(m.IsSingularTable)
	// log:
	//conn.SetLogger(logAdapter{})

	// auto hook:
	conn.Callback().Create().Replace("gorm:update_time_stamp", m.autoCreatedFields)
	conn.Callback().Update().Replace("gorm:update_time_stamp", m.autoUpdatedFields)
	// hook:
	// ref: https://gorm.io/zh_CN/docs/write_plugins.html#%E6%B3%A8%E5%86%8Ccallback%E9%A1%BA%E5%BA%8F
	// github.com/jinzhu/gorm@v1.9.12/callback_query.go:17
	conn.Callback().Query().Before("gorm:query").Register("gorm:query_fix_deleted", func(scope *gorm.Scope) {
		// todo: fix deleted field
	})
	return conn
}

func (m *Options) DBConnV2() *gormV2.DB {
	conn, err := gormV2.Open(
		mysql.Open(m.DSN),
		&gormV2.Config{
			SkipDefaultTransaction: false,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: m.IsSingularTable, // 表单数
			},
			Logger:               nil,
			NowFunc:              nil,
			DryRun:               false,
			PrepareStmt:          false,
			DisableAutomaticPing: false,
			ClauseBuilders:       nil,
			ConnPool:             nil,
			Dialector:            nil,
		})

	if err != nil {
		log.Errorf("db dsn(%s) open error: %v", m.DSN, err)
		panic(err)
	}

	// debug log:
	if m.IsDebugMode {
		conn.Debug()
	}

	return conn
}

/////////////////////////////////////////////////////////////////////////////////////

// auto create:
func (m *Options) autoCreatedFields(scope *gorm.Scope) {
	if !scope.HasError() {
		now := gorm.NowFunc()

		// ts:
		tsFields := []string{
			m.CreatedTsName,
			m.UpdatedTsName,
			//m.DeletedTsName,
		}

		// ts:
		for _, item := range tsFields {
			if field, ok := scope.FieldByName(item); ok {
				if field.IsBlank {
					field.Set(now)
				}
			}
		}

		// soft delete:
		if deleted, ok := scope.FieldByName(m.IsDeletedName); ok {
			if deleted.IsBlank {
				zeroTs := zeroTime() // default zero time
				deleted.Set(zeroTs)
			}
		}
	}
}

// auto update:
func (m *Options) autoUpdatedFields(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn(m.UpdatedTsName, gorm.NowFunc())
	}
}

/////////////////////////////////////////////////////////////////////////////////////

//
// log for orm:
//
type logAdapter struct {
}

func (l logAdapter) Print(v ...interface{}) {
	//log.Infof(strings.Repeat("%v ", len(v)), v...)
	log.Info(v...)
}
