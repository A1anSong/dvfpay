package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExchangeRecordSearch struct{
    dvfpay.ExchangeRecord
    request.PageInfo
}
