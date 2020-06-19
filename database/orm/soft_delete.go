package orm

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	//
	// timestamp类型取值范围：1970-01-01 00:00:00 到 2037-12-31 23:59:59，
	// 初始值调整为: [1970-01-02 00:00:00 , 1970-01-01 00:00:01], 自定义零值, > 1970-01-01 00:00:00, 即可
	// https://reading.developerlearning.cn/discuss/2019-06-19-gorm-mysql-timestamp/
	MySqlZeroTimestamp = "1970-01-02 00:00:00" // 时间戳零值, 特殊值(MySQL)

)

// 软删除类型:
type DeletedAt gorm.DeletedAt

// Scan implements the Scanner interface.
func (n *DeletedAt) Scan(value interface{}) error {
	return (*gorm.DeletedAt)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n DeletedAt) Value() (driver.Value, error) {
	return (gorm.DeletedAt)(n).Value()
}

//
// 软删除-查询子句:(修正 判空条件)
//    - 自动补全
//    - where (deleted_at <= "1970-01-02 00:00:00" )
//
func (DeletedAt) QueryClauses() []clause.Interface {
	return []clause.Interface{
		clause.Where{Exprs: []clause.Expression{
			// <= zeroTime
			clause.Lte{
				Column: clause.Column{Table: clause.CurrentTable, Name: "deleted_at"},
				Value:  zeroTime(), // 默认零值
			},
		}},
	}
}

// 软删除-删除子句:
func (n DeletedAt) DeleteClauses() []clause.Interface {
	return (gorm.DeletedAt)(n).DeleteClauses()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// deleted_at = default zeroTime [zero time: 1970-01-02 00:00:00 +0000 UTC]
func zeroTime() time.Time {
	// zero time: 1970-01-02 00:00:00 +0000 UTC
	zero := "01/02/1970 00:00:00"
	ts, _ := time.Parse("1/2/2006 15:04:05", zero)
	return ts
}

func zeroTimeStr() string {
	return MySqlZeroTimestamp
}
