// 自动生成模板ExchangeRecord
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExchangeRecord 结构体
// 如果含有time.Time 请自行import time包
type ExchangeRecord struct {
	global.GVA_MODEL
	MerchantId *uint  `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:;"`
	From       string `json:"from" form:"from" gorm:"column:from;comment:;"`
	To         string `json:"to" form:"to" gorm:"column:to;comment:;"`
	Rate       *int   `json:"rate" form:"rate" gorm:"column:rate;comment:;"`
	Fee        *int   `json:"fee" form:"fee" gorm:"column:fee;comment:;"`
}

// TableName ExchangeRecord 表名
func (ExchangeRecord) TableName() string {
	return "dvfpay_exchange_record"
}
