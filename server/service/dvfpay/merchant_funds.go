package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type MerchantFundsService struct {
}

// CreateMerchantFunds 创建MerchantFunds记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantFundsService *MerchantFundsService) CreateMerchantFunds(merchantFunds dvfpay.MerchantFunds) (err error) {
	err = global.GVA_DB.Create(&merchantFunds).Error
	return err
}

// DeleteMerchantFunds 删除MerchantFunds记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantFundsService *MerchantFundsService) DeleteMerchantFunds(merchantFunds dvfpay.MerchantFunds) (err error) {
	err = global.GVA_DB.Delete(&merchantFunds).Error
	return err
}

// DeleteMerchantFundsByIds 批量删除MerchantFunds记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantFundsService *MerchantFundsService) DeleteMerchantFundsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.MerchantFunds{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMerchantFunds 更新MerchantFunds记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantFundsService *MerchantFundsService) UpdateMerchantFunds(merchantFunds dvfpay.MerchantFunds) (err error) {
	err = global.GVA_DB.Save(&merchantFunds).Error
	return err
}

// GetMerchantFunds 根据id获取MerchantFunds记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantFundsService *MerchantFundsService) GetMerchantFunds(id uint) (err error, merchantFunds dvfpay.MerchantFunds) {
	err = global.GVA_DB.Where("id = ?", id).First(&merchantFunds).Error
	return
}

// GetMerchantFundsInfoList 分页获取MerchantFunds记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantFundsService *MerchantFundsService) GetMerchantFundsInfoList(info dvfpayReq.MerchantFundsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.MerchantFunds{})
	var merchantFundss []dvfpay.MerchantFunds
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Find(&merchantFundss).Error
	return err, merchantFundss, total
}

func (merchantFundsService *MerchantFundsService) GetStatisticsMerchantFundsList() (err error, list interface{}) {
	type Funds struct {
		Currency    string `json:"currency"`
		Available   int    `json:"available"`
		Unavailable int    `json:"unavailable"`
	}
	db := global.GVA_DB.Model(&dvfpay.MerchantFunds{})
	var fundss []Funds
	if err != nil {
		return
	}
	err = db.Select("currency, sum(available) as available, sum(unavailable) as unavailable").Group("currency").Find(&fundss).Error
	return err, fundss
}

func (merchantFundsService *MerchantFundsService) GetSelfMerchantFundsInfoList(info dvfpayReq.MerchantFundsSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.MerchantFunds{}).Where("merchant_id = ?", merchantID)
	var merchantFundss []dvfpay.MerchantFunds
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&merchantFundss).Error
	return err, merchantFundss, total
}
