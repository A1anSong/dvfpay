package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TransactionRecordSearch struct{
    dvfpay.TransactionRecord
    request.PageInfo
}
