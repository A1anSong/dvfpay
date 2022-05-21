package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type TransactionRecordService struct {
}

// CreateTransactionRecord 创建TransactionRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionRecordService *TransactionRecordService) CreateTransactionRecord(transactionRecord dvfpay.TransactionRecord) (err error) {
	err = global.GVA_DB.Create(&transactionRecord).Error
	return err
}

// DeleteTransactionRecord 删除TransactionRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionRecordService *TransactionRecordService) DeleteTransactionRecord(transactionRecord dvfpay.TransactionRecord) (err error) {
	err = global.GVA_DB.Delete(&transactionRecord).Error
	return err
}

// DeleteTransactionRecordByIds 批量删除TransactionRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionRecordService *TransactionRecordService) DeleteTransactionRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.TransactionRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTransactionRecord 更新TransactionRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionRecordService *TransactionRecordService) UpdateTransactionRecord(transactionRecord dvfpay.TransactionRecord) (err error) {
	err = global.GVA_DB.Save(&transactionRecord).Error
	return err
}

// GetTransactionRecord 根据id获取TransactionRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionRecordService *TransactionRecordService) GetTransactionRecord(id uint) (err error, transactionRecord dvfpay.TransactionRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&transactionRecord).Error
	return
}

// GetTransactionRecordInfoList 分页获取TransactionRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionRecordService *TransactionRecordService) GetTransactionRecordInfoList(info dvfpayReq.TransactionRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.TransactionRecord{})
	var transactionRecords []dvfpay.TransactionRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.AdminId != nil {
		db = db.Where("admin_id = ?", info.AdminId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&transactionRecords).Error
	return err, transactionRecords, total
}
