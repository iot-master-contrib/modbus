package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"modbus/connect"
	"modbus/model"
)

// @Summary 查询客户端数量
// @Schemes
// @Description 查询客户端数量
// @Tags client
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回客户端数量
// @Router /client/count [post]
func noopClientCount() {}

// @Summary 查询客户端
// @Schemes
// @Description 查询客户端
// @Tags client
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Client] 返回客户端信息
// @Router /client/search [post]
func noopClientSearch() {}

// @Summary 查询客户端
// @Schemes
// @Description 查询客户端
// @Tags client
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Client] 返回客户端信息
// @Router /client/list [get]
func noopClientList() {}

// @Summary 创建客户端
// @Schemes
// @Description 创建客户端
// @Tags client
// @Param search body model.Client true "客户端信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Client] 返回客户端信息
// @Router /client/create [post]
func noopClientCreate() {}

// @Summary 修改客户端
// @Schemes
// @Description 修改客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Param client body model.Client true "客户端信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Client] 返回客户端信息
// @Router /client/{id} [post]
func noopClientUpdate() {}

// @Summary 获取客户端
// @Schemes
// @Description 获取客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Client] 返回客户端信息
// @Router /client/{id} [get]
func noopClientGet() {}

// @Summary 删除客户端
// @Schemes
// @Description 删除客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Client] 返回客户端信息
// @Router /client/{id}/delete [get]
func noopClientDelete() {}

// @Summary 启用客户端
// @Schemes
// @Description 启用客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Client] 返回客户端信息
// @Router /client/{id}/enable [get]
func noopClientEnable() {}

// @Summary 禁用客户端
// @Schemes
// @Description 禁用客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Client] 返回客户端信息
// @Router /client/{id}/disable [get]
func noopClientDisable() {}

// @Summary 导出客户端
// @Schemes
// @Description 导出客户端
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /client/export [get]
func noopClientExport() {}

// @Summary 导入客户端
// @Schemes
// @Description 导入客户端
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回客户端数量
// @Router /client/import [post]
func noopClientImport() {}

func clientRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Client]())

	app.POST("/search", curd.ApiSearch[model.Client]())

	app.GET("/list", curd.ApiList[model.Client]())

	app.POST("/create", curd.ApiCreate[model.Client](curd.GenerateRandomId[model.Client](8), func(value *model.Client) error {
		return connect.LoadClient(value)
	}))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Client]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Client](nil, func(value *model.Client) error {
		c := connect.GetClient(value.Id)
		err := c.Close()
		if err != nil {
			log.Error(err)
		}
		return connect.LoadClient(value)
	},
		"name", "desc", "heartbeat", "period", "interval", "protocol", "protocol_ops", "disabled", "retry", "net", "addr", "port"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Client](nil, func(value interface{}) error {
		id := value.(string)
		c := connect.GetClient(id)
		return c.Close()
	}))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Client](true, nil, func(value interface{}) error {
		id := value.(string)
		c := connect.GetClient(id)
		return c.Close()
	}))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Client](false, nil, func(value interface{}) error {
		id := value.(string)
		var m model.Client
		has, err := db.Engine.ID(id).Get(&m)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("找不到 %s", id)
		}
		return connect.LoadClient(&m)
	}))

	app.GET("/export", curd.ApiExport[model.Client]("client"))
	app.POST("/import", curd.ApiImport[model.Client]())

}
