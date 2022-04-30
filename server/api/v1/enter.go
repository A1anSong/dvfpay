package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	DvfpayApiGroup  dvfpay.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
