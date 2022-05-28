// 自动生成模板Transaction
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Transaction 结构体
// 如果含有time.Time 请自行import time包
type Transaction struct {
	global.GVA_MODEL
	MerchantId *uint          `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	Merchant   system.SysUser `json:"merchant" gorm:"foreignKey:MerchantId"`
	Operation  string         `json:"operation" form:"operation" gorm:"column:operation;comment:;"`
	Amount     *int           `json:"amount" form:"amount" gorm:"column:amount;comment:;"`
	Status     string         `json:"status" form:"status" gorm:"column:status;comment:;"`
	TxID       string         `json:"txID" form:"txID" gorm:"column:txID;comment:;"`
}

// TableName Transaction 表名
func (Transaction) TableName() string {
	return "dvfpay_transaction"
}
