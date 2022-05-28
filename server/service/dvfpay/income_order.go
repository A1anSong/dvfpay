package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
	dvfpayUtil "github.com/flipped-aurora/gin-vue-admin/server/utils/dvfpay"
	"sort"
	"time"
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
func (incomeOrderService *IncomeOrderService) DeleteIncomeOrder(incomeOrder dvfpay.IncomeOrder) (err error) {
	err = global.GVA_DB.Delete(&incomeOrder).Error
	return err
}

// DeleteIncomeOrderByIds 批量删除IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService) DeleteIncomeOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.IncomeOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIncomeOrder 更新IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService) UpdateIncomeOrder(incomeOrder dvfpay.IncomeOrder) (err error) {
	err = global.GVA_DB.Save(&incomeOrder).Error
	return err
}

// GetIncomeOrder 根据id获取IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService) GetIncomeOrder(id uint) (err error, incomeOrder dvfpay.IncomeOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&incomeOrder).Error
	return
}

// GetIncomeOrderInfoList 分页获取IncomeOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeOrderService *IncomeOrderService) GetIncomeOrderInfoList(info dvfpayReq.IncomeOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	var incomeOrders []dvfpay.IncomeOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != "" {
		db = db.Where("order_id LIKE ?", "%"+info.OrderId+"%")
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.Payer != "" {
		db = db.Where("payer LIKE ?", "%"+info.Payer+"%")
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Preload("IncomeGatewayAuth").Preload("IncomeGatewayAuth.IncomeGateway").Order("arrival_time desc").Find(&incomeOrders).Error
	return err, incomeOrders, total
}

func (incomeOrderService *IncomeOrderService) GetMerchantIncomeOrderInfoList(info dvfpayReq.IncomeOrderSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	var incomeOrders []dvfpay.IncomeOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != "" {
		db = db.Where("order_id LIKE ?", "%"+info.OrderId+"%")
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.Payer != "" {
		db = db.Where("payer LIKE ?", "%"+info.Payer+"%")
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Preload("IncomeGatewayAuth").Preload("IncomeGatewayAuth.IncomeGateway").Order("arrival_time desc").Find(&incomeOrders).Error
	return err, incomeOrders, total
}

func (incomeOrderService *IncomeOrderService) ConfirmIncomeOrder(incomeOrder dvfpay.IncomeOrder) (err error) {
	err = global.GVA_DB.First(&incomeOrder).Error
	incomeOrder.Confirmed = true
	err = global.GVA_DB.Save(&incomeOrder).Error
	return err
}

func (incomeOrderService *IncomeOrderService) GetStatisticsIncomeOrder() (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	var statisticians []dvfpayUtil.Statistics
	err = db.Select("count(*) as value, status as name").Group("status").Find(&statisticians).Error
	if err == nil {
		for i, _ := range statisticians {
			if statisticians[i].Name == "success" {
				statisticians[i].Name = "成功"
			}
			if statisticians[i].Name == "PENDING" {
				statisticians[i].Name = "待定"
			}
			if statisticians[i].Name == "SETTLED" {
				statisticians[i].Name = "结算"
			}
		}
		sort.Sort(dvfpayUtil.Statisticians(statisticians))
	}
	return err, statisticians
}

func (incomeOrderService *IncomeOrderService) GetTrendsCountIncomeOrder() (err error, usdList interface{}, eurList interface{}, gbpList interface{}) {
	type CountIncomeOrder struct {
		Count int    `json:"count"`
		Date  string `json:"date"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var usdResults []CountIncomeOrder
	var eurResults []CountIncomeOrder
	var gbpResults []CountIncomeOrder
	usdDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	err = usdDb.Select("count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "USD").Group("date(arrival_time)").Find(&usdResults).Error
	eurDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	err = eurDb.Select("count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "EUR").Group("date(arrival_time)").Find(&eurResults).Error
	gbpDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	err = gbpDb.Select("count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "GBP").Group("date(arrival_time)").Find(&gbpResults).Error
	return err, usdResults, eurResults, gbpResults
}

func (incomeOrderService *IncomeOrderService) GetTrendsSumIncomeOrder() (err error, usdList interface{}, eurList interface{}, gbpList interface{}) {
	type SumIncomeOrder struct {
		Sum  int    `json:"sum"`
		Date string `json:"date"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var usdResults []SumIncomeOrder
	var eurResults []SumIncomeOrder
	var gbpResults []SumIncomeOrder
	usdDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	err = usdDb.Select("sum(amount) as sum, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "USD").Group("date(arrival_time)").Find(&usdResults).Error
	eurDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	err = eurDb.Select("sum(amount) as sum, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "EUR").Group("date(arrival_time)").Find(&eurResults).Error
	gbpDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{})
	err = gbpDb.Select("sum(amount) as sum, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "GBP").Group("date(arrival_time)").Find(&gbpResults).Error
	return err, usdResults, eurResults, gbpResults
}

func (incomeOrderService *IncomeOrderService) GetMerchantStatisticsIncomeOrder(merchantID uint) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	var statisticians []dvfpayUtil.Statistics
	err = db.Select("count(*) as value, status as name").Group("status").Find(&statisticians).Error
	if err == nil {
		for i, _ := range statisticians {
			if statisticians[i].Name == "success" {
				statisticians[i].Name = "成功"
			}
			if statisticians[i].Name == "PENDING" {
				statisticians[i].Name = "待定"
			}
			if statisticians[i].Name == "SETTLED" {
				statisticians[i].Name = "结算"
			}
		}
		sort.Sort(dvfpayUtil.Statisticians(statisticians))
	}
	return err, statisticians
}

func (incomeOrderService *IncomeOrderService) GetMerchantTrendsCountIncomeOrder(merchantID uint) (err error, usdList interface{}, eurList interface{}, gbpList interface{}) {
	type CountIncomeOrder struct {
		Count int    `json:"count"`
		Date  string `json:"date"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var usdResults []CountIncomeOrder
	var eurResults []CountIncomeOrder
	var gbpResults []CountIncomeOrder
	usdDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = usdDb.Select("count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "USD").Group("date(arrival_time)").Find(&usdResults).Error
	eurDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = eurDb.Select("count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "EUR").Group("date(arrival_time)").Find(&eurResults).Error
	gbpDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = gbpDb.Select("count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "GBP").Group("date(arrival_time)").Find(&gbpResults).Error
	return err, usdResults, eurResults, gbpResults
}

func (incomeOrderService *IncomeOrderService) GetMerchantTrendsSumIncomeOrder(merchantID uint) (err error, usdList interface{}, eurList interface{}, gbpList interface{}) {
	type SumIncomeOrder struct {
		Sum  int    `json:"sum"`
		Date string `json:"date"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var usdResults []SumIncomeOrder
	var eurResults []SumIncomeOrder
	var gbpResults []SumIncomeOrder
	usdDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = usdDb.Select("sum(amount) as sum, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "USD").Group("date(arrival_time)").Find(&usdResults).Error
	eurDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = eurDb.Select("sum(amount) as sum, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "EUR").Group("date(arrival_time)").Find(&eurResults).Error
	gbpDb := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = gbpDb.Select("sum(amount) as sum, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Where("currency = ?", "GBP").Group("date(arrival_time)").Find(&gbpResults).Error
	return err, usdResults, eurResults, gbpResults
}
