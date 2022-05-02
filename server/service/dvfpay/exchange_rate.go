package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type ExchangeRateService struct {
}

// CreateExchangeRate 创建ExchangeRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRateService *ExchangeRateService) CreateExchangeRate(exchangeRate dvfpay.ExchangeRate) (err error) {
	err = global.GVA_DB.Create(&exchangeRate).Error
	return err
}

// DeleteExchangeRate 删除ExchangeRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRateService *ExchangeRateService)DeleteExchangeRate(exchangeRate dvfpay.ExchangeRate) (err error) {
	err = global.GVA_DB.Delete(&exchangeRate).Error
	return err
}

// DeleteExchangeRateByIds 批量删除ExchangeRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRateService *ExchangeRateService)DeleteExchangeRateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.ExchangeRate{},"id in ?",ids.Ids).Error
	return err
}

// UpdateExchangeRate 更新ExchangeRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRateService *ExchangeRateService)UpdateExchangeRate(exchangeRate dvfpay.ExchangeRate) (err error) {
	err = global.GVA_DB.Save(&exchangeRate).Error
	return err
}

// GetExchangeRate 根据id获取ExchangeRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRateService *ExchangeRateService)GetExchangeRate(id uint) (err error, exchangeRate dvfpay.ExchangeRate) {
	err = global.GVA_DB.Where("id = ?", id).First(&exchangeRate).Error
	return
}

// GetExchangeRateInfoList 分页获取ExchangeRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRateService *ExchangeRateService)GetExchangeRateInfoList(info dvfpayReq.ExchangeRateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dvfpay.ExchangeRate{})
    var exchangeRates []dvfpay.ExchangeRate
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.From != "" {
        db = db.Where("from = ?",info.From)
    }
    if info.To != "" {
        db = db.Where("to = ?",info.To)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&exchangeRates).Error
	return err, exchangeRates, total
}
