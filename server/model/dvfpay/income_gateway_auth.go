// 自动生成模板IncomeGatewayAuth
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// IncomeGatewayAuth 结构体
// 如果含有time.Time 请自行import time包
type IncomeGatewayAuth struct {
	global.GVA_MODEL
	IncomeGatewayId *uint             `json:"incomeGatewayId" form:"incomeGatewayId" gorm:"column:income_gateway_id;comment:;"`
	IncomeGateway   IncomeGateway     `json:"incomeGateway"`
	Merchants       []*system.SysUser `json:"merchants" gorm:"many2many:dvfpay_income_gateway_auth_merchants;"`
	Fee             *int              `json:"fee" form:"fee" gorm:"column:fee;comment:;"`
	LimitMax        *int              `json:"limitMax" form:"limitMax" gorm:"column:limit_max;comment:;"`
	LimitMin        *int              `json:"limitMin" form:"limitMin" gorm:"column:limit_min;comment:;"`
	LimitDay        *int              `json:"limitDay" form:"limitDay" gorm:"column:limit_day;comment:;"`
	LimitTotal      *int              `json:"limitTotal" form:"limitTotal" gorm:"column:limit_total;comment:;"`
	Explain         string            `json:"explain" form:"explain" gorm:"column:explain;type:text;comment:;"`
}

// TableName IncomeGatewayAuth 表名
func (IncomeGatewayAuth) TableName() string {
	return "dvfpay_income_gateway_auth"
}
