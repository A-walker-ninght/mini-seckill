package controllers

import (
	_interface "github.com/A-walker-ninght/mini-seckill/interface"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type OrderController struct {
	Ctx          iris.Context
	OrderService _interface.Service
}

func (o *OrderController) Get() mvc.View {
	orderArray, err := o.OrderService.GetAllInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
	}
	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orderArray,
		},
	}
}
