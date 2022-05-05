// 自动生成模板PayoutGatewayAuth
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayoutGatewayAuth 结构体
// 如果含有time.Time 请自行import time包
type PayoutGatewayAuth struct {
	global.GVA_MODEL
	PayoutGatewayId *uint `json:"payoutGatewayId" form:"payoutGatewayId" gorm:"column:payout_gateway_id;comment:;"`
	MerchantId      *uint `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Fee             *int  `json:"fee" form:"fee" gorm:"column:fee;comment:;"`
	LimitMax        *int  `json:"limitMax" form:"limitMax" gorm:"column:limit_max;comment:;"`
	LimitMin        *int  `json:"limitMin" form:"limitMin" gorm:"column:limit_min;comment:;"`
	LimitDay        *int  `json:"limitDay" form:"limitDay" gorm:"column:limit_day;comment:;"`
	LimitTotal      *int  `json:"limitTotal" form:"limitTotal" gorm:"column:limit_total;comment:;"`
}

// TableName PayoutGatewayAuth 表名
func (PayoutGatewayAuth) TableName() string {
	return "dvfpay_payout_gateway_auth"
}
