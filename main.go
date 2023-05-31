package modbus

import (
	"embed"
	"encoding/json"
	"github.com/iot-master-contrib/modbus/api"
	"github.com/iot-master-contrib/modbus/connect"
	_ "github.com/iot-master-contrib/modbus/docs"
	"github.com/iot-master-contrib/modbus/internal"
	"github.com/iot-master-contrib/modbus/types"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
)

func App() *model.App {
	return &model.App{
		Id:   "modbus",
		Name: "Modbus",
		Icon: "/app/modbus/assets/modbus.svg",
		Entries: []model.AppEntry{{
			Path: "app/modbus",
			Name: "modbus",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
	}
}

//go:embed all:app/modbus
var wwwFiles embed.FS

// @title 物联大师网关接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/modbus/api/
// @query.collection.format multi
func main() {
}

func Startup(app *web.Engine) error {

	//同步表结构
	err := db.Engine.Sync2(
		new(types.Client), new(types.Server),
		new(types.Link), new(types.Serial),
		new(types.Product), new(types.Device),
	)
	if err != nil {
		log.Fatal(err)
	}

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

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/modbus/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/modbus"), "modbus")

	return nil
}

func Register() error {
	payload, _ := json.Marshal(App())
	return mqtt.Publish("master/register", payload, false, 0)
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("/app/modbus", http.FS(wwwFiles), "", "app/modbus/index.html")
}

func Shutdown() error {

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
