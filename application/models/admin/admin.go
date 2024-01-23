package admin

import (
	"github.com/hhttco/GinDemo/pkg/db"
	"time"
)

func init() {
	db.DoSynBeans(new(Admin))
}

/**
 * 管理员表
 *
 */
type Admin struct {
	Id       int
	Mobile   string
	Password string
	Status   int
	CreateAt time.Time
}

//返回表名
func (t *Admin) TableName() string {
	return "admin"
}
