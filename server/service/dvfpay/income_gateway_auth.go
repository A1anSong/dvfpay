package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type IncomeGatewayAuthService struct {
}

// CreateIncomeGatewayAuth 创建IncomeGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayAuthService *IncomeGatewayAuthService) CreateIncomeGatewayAuth(incomeGatewayAuth dvfpay.IncomeGatewayAuth) (err error) {
	err = global.GVA_DB.Create(&incomeGatewayAuth).Error
	return err
}

// DeleteIncomeGatewayAuth 删除IncomeGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayAuthService *IncomeGatewayAuthService) DeleteIncomeGatewayAuth(incomeGatewayAuth dvfpay.IncomeGatewayAuth) (err error) {
	err = global.GVA_DB.Delete(&incomeGatewayAuth).Error
	return err
}

// DeleteIncomeGatewayAuthByIds 批量删除IncomeGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayAuthService *IncomeGatewayAuthService) DeleteIncomeGatewayAuthByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.IncomeGatewayAuth{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIncomeGatewayAuth 更新IncomeGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayAuthService *IncomeGatewayAuthService) UpdateIncomeGatewayAuth(incomeGatewayAuth dvfpay.IncomeGatewayAuth) (err error) {
	err = global.GVA_DB.Save(&incomeGatewayAuth).Error
	return err
}

// GetIncomeGatewayAuth 根据id获取IncomeGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayAuthService *IncomeGatewayAuthService) GetIncomeGatewayAuth(id uint) (err error, incomeGatewayAuth dvfpay.IncomeGatewayAuth) {
	err = global.GVA_DB.Where("id = ?", id).Preload("IncomeGateway").Preload("Merchants").First(&incomeGatewayAuth).Error
	return
}

// GetIncomeGatewayAuthInfoList 分页获取IncomeGatewayAuth记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayAuthService *IncomeGatewayAuthService) GetIncomeGatewayAuthInfoList(info dvfpayReq.IncomeGatewayAuthSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeGatewayAuth{})
	var incomeGatewayAuths []dvfpay.IncomeGatewayAuth
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IncomeGatewayId != nil {
		db = db.Where("income_gateway_id = ?", info.IncomeGatewayId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("IncomeGateway").Preload("Merchants").Find(&incomeGatewayAuths).Error
	return err, incomeGatewayAuths, total
}

func (incomeGatewayAuthService *IncomeGatewayAuthService) GetMerchantIncomeGatewayAuthInfoList(info dvfpayReq.IncomeGatewayAuthSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var authIds []uint
	global.GVA_DB.Table("dvfpay_income_gateway_auth_merchants").Where("sys_user_id = ?", merchantID).Select("distinct income_gateway_auth_id").Find(&authIds)
	db := global.GVA_DB.Model(&dvfpay.IncomeGatewayAuth{}).Where("deleted_at is null and id in ?", authIds)
	var incomeGatewayAuths []dvfpay.IncomeGatewayAuth
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IncomeGatewayId != nil {
		db = db.Where("income_gateway_id = ?", info.IncomeGatewayId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("IncomeGateway").Find(&incomeGatewayAuths).Error
	if err == nil {
		for i, _ := range incomeGatewayAuths {
			incomeGatewayAuths[i].IncomeGateway.Parameter = ""
			*incomeGatewayAuths[i].IncomeGateway.Fee = -1
		}
	}
	return err, incomeGatewayAuths, total
}
