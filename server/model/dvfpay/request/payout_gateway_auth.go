package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
)

type PayoutGatewayAuthSearch struct {
	dvfpay.PayoutGatewayAuth
	request.PageInfo
}
