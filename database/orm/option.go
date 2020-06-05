package orm

import (
	"time"

	xtime "github.com/better-go/pkg/time"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

/*

ref:
	- https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	- gorm 默认字段命名同此: https://gorm.io/zh_CN/docs/conventions.html
	- gorm 2.0 特性:
		- https://www.youtube.com/watch?v=NCZHe6zb2Sg
		- https://github.com/talk-go/night/issues/511
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
	DSN         string         // data source name.
	ActiveNum   int            // pool
	IdleNum     int            // pool
	IdleTimeout xtime.Duration // connect max life time.

	// option for table:
	IsSingularTable bool // orm 默认表名: 单数

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
	var expire xtime.Duration
	expire.UnmarshalText([]byte(DefaultIdleExpire))

	// default:
	opt := Options{
		Dialect:     DefaultDialect,
		ActiveNum:   DefaultActiveNum,
		IdleNum:     DefaultIdleNum,
		IdleTimeout: expire,
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

// 连接参数定制:
func ConnParams(activeNum int, idleNum int, idleExpire xtime.Duration) OptionFunc {
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
	// debug mode:
	//conn.Debug()
	// log:
	//conn.SetLogger(logAdapter{})

	// auto hook:
	conn.Callback().Create().Replace("gorm:update_time_stamp", m.autoCreatedFields)
	conn.Callback().Update().Replace("gorm:update_time_stamp", m.autoUpdatedFields)
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
			m.DeletedTsName,
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
		if updatedAtField, ok := scope.FieldByName(m.IsDeletedName); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(false)
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