package orm

import (
	"time"
)

const (
	//
	// timestamp类型取值范围：1970-01-01 00:00:00 到 2037-12-31 23:59:59，
	// 初始值调整为: [1970-01-02 00:00:00 , 1970-01-01 00:00:01], 自定义零值, > 1970-01-01 00:00:00, 即可
	//
	MySqlZeroTimestamp = "1970-01-02 00:00:00" // 时间戳零值, 特殊值(MySQL)

)

// ref: github.com/jinzhu/gorm/model.go:9
// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        uint64    `gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	CreatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "` // 删除默认时间戳
	// 1970-01-01 00:00:00 // 1970-01-02 00:00:00
}

// user model:
type UserModel struct {
	ID        uint64    `gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	UserID    uint64    `gorm:"index;       type:int(11) unsigned; NOT NULL;                 DEFAULT:'0';                                           COMMENT:'用户 ID' "` // 用户 ID
	CreatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}

// db 事务: 乐观锁
type TxModel struct {
	ID        uint64    `gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Ver       uint64    `gorm:"             type:int(11) unsigned; NOT NULL; AUTO_INCREMENT; DEFAULT:'1';                                           COMMENT:'乐观锁版本号' "` // 乐观锁
	CreatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}
