package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type ExchangeRecordService struct {
}

// CreateExchangeRecord 创建ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) CreateExchangeRecord(exchangeRecord dvfpay.ExchangeRecord) (err error) {
	err = global.GVA_DB.Create(&exchangeRecord).Error
	return err
}

// DeleteExchangeRecord 删除ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService)DeleteExchangeRecord(exchangeRecord dvfpay.ExchangeRecord) (err error) {
	err = global.GVA_DB.Delete(&exchangeRecord).Error
	return err
}

// DeleteExchangeRecordByIds 批量删除ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService)DeleteExchangeRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.ExchangeRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateExchangeRecord 更新ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService)UpdateExchangeRecord(exchangeRecord dvfpay.ExchangeRecord) (err error) {
	err = global.GVA_DB.Save(&exchangeRecord).Error
	return err
}

// GetExchangeRecord 根据id获取ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService)GetExchangeRecord(id uint) (err error, exchangeRecord dvfpay.ExchangeRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&exchangeRecord).Error
	return
}

// GetExchangeRecordInfoList 分页获取ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService)GetExchangeRecordInfoList(info dvfpayReq.ExchangeRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dvfpay.ExchangeRecord{})
    var exchangeRecords []dvfpay.ExchangeRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.From != "" {
        db = db.Where("from = ?",info.From)
    }
    if info.To != "" {
        db = db.Where("to = ?",info.To)
    }
    if info.MerchantId != nil {
        db = db.Where("merchant_id = ?",info.MerchantId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&exchangeRecords).Error
	return err, exchangeRecords, total
}
