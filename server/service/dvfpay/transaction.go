package dvfpay

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
	"gorm.io/gorm"
)

type TransactionService struct {
}

// CreateTransaction 创建Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) CreateTransaction(transaction dvfpay.Transaction, merchantID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//如果是提现操作，同时检查available资金余额，并把相应available资金归为pending
		if transaction.Operation == "WITHDRAW" {
			//检查商户资金
			var merchantFunds dvfpay.MerchantFunds
			if err := tx.Model(&dvfpay.MerchantFunds{}).Where("merchant_id = ? and currency = 'USDT'", merchantID).First(&merchantFunds).Error; err != nil {
				return err
			}
			if *transaction.Amount > *merchantFunds.Available {
				return errors.New("余额不足，提现失败！")
			}
			*merchantFunds.Available -= *transaction.Amount
			*merchantFunds.Unavailable += *transaction.Amount
			if err := tx.Save(&merchantFunds).Error; err != nil {
				return err
			}
		}
		trans := dvfpay.Transaction{
			MerchantId: &merchantID,
			Operation:  transaction.Operation,
			Address:    transaction.Address,
			Amount:     transaction.Amount,
			Status:     "REVIEWING",
			TxID:       transaction.TxID,
		}
		if err := tx.Create(&trans).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTransaction 删除Transaction记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionService *TransactionService) DeleteTransaction(transaction dvfpay.Transaction, adminID uint) (err error) {
	//变更为拒绝放行，恢复相应资金表，并添加相应的Record记录
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&dvfpay.Transaction{}).Where("id = ?", transaction.ID).First(&transaction).Error; err != nil {
			return err
		}
		transaction.Status = "REJECTED"
		if err = tx.Save(&transaction).Error; err != nil {
			return err
		}
		var merchantFundsUSDT dvfpay.MerchantFunds
		if err := tx.Model(&dvfpay.MerchantFunds{}).Where("merchant_id = ? and currency = 'USDT'", *transaction.MerchantId).First(&merchantFundsUSDT).Error; err != nil {
			return err
		}
		*merchantFundsUSDT.Available += *transaction.Amount
		*merchantFundsUSDT.Unavailable -= *transaction.Amount
		if err = tx.Save(&merchantFundsUSDT).Error; err != nil {
			return err
		}
		transactionRecord := dvfpay.TransactionRecord{
			TransactionId: &transaction.ID,
			AdminId:       &adminID,
			Operation:     "REJECTED",
		}
		if err = tx.Save(&transactionRecord).Error; err != nil {
			return err
		}
		return nil
	})
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
func (transactionService *TransactionService) UpdateTransaction(transaction dvfpay.Transaction, adminID uint) (err error) {
	//变更为放行，操作相应资金表，并添加相应的Record记录
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		transaction.Status = "RELEASED"
		if err = tx.Save(&transaction).Error; err != nil {
			return err
		}
		var merchantFundsUSDT dvfpay.MerchantFunds
		if err := tx.Model(&dvfpay.MerchantFunds{}).Where("merchant_id = ? and currency = 'USDT'", *transaction.MerchantId).First(&merchantFundsUSDT).Error; err != nil {
			return err
		}
		*merchantFundsUSDT.Unavailable -= *transaction.Amount
		if err = tx.Save(&merchantFundsUSDT).Error; err != nil {
			return err
		}
		transactionRecord := dvfpay.TransactionRecord{
			TransactionId: &transaction.ID,
			AdminId:       &adminID,
			Operation:     "RELEASED",
		}
		if err = tx.Save(&transactionRecord).Error; err != nil {
			return err
		}
		return nil
	})
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
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Order("created_at desc").Find(&transactions).Error
	return err, transactions, total
}

func (transactionService *TransactionService) GetMerchantTransactionInfoList(info dvfpayReq.TransactionSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.Transaction{}).Where("merchant_id = ?", merchantID)
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
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&transactions).Error
	return err, transactions, total
}
