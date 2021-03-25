package orm

import (
	"time"
)

/*
ref:
	mysql 数据类型:
		- https://www.runoob.com/mysql/mysql-data-types.html


*/

type Model struct {
	ID        uint64    `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	CreatedAt time.Time `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "` // 删除默认时间戳
	// 1970-01-01 00:00:00 // 1970-01-02 00:00:00
}

// user model:
type UserModel struct {
	ID        uint64    `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	UserID    uint64    `db:"user_d"     gorm:"index;       type:BIGINT(20) unsigned; NOT NULL;              DEFAULT:'0';                                           COMMENT:'用户 ID' "` // 用户 ID
	CreatedAt time.Time `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}

// 唯一索引:
type UserUniqueModel struct {
	ID        uint64    `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	UserID    uint64    `db:"user_d"     gorm:"UNIQUE;      type:BIGINT(20) unsigned; NOT NULL;              DEFAULT:'0';                                           COMMENT:'用户 ID' "` // 用户 ID
	CreatedAt time.Time `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}

// has status:
type StatusModel struct {
	ID        uint64    `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Status    int64     `db:"status"     gorm:"index;       type:int(11);          NOT NULL;                 DEFAULT:'-1';                                          COMMENT:'状态: (-1:无效, 1:正常, >1: 正常中间状态, <-1: 异常中间状态)'; "`
	CreatedAt time.Time `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "` // 删除默认时间戳
}

// db 事务: 乐观锁
type TxModel struct {
	ID        uint64    `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Ver       uint64    `db:"ver"        gorm:"             type:int(11) unsigned; NOT NULL; AUTO_INCREMENT; DEFAULT:'1';                                           COMMENT:'乐观锁版本号' "` // 乐观锁
	CreatedAt time.Time `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
}

// 集大成者:
type TemplateModel struct {
	ID        uint64    `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Ver       uint64    `db:"ver"        gorm:"             type:int(11) unsigned; NOT NULL; AUTO_INCREMENT; DEFAULT:'1';                                           COMMENT:'乐观锁版本号' "` // 乐观锁
	CreatedAt time.Time `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "`
	//
	Name       string `db:"name"         gorm:"type=varchar(128) CHARACTER SET utf8mb4;  NOT NULL; DEFAULT:''; COMMENT:'姓名'; "`
	BankCardNo int64  `db:"bank_card_no" gorm:"type=int(11);                             NOT NULL; DEFAULT:''; COMMENT:'银行卡号'; "`
}
