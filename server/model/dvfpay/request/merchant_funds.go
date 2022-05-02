package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MerchantFundsSearch struct{
    dvfpay.MerchantFunds
    request.PageInfo
}
