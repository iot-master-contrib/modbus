package main

import (
	"embed"
	"encoding/json"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"modbus/api"
	"modbus/config"
	"modbus/connect"
	_ "modbus/docs"
	"modbus/internal"
	"modbus/types"
	"net/http"
)

//go:embed all:app/modbus
var wwwFiles embed.FS

// @title 物联大师网关接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/modbus/api/
// @query.collection.format multi
func main() {
	banner.Print("iot-master-plugin:alarm")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open()
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
	err = mqtt.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer mqtt.Close()

	//注册应用
	payload, _ := json.Marshal(model.App{
		Id:   "alarm",
		Name: "modbus",
		Entries: []model.AppEntry{{
			Path: "app/modbus",
			Name: "modbus",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
	})
	_ = mqtt.Publish("master/register", payload, false, 0)
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

	app := web.CreateEngine()

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/modbus/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/modbus"))

	//前端静态文件
	app.RegisterFS(http.FS(wwwFiles), "", "app/modbus/index.html")

	//监听HTTP
	app.Serve()
}
