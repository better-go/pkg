package builder

import (
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
)

// has status:
type statusModel struct {
	sqlBuilderMixin // gen sql:

	ID        uint64         `db:"id"         gorm:"primary_key; type:int(11) unsigned; NOT NULL; AUTO_INCREMENT;                                                        COMMENT:'自增主键' "`
	Status    int64          `db:"status"     gorm:"index;       type:int(11);          NOT NULL;                 DEFAULT:'-1';                                          COMMENT:'状态: (-1:无效, 1:正常, >1: 正常中间状态, <-1: 异常中间状态)'; "`
	CreatedAt time.Time      `db:"created_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP;                             COMMENT:'创建时间' "`
	UpdatedAt time.Time      `db:"updated_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:'更新时间' "`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index;       type:timestamp;        NOT NULL;                 DEFAULT:'1970-01-01 00:00:01';                         COMMENT:'删除时间' "` // 删除默认时间戳
}

// gen sql:
type sqlBuilderMixin struct {
	_sqlSelectFields string // not db table fields
}

// model: 传数据实体的类型
func (m *sqlBuilderMixin) ToSelectFields(model interface{}) string {
	if m._sqlSelectFields == "" {
		fields := FieldNames(model)
		//fields = builderx.FieldNames(m)
		m._sqlSelectFields = strings.Join(fields, ",")
	}
	return m._sqlSelectFields
}

// 用户 Meta 信息
type userMeta struct {
	statusModel

	UserID     uint64 `db:"user_id" gorm:"UNIQUE; type:BIGINT(20) unsigned; NOT NULL;      DEFAULT:'0'; COMMENT:'用户 ID'; "`        // 用户 ID
	UniqueName string `db:"unique_name" gorm:"index;  type:varchar(128) CHARACTER SET utf8mb4; DEFAULT:'';  COMMENT:'方案1:唯一昵称'; "` // 二选一方案: 唯一昵称, 不允许重名
	NickName   string `db:"nick_name" gorm:"index;  type:varchar(128) CHARACTER SET utf8mb4; DEFAULT:'';  COMMENT:'方案2:昵称+编号'; "`  // 二选一方案: 昵称 + 昵称编号
	NickNo     int64  `db:"nick_no" gorm:"index;  type:varchar(128) CHARACTER SET utf8mb4; DEFAULT:'';  COMMENT:'方案2:昵称+编号'; "`    // 二选一方案: 昵称编号
	// auth:
	Register   string `db:"register" `    // 注册信息
	MobileNo   string `db:"mobile_no" `   // 手机号
	MobileCode string `db:"mobile_code" ` // 手机号国家码
	Email      string `db:"email" `       // 邮箱
	Password   string `db:"password" `    // 密码

	_sqlSelectFields string `` // sql
}

func TestFieldNames(t *testing.T) {
	m := new(userMeta)
	t.Logf("sql select fields: %v", m.ToSelectFields(m))
}
