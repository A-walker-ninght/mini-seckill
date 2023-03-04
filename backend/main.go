package main

import (
	context2 "context"
	"github.com/A-walker-ninght/mini-seckill/backend/web/controllers"
	"github.com/A-walker-ninght/mini-seckill/common"
	_interface "github.com/A-walker-ninght/mini-seckill/interface"
	"github.com/A-walker-ninght/mini-seckill/repositories"
	"github.com/A-walker-ninght/mini-seckill/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"log"
)

func Register(ctx context2.Context, service _interface.Service, tableName string, app *iris.Application) *mvc.Application {
	Party := app.Party("/" + tableName)
	p := mvc.New(Party)
	p.Register(ctx, service)
	return p
}

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	// 注册模板
	tmplate := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	// 设置静态文件
	app.HandleDir("/assets", "./backend/web/assets")

	// 出现异常跳转页面
	app.OnAnyErrorCode(func(ctx context.Context) {
		ctx.ViewData("message", "访问的页面出错")
		ctx.ViewLayout("")
		_ = ctx.View("shared/error.html")
	})

	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatalf("mysql: %s", err)
	}

	ctx, cancel := context2.WithCancel(context2.Background())
	defer func() {
		cancel()
	}()

	// 注册控制器
	productService := services.NewProductService(repositories.NewProductManager("product", db))
	p := Register(ctx, productService, "product", app)
	p.Handle(new(controllers.ProductController))

	orderService := services.NewOrderService(repositories.NewOrderManagerRepository("order", db))
	o := Register(ctx, orderService, "order", app)
	o.Handle(new(controllers.OrderController))

	// 启动服务
	app.Run(iris.Addr("localhost:8082"))
}
