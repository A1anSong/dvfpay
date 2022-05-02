// 自动生成模板Transaction
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Transaction 结构体
// 如果含有time.Time 请自行import time包
type Transaction struct {
	global.GVA_MODEL
	MerchantId *uint  `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;"`
	Currency   string `json:"currency" form:"currency" gorm:"column:currency;comment:;"`
	Amount     *int   `json:"amount" form:"amount" gorm:"column:amount;comment:;"`
	Status     string `json:"status" form:"status" gorm:"column:status;comment:;"`
}

// TableName Transaction 表名
func (Transaction) TableName() string {
	return "dvfpay_transaction"
}
