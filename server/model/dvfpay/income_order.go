// 自动生成模板IncomeOrder
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"time"
)

// IncomeOrder 结构体
// 如果含有time.Time 请自行import time包
type IncomeOrder struct {
	global.GVA_MODEL
	ArrivalTime         *time.Time        `json:"arrivalTime" form:"arrivalTime" gorm:"column:arrival_time;comment:;"`
	OrderId             string            `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;"`
	Amount              *int              `json:"amount" form:"amount" gorm:"column:amount;comment:;"`
	Currency            string            `json:"currency" form:"currency" gorm:"column:currency;comment:;"`
	Status              string            `json:"status" form:"status" gorm:"column:status;comment:;"`
	Payer               string            `json:"payer" form:"payer" gorm:"column:payer;comment:;"`
	Remark              string            `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
	MerchantId          *uint             `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Merchant            system.SysUser    `json:"merchant" gorm:"foreignKey:MerchantId"`
	IncomeGatewayAuthId *uint             `json:"incomeGatewayAuthId" form:"incomeGatewayAuthId" gorm:"column:income_gateway_auth_id;comment:;"`
	IncomeGatewayAuth   IncomeGatewayAuth `json:"incomeGatewayAuth" form:"incomeGatewayAuth"`
	Confirmed           bool              `json:"confirmed" form:"confirmed" gorm:"column:confirmed;comment:;"`
	NotifyURL           string            `json:"-" form:"notifyURL" gorm:"column:notify_url;comment:;"`
	NotifyStatus        int               `json:"-" form:"notifyStatus" gorm:"column:notify_status;comment:;"`
	MetaData            string            `json:"-" form:"metaData" gorm:"column:meta_data;type:text;comment:;"`
	MerchantOrderId     string            `json:"-" form:"merchantOrderId" gorm:"column:merchant_order_id;comment:;"`
}

// TableName IncomeOrder 表名
func (IncomeOrder) TableName() string {
	return "dvfpay_income_order"
}
