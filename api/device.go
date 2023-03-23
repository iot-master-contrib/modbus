package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"modbus/model"
)

// @Summary 查询设备数量
// @Schemes
// @Description 查询设备数量
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回设备数量
// @Router /device/count [post]
func noopDeviceCount() {}

// @Summary 查询设备
// @Schemes
// @Description 查询设备
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Device] 返回设备信息
// @Router /device/search [post]
func noopDeviceSearch() {}

// @Summary 查询设备
// @Schemes
// @Description 查询设备
// @Tags device
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Device] 返回设备信息
// @Router /device/list [get]
func noopDeviceList() {}

// @Summary 创建设备
// @Schemes
// @Description 创建设备
// @Tags device
// @Param search body model.Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Device] 返回设备信息
// @Router /device/create [post]
func noopDeviceCreate() {}

// @Summary 修改设备
// @Schemes
// @Description 修改设备
// @Tags device
// @Param id path int true "设备ID"
// @Param device body model.Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Device] 返回设备信息
// @Router /device/{id} [post]
func noopDeviceUpdate() {}

// @Summary 获取设备
// @Schemes
// @Description 获取设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Device] 返回设备信息
// @Router /device/{id} [get]
func noopDeviceGet() {}

// @Summary 删除设备
// @Schemes
// @Description 删除设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Device] 返回设备信息
// @Router /device/{id}/delete [get]
func noopDeviceDelete() {}

// @Summary 启用设备
// @Schemes
// @Description 启用设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Device] 返回设备信息
// @Router /device/{id}/enable [get]
func noopDeviceEnable() {}

// @Summary 禁用设备
// @Schemes
// @Description 禁用设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Device] 返回设备信息
// @Router /device/{id}/disable [get]
func noopDeviceDisable() {}

func deviceRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Device]())
	app.POST("/search", curd.ApiSearch[model.Device]())
	app.GET("/list", curd.ApiList[model.Device]())
	app.POST("/create", curd.ApiCreate[model.Device](curd.GenerateRandomKey(8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Device]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Device](nil, nil,
		"name", "desc", "tunnel_id", "product_id", "slave", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Device](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Device](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Device](false, nil, nil))
}
