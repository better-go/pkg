package orm

import (
	"time"
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
	DeletedAt DeletedAt `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "` // 删除默认时间戳
	// 1970-01-01 00:00:00 // 1970-01-02 00:00:00
}

// user model:
type UserModel struct {
	ID        uint64    `gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	UserID    uint64    `gorm:"index;       type:int(11) unsigned; NOT NULL;                 DEFAULT:'0';                                           COMMENT:'用户 ID' "` // 用户 ID
	CreatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}

// db 事务: 乐观锁
type TxModel struct {
	ID        uint64    `gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Ver       uint64    `gorm:"             type:int(11) unsigned; NOT NULL; AUTO_INCREMENT; DEFAULT:'1';                                           COMMENT:'乐观锁版本号' "` // 乐观锁
	CreatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}

// 集大成者:
type TemplateModel struct {
	ID        uint64    `gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Ver       uint64    `gorm:"             type:int(11) unsigned; NOT NULL; AUTO_INCREMENT; DEFAULT:'1';                                           COMMENT:'乐观锁版本号' "` // 乐观锁
	CreatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
	//
	Name       string `gorm:"type=varchar(128) CHARACTER SET utf8mb4;  NOT NULL; DEFAULT:''; COMMENT:'姓名'; "`
	BankCardNo int64  `gorm:"type=int(11);                             NOT NULL; DEFAULT:''; COMMENT:'银行卡号'; "`
}
