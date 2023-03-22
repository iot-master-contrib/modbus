package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"modbus/model"
)

// @Summary 查询串口数量
// @Schemes
// @Description 查询串口数量
// @Tags serial
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回串口数量
// @Router /serial/count [post]
func noopSerialCount() {}

// @Summary 查询串口
// @Schemes
// @Description 查询串口
// @Tags serial
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Serial] 返回串口信息
// @Router /serial/search [post]
func noopSerialSearch() {}

// @Summary 查询串口
// @Schemes
// @Description 查询串口
// @Tags serial
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Serial] 返回串口信息
// @Router /serial/list [get]
func noopSerialList() {}

// @Summary 创建串口
// @Schemes
// @Description 创建串口
// @Tags serial
// @Param search body model.Serial true "串口信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Serial] 返回串口信息
// @Router /serial/create [post]
func noopSerialCreate() {}

// @Summary 修改串口
// @Schemes
// @Description 修改串口
// @Tags serial
// @Param id path int true "串口ID"
// @Param serial body model.Serial true "串口信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Serial] 返回串口信息
// @Router /serial/{id} [post]
func noopSerialUpdate() {}

// @Summary 获取串口
// @Schemes
// @Description 获取串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Serial] 返回串口信息
// @Router /serial/{id} [get]
func noopSerialGet() {}

// @Summary 删除串口
// @Schemes
// @Description 删除串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Serial] 返回串口信息
// @Router /serial/{id}/delete [get]
func noopSerialDelete() {}

// @Summary 启用串口
// @Schemes
// @Description 启用串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Serial] 返回串口信息
// @Router /serial/{id}/enable [get]
func noopSerialEnable() {}

// @Summary 禁用串口
// @Schemes
// @Description 禁用串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Serial] 返回串口信息
// @Router /serial/{id}/disable [get]
func noopSerialDisable() {}

func serialRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Serial]())
	app.POST("/search", curd.ApiSearch[model.Serial]())
	app.GET("/list", curd.ApiList[model.Serial]())
	app.POST("/create", curd.ApiCreate[model.Serial](curd.GenerateRandomKey(8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Serial]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Serial](nil, nil,
		"name", "desc", "heartbeat", "period", "interval", "retry", "options", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Serial](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Serial](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Serial](false, nil, nil))

}
