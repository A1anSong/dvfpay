// 自动生成模板PayoutGateway
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayoutGateway 结构体
// 如果含有time.Time 请自行import time包
type PayoutGateway struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Type      string `json:"type" form:"type" gorm:"column:type;comment:;"`
	Currency  string `json:"currency" form:"currency" gorm:"column:currency;comment:;"`
	Parameter string `json:"parameter" form:"parameter" gorm:"column:parameter;comment:;"`
	Status    string `json:"status" form:"status" gorm:"column:status;comment:;"`
	LimitMax  *int   `json:"limitMax" form:"limitMax" gorm:"column:limit_max;comment:;"`
	LimitMin  *int   `json:"limitMin" form:"limitMin" gorm:"column:limit_min;comment:;"`
	Fee       *int   `json:"fee" form:"fee" gorm:"column:fee;comment:;"`
}

// TableName PayoutGateway 表名
func (PayoutGateway) TableName() string {
	return "dvfpay_payout_gateway"
}
