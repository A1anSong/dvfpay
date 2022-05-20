// 自动生成模板ExchangeRecord
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// ExchangeRecord 结构体
// 如果含有time.Time 请自行import time包
type ExchangeRecord struct {
	global.GVA_MODEL
	MerchantId *uint          `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Merchant   system.SysUser `json:"merchant" gorm:"foreignKey:MerchantId"`
	From       string         `json:"from" form:"from" gorm:"column:from;comment:;"`
	To         string         `json:"to" form:"to" gorm:"column:to;comment:;"`
	FromAmount *int           `json:"fromAmount" form:"fromAmount"  gorm:"column:fromAmount;comment:;"`
	ToAmount   *int           `json:"toAmount" form:"toAmount"  gorm:"column:toAmount;comment:;"`
	Rate       *int           `json:"rate" form:"rate" gorm:"column:rate;comment:;"`
	Fee        *int           `json:"fee" form:"fee" gorm:"column:fee;comment:;"`
}

// TableName ExchangeRecord 表名
func (ExchangeRecord) TableName() string {
	return "dvfpay_exchange_record"
}
