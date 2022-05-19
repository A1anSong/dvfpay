// 自动生成模板PayoutGatewayAuth
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// PayoutGatewayAuth 结构体
// 如果含有time.Time 请自行import time包
type PayoutGatewayAuth struct {
	global.GVA_MODEL
	PayoutGatewayId *uint             `json:"payoutGatewayId" form:"payoutGatewayId" gorm:"column:payout_gateway_id;comment:;"`
	PayoutGateway   PayoutGateway     `json:"payoutGateway"`
	Merchants       []*system.SysUser `json:"merchants" gorm:"many2many:dvfpay_payout_gateway_auth_merchants;"`
	Fee             *int              `json:"fee" form:"fee" gorm:"column:fee;comment:;"`
	LimitMax        *int              `json:"limitMax" form:"limitMax" gorm:"column:limit_max;comment:;"`
	LimitMin        *int              `json:"limitMin" form:"limitMin" gorm:"column:limit_min;comment:;"`
	LimitDay        *int              `json:"limitDay" form:"limitDay" gorm:"column:limit_day;comment:;"`
	LimitTotal      *int              `json:"limitTotal" form:"limitTotal" gorm:"column:limit_total;comment:;"`
	Explain         string            `json:"explain" form:"explain" gorm:"column:explain;comment:;"`
}

// TableName PayoutGatewayAuth 表名
func (PayoutGatewayAuth) TableName() string {
	return "dvfpay_payout_gateway_auth"
}
