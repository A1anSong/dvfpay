package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type PayoutGatewayAuthService struct {
}

// CreatePayoutGatewayAuth 创建PayoutGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayAuthService *PayoutGatewayAuthService) CreatePayoutGatewayAuth(payoutGatewayAuth dvfpay.PayoutGatewayAuth) (err error) {
	err = global.GVA_DB.Create(&payoutGatewayAuth).Error
	return err
}

// DeletePayoutGatewayAuth 删除PayoutGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayAuthService *PayoutGatewayAuthService) DeletePayoutGatewayAuth(payoutGatewayAuth dvfpay.PayoutGatewayAuth) (err error) {
	err = global.GVA_DB.Delete(&payoutGatewayAuth).Error
	return err
}

// DeletePayoutGatewayAuthByIds 批量删除PayoutGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayAuthService *PayoutGatewayAuthService) DeletePayoutGatewayAuthByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.PayoutGatewayAuth{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayoutGatewayAuth 更新PayoutGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayAuthService *PayoutGatewayAuthService) UpdatePayoutGatewayAuth(payoutGatewayAuth dvfpay.PayoutGatewayAuth) (err error) {
	err = global.GVA_DB.Save(&payoutGatewayAuth).Error
	return err
}

// GetPayoutGatewayAuth 根据id获取PayoutGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayAuthService *PayoutGatewayAuthService) GetPayoutGatewayAuth(id uint) (err error, payoutGatewayAuth dvfpay.PayoutGatewayAuth) {
	err = global.GVA_DB.Where("id = ?", id).Preload("PayoutGateway").Preload("Merchants").First(&payoutGatewayAuth).Error
	return
}

// GetPayoutGatewayAuthInfoList 分页获取PayoutGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayAuthService *PayoutGatewayAuthService) GetPayoutGatewayAuthInfoList(info dvfpayReq.PayoutGatewayAuthSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutGatewayAuth{})
	var payoutGatewayAuths []dvfpay.PayoutGatewayAuth
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PayoutGatewayId != nil {
		db = db.Where("payout_gateway_id = ?", info.PayoutGatewayId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("PayoutGateway").Preload("Merchants").Find(&payoutGatewayAuths).Error
	return err, payoutGatewayAuths, total
}

func (payoutGatewayAuthService *PayoutGatewayAuthService) GetMerchantPayoutGatewayAuthInfoList(info dvfpayReq.PayoutGatewayAuthSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var authIds []uint
	global.GVA_DB.Table("dvfpay_payout_gateway_auth_merchants").Where("sys_user_id = ?", merchantID).Select("distinct payout_gateway_auth_id").Find(&authIds)
	db := global.GVA_DB.Model(&dvfpay.PayoutGatewayAuth{}).Where("deleted_at is null and id in ?", authIds)
	var payoutGatewayAuths []dvfpay.PayoutGatewayAuth
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PayoutGatewayId != nil {
		db = db.Where("payout_gateway_id = ?", info.PayoutGatewayId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("PayoutGateway").Find(&payoutGatewayAuths).Error
	return err, payoutGatewayAuths, total
}
