package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type PayoutOrderService struct {
}

// CreatePayoutOrder 创建PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService) CreatePayoutOrder(payoutOrder dvfpay.PayoutOrder) (err error) {
	err = global.GVA_DB.Create(&payoutOrder).Error
	return err
}

// DeletePayoutOrder 删除PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService)DeletePayoutOrder(payoutOrder dvfpay.PayoutOrder) (err error) {
	err = global.GVA_DB.Delete(&payoutOrder).Error
	return err
}

// DeletePayoutOrderByIds 批量删除PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService)DeletePayoutOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.PayoutOrder{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePayoutOrder 更新PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService)UpdatePayoutOrder(payoutOrder dvfpay.PayoutOrder) (err error) {
	err = global.GVA_DB.Save(&payoutOrder).Error
	return err
}

// GetPayoutOrder 根据id获取PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService)GetPayoutOrder(id uint) (err error, payoutOrder dvfpay.PayoutOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&payoutOrder).Error
	return
}

// GetPayoutOrderInfoList 分页获取PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService)GetPayoutOrderInfoList(info dvfpayReq.PayoutOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutOrder{})
    var payoutOrders []dvfpay.PayoutOrder
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.OrderId != "" {
        db = db.Where("order_id LIKE ?","%"+ info.OrderId+"%")
    }
    if info.Amount != nil {
        db = db.Where("amount = ?",info.Amount)
    }
    if info.Currency != "" {
        db = db.Where("currency = ?",info.Currency)
    }
    if info.Status != "" {
        db = db.Where("status = ?",info.Status)
    }
    if info.Payer != "" {
        db = db.Where("payer LIKE ?","%"+ info.Payer+"%")
    }
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&payoutOrders).Error
	return err, payoutOrders, total
}
