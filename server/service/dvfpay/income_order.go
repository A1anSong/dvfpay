package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type IncomeOrderService struct {
}

// CreateIncomeOrder 创建IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService) CreateIncomeOrder(incomeOrder dvfpay.IncomeOrder) (err error) {
	err = global.GVA_DB.Create(&incomeOrder).Error
	return err
}

// DeleteIncomeOrder 删除IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService)DeleteIncomeOrder(incomeOrder dvfpay.IncomeOrder) (err error) {
	err = global.GVA_DB.Delete(&incomeOrder).Error
	return err
}

// DeleteIncomeOrderByIds 批量删除IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService)DeleteIncomeOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.IncomeOrder{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIncomeOrder 更新IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService)UpdateIncomeOrder(incomeOrder dvfpay.IncomeOrder) (err error) {
	err = global.GVA_DB.Save(&incomeOrder).Error
	return err
}

// GetIncomeOrder 根据id获取IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService)GetIncomeOrder(id uint) (err error, incomeOrder dvfpay.IncomeOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&incomeOrder).Error
	return
}

// GetIncomeOrderInfoList 分页获取IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService)GetIncomeOrderInfoList(info dvfpayReq.IncomeOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
    var incomeOrders []dvfpay.IncomeOrder
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
	err = db.Limit(limit).Offset(offset).Find(&incomeOrders).Error
	return err, incomeOrders, total
}
