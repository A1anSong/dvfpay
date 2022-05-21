package dvfpay

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
	"gorm.io/gorm"
)

type ExchangeRecordService struct {
}

// CreateExchangeRecord 创建ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) CreateExchangeRecord(exchangeRecord dvfpay.ExchangeRecord, merchantID uint) (err error) {
	//err = global.GVA_DB.Create(&exchangeRecord).Error
	//request获取的是from, to, fromAmount的值，merchantID在gin的上下文context中获取了并已经传递过来
	//在数据库中检查对应merchantID的钱包余额
	//获取对应的exchangeRate余额并开始计算
	//开始事务
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//检查商户资金
		var merchantFundsFrom dvfpay.MerchantFunds
		if err := tx.Model(&dvfpay.MerchantFunds{}).Where("merchant_id = ? and currency = ?", merchantID, exchangeRecord.From).First(&merchantFundsFrom).Error; err != nil {
			return err
		}
		if *exchangeRecord.FromAmount > *merchantFundsFrom.Available {
			return errors.New("余额不足，兑换失败！")
		}

		//检查兑换币种
		var rate string
		if exchangeRecord.To == "USDT" {
			rate = exchangeRecord.From
		} else if exchangeRecord.From == "USDT" {
			rate = exchangeRecord.To
		}
		if rate == "" {
			return errors.New("兑换币种错误！")
		}
		var exchangeRate dvfpay.ExchangeRate
		if err := tx.Model(&dvfpay.ExchangeRate{}).Where("name = ?", rate).First(&exchangeRate).Error; err != nil {
			return err
		}

		//开始计算
		var toAmount int
		var fee int
		feeRate := 992
		if exchangeRecord.To == "USDT" {
			toAmount = (*exchangeRecord.FromAmount) * (*exchangeRate.Rate) / 10000
		} else if exchangeRecord.From == "USDT" {
			toAmount = (*exchangeRecord.FromAmount) * 10000 / (*exchangeRate.Rate)
		}
		toAmount = toAmount * feeRate / 1000
		fee = (*exchangeRecord.FromAmount) * (1000 - feeRate) / 1000

		//操作对应的资金表
		var merchantFundsTo dvfpay.MerchantFunds
		available := 0
		unavailable := 0
		if err := tx.Where(dvfpay.MerchantFunds{MerchantId: &merchantID, Currency: exchangeRecord.To}).Attrs(dvfpay.MerchantFunds{Available: &available, Unavailable: &unavailable}).FirstOrCreate(&merchantFundsTo).Error; err != nil {
			return nil
		}
		*merchantFundsFrom.Available = *merchantFundsFrom.Available - *exchangeRecord.FromAmount
		*merchantFundsTo.Available = *merchantFundsTo.Available + toAmount
		if err := tx.Save(&merchantFundsFrom).Error; err != nil {
			return err
		}
		if err := tx.Save(&merchantFundsTo).Error; err != nil {
			return err
		}

		//分别给merchantId, toAmount, rate, fee赋值并存储
		exchangeRecord.MerchantId = &merchantID
		exchangeRecord.ToAmount = &toAmount
		exchangeRecord.Rate = exchangeRate.Rate
		exchangeRecord.Fee = &fee
		if err := tx.Create(&exchangeRecord).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// DeleteExchangeRecord 删除ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) DeleteExchangeRecord(exchangeRecord dvfpay.ExchangeRecord) (err error) {
	err = global.GVA_DB.Delete(&exchangeRecord).Error
	return err
}

// DeleteExchangeRecordByIds 批量删除ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) DeleteExchangeRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.ExchangeRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExchangeRecord 更新ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) UpdateExchangeRecord(exchangeRecord dvfpay.ExchangeRecord) (err error) {
	err = global.GVA_DB.Save(&exchangeRecord).Error
	return err
}

// GetExchangeRecord 根据id获取ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) GetExchangeRecord(id uint) (err error, exchangeRecord dvfpay.ExchangeRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&exchangeRecord).Error
	return
}

// GetExchangeRecordInfoList 分页获取ExchangeRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangeRecordService *ExchangeRecordService) GetExchangeRecordInfoList(info dvfpayReq.ExchangeRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.ExchangeRecord{})
	var exchangeRecords []dvfpay.ExchangeRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.From != "" {
		db = db.Where("from = ?", info.From)
	}
	if info.To != "" {
		db = db.Where("to = ?", info.To)
	}
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Order("arrival_time desc").Find(&exchangeRecords).Error
	return err, exchangeRecords, total
}

func (exchangeRecordService *ExchangeRecordService) GetMerchantExchangeRecordInfoList(info dvfpayReq.ExchangeRecordSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.ExchangeRecord{}).Where("merchant_id = ?", merchantID)
	var exchangeRecords []dvfpay.ExchangeRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.From != "" {
		db = db.Where("from = ?", info.From)
	}
	if info.To != "" {
		db = db.Where("to = ?", info.To)
	}
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&exchangeRecords).Error
	return err, exchangeRecords, total
}
