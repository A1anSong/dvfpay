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
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.MetaData != "" {
		db = db.Where("meta_data LIKE ?", "%"+info.MetaData+"%")
	}
	if info.ArrivalTime != nil {
		db = db.Where("date(arrival_time) = date(?)", info.ArrivalTime)
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
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.MetaData != "" {
		db = db.Where("meta_data LIKE ?", "%"+info.MetaData+"%")
	}
	if info.ArrivalTime != nil {
		db = db.Where("date(arrival_time) = date(?)", info.ArrivalTime)
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

func (incomeOrderService *IncomeOrderService) GetTrendsCountIncomeOrder() (err error, list interface{}) {
	type CountIncomeOrder struct {
		Count    int    `json:"count"`
		Date     string `json:"date"`
		Currency string `json:"currency"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var countIncomeOrders []CountIncomeOrder
	err = global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Select("count(*) as count, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&countIncomeOrders).Error
	countList := make(map[string][]CountIncomeOrder)
	for _, countIncomeOrder := range countIncomeOrders {
		if _, ok := countList[countIncomeOrder.Currency]; ok {
			countList[countIncomeOrder.Currency] = append(countList[countIncomeOrder.Currency], countIncomeOrder)
		} else {
			countList[countIncomeOrder.Currency] = []CountIncomeOrder{countIncomeOrder}
		}
	}
	return err, countList
}

func (incomeOrderService *IncomeOrderService) GetTrendsSumIncomeOrder() (err error, list interface{}) {
	type SumIncomeOrder struct {
		Sum      int    `json:"sum"`
		Date     string `json:"date"`
		Currency string `json:"currency"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var sumIncomeOrders []SumIncomeOrder
	err = global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Select("sum(amount) as sum, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&sumIncomeOrders).Error
	sumList := make(map[string][]SumIncomeOrder)
	for _, sumIncomeOrder := range sumIncomeOrders {
		if _, ok := sumList[sumIncomeOrder.Currency]; ok {
			sumList[sumIncomeOrder.Currency] = append(sumList[sumIncomeOrder.Currency], sumIncomeOrder)
		} else {
			sumList[sumIncomeOrder.Currency] = []SumIncomeOrder{sumIncomeOrder}
		}
	}
	return err, sumList
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

func (incomeOrderService *IncomeOrderService) GetMerchantTrendsCountIncomeOrder(merchantID uint) (err error, list interface{}) {
	type CountIncomeOrder struct {
		Count    int    `json:"count"`
		Date     string `json:"date"`
		Currency string `json:"currency"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var countIncomeOrders []CountIncomeOrder
	err = global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID).Select("count(*) as count, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&countIncomeOrders).Error
	countList := make(map[string][]CountIncomeOrder)
	for _, countIncomeOrder := range countIncomeOrders {
		if _, ok := countList[countIncomeOrder.Currency]; ok {
			countList[countIncomeOrder.Currency] = append(countList[countIncomeOrder.Currency], countIncomeOrder)
		} else {
			countList[countIncomeOrder.Currency] = []CountIncomeOrder{countIncomeOrder}
		}
	}
	return err, countList
}

func (incomeOrderService *IncomeOrderService) GetMerchantTrendsSumIncomeOrder(merchantID uint) (err error, list interface{}) {
	type SumIncomeOrder struct {
		Sum      int    `json:"sum"`
		Date     string `json:"date"`
		Currency string `json:"currency"`
	}
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime <= endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var sumIncomeOrders []SumIncomeOrder
	err = global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID).Select("sum(amount) as sum, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&sumIncomeOrders).Error
	sumList := make(map[string][]SumIncomeOrder)
	for _, sumIncomeOrder := range sumIncomeOrders {
		if _, ok := sumList[sumIncomeOrder.Currency]; ok {
			sumList[sumIncomeOrder.Currency] = append(sumList[sumIncomeOrder.Currency], sumIncomeOrder)
		} else {
			sumList[sumIncomeOrder.Currency] = []SumIncomeOrder{sumIncomeOrder}
		}
	}
	return err, sumList
}
