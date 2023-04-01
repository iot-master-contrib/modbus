package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"modbus/connect"
	"modbus/model"
)

// @Summary 查询服务器数量
// @Schemes
// @Description 查询服务器数量
// @Tags server
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回服务器数量
// @Router /server/count [post]
func noopServerCount() {}

// @Summary 查询服务器
// @Schemes
// @Description 查询服务器
// @Tags server
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Server] 返回服务器信息
// @Router /server/search [post]
func noopServerSearch() {}

// @Summary 查询服务器
// @Schemes
// @Description 查询服务器
// @Tags server
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Server] 返回服务器信息
// @Router /server/list [get]
func noopServerList() {}

// @Summary 创建服务器
// @Schemes
// @Description 创建服务器
// @Tags server
// @Param search body model.Server true "服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Server] 返回服务器信息
// @Router /server/create [post]
func noopServerCreate() {}

// @Summary 修改服务器
// @Schemes
// @Description 修改服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Param server body model.Server true "服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Server] 返回服务器信息
// @Router /server/{id} [post]
func noopServerUpdate() {}

// @Summary 获取服务器
// @Schemes
// @Description 获取服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Server] 返回服务器信息
// @Router /server/{id} [get]
func noopServerGet() {}

// @Summary 删除服务器
// @Schemes
// @Description 删除服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Server] 返回服务器信息
// @Router /server/{id}/delete [get]
func noopServerDelete() {}

// @Summary 启用服务器
// @Schemes
// @Description 启用服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Server] 返回服务器信息
// @Router /server/{id}/enable [get]
func noopServerEnable() {}

// @Summary 禁用服务器
// @Schemes
// @Description 禁用服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Server] 返回服务器信息
// @Router /server/{id}/disable [get]
func noopServerDisable() {}

func serverRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Server]())

	app.POST("/search", curd.ApiSearch[model.Server]())

	app.GET("/list", curd.ApiList[model.Server]())

	app.POST("/create", curd.ApiCreate[model.Server](curd.GenerateRandomKey(8), func(value interface{}) error {
		return connect.LoadServer(value.(*model.Server))
	}))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Server]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Server](nil, func(value interface{}) error {
		m := value.(*model.Server)
		c := connect.GetServer(m.Id)
		err := c.Close()
		if err != nil {
			return err
		}
		return connect.LoadServer(m)
	},
		"name", "desc", "heartbeat", "period", "interval", "retry", "options", "disabled", "port", "standalone"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Server](nil, func(value interface{}) error {
		id := value.(string)
		c := connect.GetServer(id)
		return c.Close()
	}))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Server](true, nil, func(value interface{}) error {
		id := value.(string)
		c := connect.GetServer(id)
		return c.Close()
	}))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Server](false, nil, func(value interface{}) error {
		id := value.(string)
		var m model.Server
		has, err := db.Engine.ID(id).Get(&m)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("找不到 %s", id)
		}
		return connect.LoadServer(&m)
	}))

}
