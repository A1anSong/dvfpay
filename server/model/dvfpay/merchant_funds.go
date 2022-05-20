// 自动生成模板MerchantFunds
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// MerchantFunds 结构体
// 如果含有time.Time 请自行import time包
type MerchantFunds struct {
	global.GVA_MODEL
	MerchantId  *uint          `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Merchant    system.SysUser `json:"merchant" gorm:"foreignKey:MerchantId"`
	Currency    string         `json:"currency" form:"currency" gorm:"column:currency;comment:;"`
	Available   *int           `json:"available" form:"available" gorm:"column:available;comment:;"`
	Unavailable *int           `json:"unavailable" form:"unavailable" gorm:"column:unavailable;comment:;"`
}

// TableName MerchantFunds 表名
func (MerchantFunds) TableName() string {
	return "dvfpay_merchant_funds"
}
