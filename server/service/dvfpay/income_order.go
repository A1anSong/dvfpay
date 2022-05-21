package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
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

type Statistics struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}
type Statisticians []Statistics

func (s Statisticians) Len() int { return len(s) }
func (s Statisticians) Less(i, j int) bool {
	statisticiansMap := map[string]int{
		"成功":  0,
		"处理中": 1,
	}
	return statisticiansMap[s[i].Name] < statisticiansMap[s[j].Name]
}
func (s Statisticians) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (incomeOrderService *IncomeOrderService) GetStatisticsMerchantIncomeOrder(merchantID uint) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	var statisticians []Statistics
	err = db.Select("count(*) as value, status as name").Group("status").Find(&statisticians).Error
	if err == nil {
		for i, _ := range statisticians {
			if statisticians[i].Name == "success" {
				statisticians[i].Name = "成功"
			}
			if statisticians[i].Name == "PENDING" {
				statisticians[i].Name = "处理中"
			}
		}
		sort.Sort(Statisticians(statisticians))
	}
	return err, statisticians
}

func (incomeOrderService *IncomeOrderService) GetTrendsMerchantIncomeOrder(merchantID uint) (err error, sumList interface{}, countList interface{}) {
	type SumIncomeOrder struct {
		Sum  int    `json:"sum"`
		Date string `json:"date"`
	}
	type CountIncomeOrder struct {
		Count int    `json:"count"`
		Date  string `json:"date"`
	}
	type Result struct {
		Sum   int
		Count int
		Date  string
	}
	var sumIncomeOrders []SumIncomeOrder
	var countIncomeOrders []CountIncomeOrder
	endTime := time.Now().Unix()
	loopTime := endTime - 86400*30
	var dateList []string
	for loopTime < endTime {
		dateList = append(dateList, time.Unix(loopTime, 0).Format("2006-01-02"))
		loopTime += 86400
	}
	var results []Result
	db := global.GVA_DB.Model(&dvfpay.IncomeOrder{}).Where("merchant_id = ?", merchantID)
	err = db.Select("coalesce(sum(amount), 0) as sum, count(*) as count, date(arrival_time) as date").Where("date(arrival_time) in ?", dateList).Group("date(arrival_time)").Find(&results).Error
	for _, result := range results {
		sumIncomeOrder := SumIncomeOrder{
			Sum:  result.Sum,
			Date: result.Date,
		}
		countIncomeOrder := CountIncomeOrder{
			Count: result.Count,
			Date:  result.Date,
		}
		sumIncomeOrders = append(sumIncomeOrders, sumIncomeOrder)
		countIncomeOrders = append(countIncomeOrders, countIncomeOrder)
	}
	return err, sumIncomeOrders, countIncomeOrders
}
