package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
)

type TransactionService struct {
}

// CreateTransaction 创建Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) CreateTransaction(transaction dvfpay.Transaction) (err error) {
	err = global.GVA_DB.Create(&transaction).Error
	return err
}

// DeleteTransaction 删除Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) DeleteTransaction(transaction dvfpay.Transaction) (err error) {
	err = global.GVA_DB.Delete(&transaction).Error
	return err
}

// DeleteTransactionByIds 批量删除Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) DeleteTransactionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.Transaction{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTransaction 更新Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) UpdateTransaction(transaction dvfpay.Transaction) (err error) {
	err = global.GVA_DB.Save(&transaction).Error
	return err
}

// GetTransaction 根据id获取Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) GetTransaction(id uint) (err error, transaction dvfpay.Transaction) {
	err = global.GVA_DB.Where("id = ?", id).First(&transaction).Error
	return
}

// GetTransactionInfoList 分页获取Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) GetTransactionInfoList(info dvfpayReq.TransactionSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.Transaction{})
	var transactions []dvfpay.Transaction
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.Operation != "" {
		db = db.Where("type = ?", info.Operation)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&transactions).Error
	return err, transactions, total
}
