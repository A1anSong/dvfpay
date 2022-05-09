// 自动生成模板ExchangeRate
package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExchangeRate 结构体
// 如果含有time.Time 请自行import time包
type ExchangeRate struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Rate *int   `json:"rate" form:"rate" gorm:"column:rate;comment:;"`
}

// TableName ExchangeRate 表名
func (ExchangeRate) TableName() string {
	return "dvfpay_exchange_rate"
}
