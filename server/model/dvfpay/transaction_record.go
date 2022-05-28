// 自动生成模板TransactionRecord
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// TransactionRecord 结构体
// 如果含有time.Time 请自行import time包
type TransactionRecord struct {
	global.GVA_MODEL
	TransactionId *uint          `json:"transactionId" form:"transactionId" gorm:"column:transaction_id;comment:;"`
	Transaction   Transaction    `json:"transaction"`
	AdminId       *uint          `json:"adminId" form:"adminId" gorm:"column:admin_id;comment:;"`
	Admin         system.SysUser `json:"admin" gorm:"foreignKey:AdminId"`
	Operation     string         `json:"operation" form:"operation" gorm:"column:operation;comment:;"`
}

// TableName TransactionRecord 表名
func (TransactionRecord) TableName() string {
	return "dvfpay_transaction_record"
}
