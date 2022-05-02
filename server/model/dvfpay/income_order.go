// 自动生成模板IncomeOrder
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// IncomeOrder 结构体
// 如果含有time.Time 请自行import time包
type IncomeOrder struct {
	global.GVA_MODEL
	ArrivalTime *time.Time `json:"arrivalTime" form:"arrivalTime" gorm:"column:arrival_time;comment:;"`
	OrderId     string     `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;"`
	Amount      *int       `json:"amount" form:"amount" gorm:"column:amount;comment:;"`
	Currency    string     `json:"currency" form:"currency" gorm:"column:currency;comment:;"`
	Status      string     `json:"status" form:"status" gorm:"column:status;comment:;"`
	Payer       string     `json:"payer" form:"payer" gorm:"column:payer;comment:;"`
	Remark      string     `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
	MerchantId  *uint      `json:"merchantId" form:"merchantId" gorm:"column:merchantId;comment:;"`
	GatewayId   *uint      `json:"gatewayId" form:"gatewayId" gorm:"column:gatewayId;comment:;"`
}

// TableName IncomeOrder 表名
func (IncomeOrder) TableName() string {
	return "dvfpay_income_order"
}
