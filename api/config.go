package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"modbus/config"
)

type webOptions struct {
	Addr  string `yaml:"addr" json:"addr"`
	Debug bool   `yaml:"debug,omitempty" json:"debug,omitempty"`
	Cors  bool   `json:"cors,omitempty" json:"cors,omitempty"`
	Gzip  bool   `json:"gzip,omitempty" json:"gzip,omitempty"`
}

// @Summary 查询WEB配置
// @Schemes
// @Description 查询WEB配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[webOptions] 返回WEB配置
// @Router /config/web [get]
func configGetWeb(ctx *gin.Context) {
	curd.OK(ctx, &config.Config.Web)
}

// @Summary 修改WEB配置
// @Schemes
// @Description 修改WEB配置
// @Tags config
// @Param cfg body webOptions true "WEB配置"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int]
// @Router /config/web [post]
func configSetWeb(ctx *gin.Context) {
	var conf web.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	config.Config.Web = conf
	err = config.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

type logOutput struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size,omitempty"`
	MaxAge     int    `json:"max_age,omitempty"`
	MaxBackups int    `json:"max_backups,omitempty"`
}
type logOptions struct {
	Level  string    `json:"level"`
	Caller bool      `json:"caller,omitemptys"`
	Text   bool      `json:"text,omitempty"`
	Format string    `json:"format,omitempty"`
	Output logOutput `json:"output"`
}

// @Summary 查询日志配置
// @Schemes
// @Description 查询日志配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[logOptions] 返回日志配置
// @Router /config/log [get]
func configGetLog(ctx *gin.Context) {
	curd.OK(ctx, &config.Config.Log)
}

// @Summary 修改日志配置
// @Schemes
// @Description 修改日志配置
// @Tags config
// @Param cfg body logOptions true "日志配置"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int]
// @Router /config/log [post]
func configSetLog(ctx *gin.Context) {
	var conf log.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	config.Config.Log = conf
	err = config.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

type mqttOptions struct {
	Url      string `json:"url,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// @Summary 查询MQTT配置
// @Schemes
// @Description 查询MQTT配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[mqttOptions] 返回MQTT配置
// @Router /config/mqtt [get]
func configGetMqtt(ctx *gin.Context) {
	curd.OK(ctx, &config.Config.Mqtt)
}

// @Summary 修改MQTT配置
// @Schemes
// @Description 修改MQTT配置
// @Tags config
// @Param cfg body mqttOptions true "MQTT配置"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int]
// @Router /config/mqtt [post]
func configSetMqtt(ctx *gin.Context) {
	var conf mqtt.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	config.Config.Mqtt = conf
	err = config.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

type dbOptions struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Debug    bool   `json:"debug,omitempty"`
	LogLevel int    `json:"log_level"`
}

// @Summary 查询数据库配置
// @Schemes
// @Description 查询数据库配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[dbOptions] 返回数据库配置
// @Router /config/db [get]
func configGetDatabase(ctx *gin.Context) {
	curd.OK(ctx, &config.Config.Database)
}

// @Summary 修改数据库配置
// @Schemes
// @Description 修改数据库配置
// @Tags config
// @Param cfg body dbOptions true "数据库配置"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int]
// @Router /config/db [post]
func configSetDatabase(ctx *gin.Context) {
	var conf db.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	config.Config.Database = conf
	err = config.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func configRouter(app *gin.RouterGroup) {

	app.POST("/web", configSetWeb)
	app.GET("/web", configGetWeb)

	app.POST("/log", configSetLog)
	app.GET("/log", configGetLog)

	app.POST("/mqtt", configSetMqtt)
	app.GET("/mqtt", configGetMqtt)

	app.POST("/database", configSetDatabase)
	app.GET("/database", configGetDatabase)

}
