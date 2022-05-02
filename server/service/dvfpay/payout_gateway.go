package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type PayoutGatewayService struct {
}

// CreatePayoutGateway 创建PayoutGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayService *PayoutGatewayService) CreatePayoutGateway(payoutGateway dvfpay.PayoutGateway) (err error) {
	err = global.GVA_DB.Create(&payoutGateway).Error
	return err
}

// DeletePayoutGateway 删除PayoutGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayService *PayoutGatewayService)DeletePayoutGateway(payoutGateway dvfpay.PayoutGateway) (err error) {
	err = global.GVA_DB.Delete(&payoutGateway).Error
	return err
}

// DeletePayoutGatewayByIds 批量删除PayoutGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayService *PayoutGatewayService)DeletePayoutGatewayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.PayoutGateway{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePayoutGateway 更新PayoutGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayService *PayoutGatewayService)UpdatePayoutGateway(payoutGateway dvfpay.PayoutGateway) (err error) {
	err = global.GVA_DB.Save(&payoutGateway).Error
	return err
}

// GetPayoutGateway 根据id获取PayoutGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayService *PayoutGatewayService)GetPayoutGateway(id uint) (err error, payoutGateway dvfpay.PayoutGateway) {
	err = global.GVA_DB.Where("id = ?", id).First(&payoutGateway).Error
	return
}

// GetPayoutGatewayInfoList 分页获取PayoutGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutGatewayService *PayoutGatewayService)GetPayoutGatewayInfoList(info dvfpayReq.PayoutGatewaySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutGateway{})
    var payoutGateways []dvfpay.PayoutGateway
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Parameter != "" {
        db = db.Where("parameter LIKE ?","%"+ info.Parameter+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&payoutGateways).Error
	return err, payoutGateways, total
}
