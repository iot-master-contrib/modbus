package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"go.bug.st/serial"
	"modbus/connect"
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

// @Summary 导出串口
// @Schemes
// @Description 导出串口
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /serial/export [get]
func noopSerialExport() {}

// @Summary 导入串口
// @Schemes
// @Description 导入串口
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回串口数量
// @Router /serial/import [post]
func noopSerialImport() {}

// @Summary 串口列表
// @Schemes
// @Description 串口列表
// @Tags serial
// @Produce json
// @Success 200 {object} ReplyData[string] 返回串口列表
// @Router /serial/ports [get]
func noopSerialPorts() {}

func serialRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Serial]())

	app.POST("/search", curd.ApiSearch[model.Serial]())

	app.GET("/list", curd.ApiList[model.Serial]())

	app.POST("/create", curd.ApiCreate[model.Serial](curd.GenerateRandomId[model.Serial](8), func(value *model.Serial) error {
		return connect.LoadSerial(value)
	}))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Serial]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Serial](nil, func(value *model.Serial) error {
		c := connect.GetSerial(value.Id)
		err := c.Close()
		if err != nil {
			log.Error(err)
		}
		return connect.LoadSerial(value)
	},
		"name", "desc", "heartbeat", "period", "interval", "protocol", "protocol_ops", "port", "retry", "options", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Serial](nil, func(value interface{}) error {
		id := value.(string)
		c := connect.GetSerial(id)
		return c.Close()
	}))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Serial](true, nil, func(value interface{}) error {
		id := value.(string)
		c := connect.GetSerial(id)
		return c.Close()
	}))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Serial](false, nil, func(value interface{}) error {
		id := value.(string)
		var m model.Serial
		has, err := db.Engine.ID(id).Get(&m)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("找不到 %s", id)
		}
		return connect.LoadSerial(&m)
	}))

	app.GET("/export", curd.ApiExport[model.Serial]("serial"))
	app.POST("/import", curd.ApiImport[model.Serial]())

	app.GET("ports", func(ctx *gin.Context) {
		list, err := serial.GetPortsList()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, list)
	})

}
