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
func (payoutOrderService *PayoutOrderService) DeletePayoutOrder(payoutOrder dvfpay.PayoutOrder) (err error) {
	err = global.GVA_DB.Delete(&payoutOrder).Error
	return err
}

// DeletePayoutOrderByIds 批量删除PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService) DeletePayoutOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dvfpay.PayoutOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayoutOrder 更新PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService) UpdatePayoutOrder(payoutOrder dvfpay.PayoutOrder) (err error) {
	err = global.GVA_DB.Save(&payoutOrder).Error
	return err
}

// GetPayoutOrder 根据id获取PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService) GetPayoutOrder(id uint) (err error, payoutOrder dvfpay.PayoutOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&payoutOrder).Error
	return
}

// GetPayoutOrderInfoList 分页获取PayoutOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payoutOrderService *PayoutOrderService) GetPayoutOrderInfoList(info dvfpayReq.PayoutOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutOrder{})
	var payoutOrders []dvfpay.PayoutOrder
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
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Preload("PayoutGatewayAuth").Preload("PayoutGatewayAuth.PayoutGateway").Order("arrival_time desc").Find(&payoutOrders).Error
	return err, payoutOrders, total
}

func (payoutOrderService *PayoutOrderService) GetMerchantPayoutOrderInfoList(info dvfpayReq.PayoutOrderSearch, merchantID uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutOrder{}).Where("merchant_id = ?", merchantID)
	var payoutOrders []dvfpay.PayoutOrder
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
	err = db.Limit(limit).Offset(offset).Preload("Merchant").Preload("PayoutGatewayAuth").Preload("PayoutGatewayAuth.PayoutGateway").Order("arrival_time desc").Find(&payoutOrders).Error
	return err, payoutOrders, total
}

func (payoutOrderService *PayoutOrderService) GetStatisticsPayoutOrder() (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutOrder{})
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

func (payoutOrderService *PayoutOrderService) GetTrendsCountPayoutOrder() (err error, list interface{}) {
	type CountPayoutOrder struct {
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
	var countPayoutOrders []CountPayoutOrder
	err = global.GVA_DB.Model(&dvfpay.PayoutOrder{}).Select("count(*) as count, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&countPayoutOrders).Error
	countList := make(map[string][]CountPayoutOrder)
	for _, countPayoutOrder := range countPayoutOrders {
		if _, ok := countList[countPayoutOrder.Currency]; ok {
			countList[countPayoutOrder.Currency] = append(countList[countPayoutOrder.Currency], countPayoutOrder)
		} else {
			countList[countPayoutOrder.Currency] = []CountPayoutOrder{countPayoutOrder}
		}
	}
	return err, countList
}

func (payoutOrderService *PayoutOrderService) GetTrendsSumPayoutOrder() (err error, list interface{}) {
	type SumPayoutOrder struct {
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
	var sumPayoutOrders []SumPayoutOrder
	err = global.GVA_DB.Model(&dvfpay.PayoutOrder{}).Select("sum(amount) as sum, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&sumPayoutOrders).Error
	sumList := make(map[string][]SumPayoutOrder)
	for _, sumPayoutOrder := range sumPayoutOrders {
		if _, ok := sumList[sumPayoutOrder.Currency]; ok {
			sumList[sumPayoutOrder.Currency] = append(sumList[sumPayoutOrder.Currency], sumPayoutOrder)
		} else {
			sumList[sumPayoutOrder.Currency] = []SumPayoutOrder{sumPayoutOrder}
		}
	}
	return err, sumList
}

func (payoutOrderService *PayoutOrderService) GetMerchantStatisticsPayoutOrder(merchantID uint) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.PayoutOrder{}).Where("merchant_id = ?", merchantID)
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

func (payoutOrderService *PayoutOrderService) GetMerchantTrendsCountPayoutOrder(merchantID uint) (err error, list interface{}) {
	type CountPayoutOrder struct {
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
	var countPayoutOrders []CountPayoutOrder
	err = global.GVA_DB.Model(&dvfpay.PayoutOrder{}).Where("merchant_id = ?", merchantID).Select("count(*) as count, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&countPayoutOrders).Error
	countList := make(map[string][]CountPayoutOrder)
	for _, countPayoutOrder := range countPayoutOrders {
		if _, ok := countList[countPayoutOrder.Currency]; ok {
			countList[countPayoutOrder.Currency] = append(countList[countPayoutOrder.Currency], countPayoutOrder)
		} else {
			countList[countPayoutOrder.Currency] = []CountPayoutOrder{countPayoutOrder}
		}
	}
	return err, countList
}

func (payoutOrderService *PayoutOrderService) GetMerchantTrendsSumPayoutOrder(merchantID uint) (err error, list interface{}) {
	type SumPayoutOrder struct {
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
	var sumPayoutOrders []SumPayoutOrder
	err = global.GVA_DB.Model(&dvfpay.PayoutOrder{}).Where("merchant_id = ?", merchantID).Select("sum(amount) as sum, date(arrival_time) as date, currency").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&sumPayoutOrders).Error
	sumList := make(map[string][]SumPayoutOrder)
	for _, sumPayoutOrder := range sumPayoutOrders {
		if _, ok := sumList[sumPayoutOrder.Currency]; ok {
			sumList[sumPayoutOrder.Currency] = append(sumList[sumPayoutOrder.Currency], sumPayoutOrder)
		} else {
			sumList[sumPayoutOrder.Currency] = []SumPayoutOrder{sumPayoutOrder}
		}
	}
	return err, sumList
}
