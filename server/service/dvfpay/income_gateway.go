package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type IncomeGatewayService struct {
}

// CreateIncomeGateway 创建IncomeGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayService *IncomeGatewayService) CreateIncomeGateway(incomeGateway dvfpay.IncomeGateway) (err error) {
	err = global.GVA_DB.Create(&incomeGateway).Error
	return err
}

// DeleteIncomeGateway 删除IncomeGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayService *IncomeGatewayService) DeleteIncomeGateway(incomeGateway dvfpay.IncomeGateway) (err error) {
	err = global.GVA_DB.Delete(&incomeGateway).Error
	return err
}

// DeleteIncomeGatewayByIds 批量删除IncomeGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayService *IncomeGatewayService) DeleteIncomeGatewayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.IncomeGateway{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIncomeGateway 更新IncomeGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayService *IncomeGatewayService) UpdateIncomeGateway(incomeGateway dvfpay.IncomeGateway) (err error) {
	err = global.GVA_DB.Save(&incomeGateway).Error
	return err
}

// GetIncomeGateway 根据id获取IncomeGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayService *IncomeGatewayService) GetIncomeGateway(id uint) (err error, incomeGateway dvfpay.IncomeGateway) {
	err = global.GVA_DB.Where("id = ?", id).First(&incomeGateway).Error
	return
}

// GetIncomeGatewayInfoList 分页获取IncomeGateway记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeGatewayService *IncomeGatewayService) GetIncomeGatewayInfoList(info dvfpayReq.IncomeGatewaySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeGateway{})
	var incomeGateways []dvfpay.IncomeGateway
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Parameter != "" {
		db = db.Where("parameter LIKE ?", "%"+info.Parameter+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&incomeGateways).Error
	return err, incomeGateways, total
}
