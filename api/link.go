package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"modbus/model"
)

// @Summary 查询连接数量
// @Schemes
// @Description 查询连接数量
// @Tags link
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回连接数量
// @Router /link/count [post]
func noopLinkCount() {}

// @Summary 查询连接
// @Schemes
// @Description 查询连接
// @Tags link
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Link] 返回连接信息
// @Router /link/search [post]
func noopLinkSearch() {}

// @Summary 查询连接
// @Schemes
// @Description 查询连接
// @Tags link
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Link] 返回连接信息
// @Router /link/list [get]
func noopLinkList() {}

// @Summary 修改连接
// @Schemes
// @Description 修改连接
// @Tags link
// @Param id path int true "连接ID"
// @Param link body model.Link true "连接信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Link] 返回连接信息
// @Router /link/{id} [post]
func noopLinkUpdate() {}

// @Summary 获取连接
// @Schemes
// @Description 获取连接
// @Tags link
// @Param id path int true "连接ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Link] 返回连接信息
// @Router /link/{id} [get]
func noopLinkGet() {}

// @Summary 删除连接
// @Schemes
// @Description 删除连接
// @Tags link
// @Param id path int true "连接ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Link] 返回连接信息
// @Router /link/{id}/delete [get]
func noopLinkDelete() {}

// @Summary 启用连接
// @Schemes
// @Description 启用连接
// @Tags link
// @Param id path int true "连接ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Link] 返回连接信息
// @Router /link/{id}/enable [get]
func noopLinkEnable() {}

// @Summary 禁用连接
// @Schemes
// @Description 禁用连接
// @Tags link
// @Param id path int true "连接ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Link] 返回连接信息
// @Router /link/{id}/disable [get]
func noopLinkDisable() {}

// @Summary 导出连接
// @Schemes
// @Description 导出连接
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /link/export [get]
func noopLinkExport() {}

// @Summary 导入连接
// @Schemes
// @Description 导入连接
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回连接数量
// @Router /link/import [post]
func noopLinkImport() {}

func linkRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Link]())
	app.POST("/search", curd.ApiSearch[model.Link]())
	app.GET("/list", curd.ApiList[model.Link]())
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Link]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Link](nil, nil,
		"name", "desc", "heartbeat", "poller_period", "poller_interval", "protocol_name", "protocol_options", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Link](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Link](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Link](false, nil, nil))

	app.GET("/export", curd.ApiExport[model.Link]("link"))
	app.POST("/import", curd.ApiImport[model.Link]())

}
