package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/dvfpay"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Dvfpay  dvfpay.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
