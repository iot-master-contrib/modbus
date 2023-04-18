package main

import (
	"embed"
	"github.com/kardianos/service"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	swaggerFiles "github.com/zgwit/swagger-files"
	"modbus/api"
	"modbus/args"
	"modbus/config"
	"modbus/connect"
	_ "modbus/docs"
	"modbus/internal"
	"modbus/types"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:embed all:www
var wwwFiles embed.FS

var serviceConfig = &service.Config{
	Name:        "iot-master-gateway",
	DisplayName: "物联大师网关",
	Description: "物联大师网关",
	Arguments:   nil,
}

// @title 物联大师网关接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @query.collection.format multi
func main() {
	args.Parse()

	//传递参数到服务
	serviceConfig.Arguments = []string{"-c", args.ConfigPath}

	// 构建服务对象
	program := &Program{}
	s, err := service.New(program, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if args.Uninstall {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("卸载服务成功")
		return
	}

	if args.Install {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("安装服务成功")
		return
	}

	err = s.Run()
	if err != nil {
		_ = logger.Error(err)
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	//log.Println("===开始服务===")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	//log.Println("===停止服务===")
	_ = shutdown()
	return nil
}

func (p *Program) run() {

	// 此处编写具体的服务代码
	hup := make(chan os.Signal, 2)
	signal.Notify(hup, syscall.SIGHUP)
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, os.Kill)

	//原本的Main函数
	originMain()

	select {
	case <-hup:
	case <-quit:
		//优雅地结束
		_ = shutdown()
		//os.Exit(0)
	}
}

func originMain() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = log.Open(config.Config.Log)
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open(config.Config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(types.Client), new(types.Server),
		new(types.Link), new(types.Serial),
		new(types.Product), new(types.Device),
	)
	if err != nil {
		log.Fatal(err)
	}

	//MQTT总线
	err = mqtt.Open(config.Config.Mqtt)
	if err != nil {
		log.Fatal(err)
	}
	defer mqtt.Close()

	//内部加载
	err = internal.LoadProducts()
	if err != nil {
		log.Fatal(err)
	}

	//连接
	err = connect.Load()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()

	//
	////加载主程序
	//err = internal.Open()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer internal.Close()

	app := web.CreateEngine(config.Config.Web)

	//注册前端接口
	api.RegisterRoutes(app.Group("/api"))

	//注册接口文档
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//前端静态文件
	web.RegisterFS(app, http.FS(wwwFiles), "www", "index.html")

	//监听HTTP
	err = app.Run(config.Config.Web.Addr)
	if err != nil {
		log.Fatal("HTTP 服务启动错误", err)
	}
}

func shutdown() error {

	//_ = database.Close()
	//_ = tsdb.Close()
	//connect.Close()
	//master.Close()

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
