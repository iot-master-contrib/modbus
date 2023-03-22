package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
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

func clientRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Client]())

	app.POST("/search", curd.ApiSearch[model.Client]())

	app.GET("/list", curd.ApiList[model.Client]())
	app.POST("/create", curd.ApiCreate[model.Client](curd.GenerateRandomKey(8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Client]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Client](nil, nil,
		"name", "desc", "heartbeat", "period", "interval", "disabled", "retry", "net", "addr", "port"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Client](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Client](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Client](false, nil, nil))

}
