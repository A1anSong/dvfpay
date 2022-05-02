// 自动生成模板TransactionRecord
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TransactionRecord 结构体
// 如果含有time.Time 请自行import time包
type TransactionRecord struct {
	global.GVA_MODEL
	MerchantId *uint  `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;"`
	Currency   string `json:"currency" form:"currency" gorm:"column:currency;comment:;"`
	Amount     *int   `json:"amount" form:"amount" gorm:"column:amount;comment:;"`
	AdminId    *uint  `json:"adminId" form:"adminId" gorm:"column:admin_id;comment:;"`
	Operation  string `json:"operation" form:"operation" gorm:"column:operation;comment:;"`
}

// TableName TransactionRecord 表名
func (TransactionRecord) TableName() string {
	return "dvfpay_transaction_record"
}
