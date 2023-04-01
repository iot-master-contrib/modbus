package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"modbus/model"
)

// @Summary 查询产品数量
// @Schemes
// @Description 查询产品数量
// @Tags product
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回产品数量
// @Router /product/count [post]
func noopProductCount() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Product] 返回产品信息
// @Router /product/search [post]
func noopProductSearch() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[model.Product] 返回产品信息
// @Router /product/list [get]
func noopProductList() {}

// @Summary 创建产品
// @Schemes
// @Description 创建产品
// @Tags product
// @Param search body model.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Product] 返回产品信息
// @Router /product/create [post]
func noopProductCreate() {}

// @Summary 修改产品
// @Schemes
// @Description 修改产品
// @Tags product
// @Param id path int true "产品ID"
// @Param product body model.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Product] 返回产品信息
// @Router /product/{id} [post]
func noopProductUpdate() {}

// @Summary 获取产品
// @Schemes
// @Description 获取产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Product] 返回产品信息
// @Router /product/{id} [get]
func noopProductGet() {}

// @Summary 删除产品
// @Schemes
// @Description 删除产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[model.Product] 返回产品信息
// @Router /product/{id}/delete [get]
func noopProductDelete() {}

func productRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Product]())
	app.POST("/search", curd.ApiSearch[model.Product]())
	app.GET("/list", curd.ApiList[model.Product]())
	app.POST("/create", curd.ApiCreate[model.Product](curd.GenerateRandomKey(8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Product]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Product](nil, nil,
		"id", "name", "desc", "mappers", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Product](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[model.Product](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[model.Product](false, nil, nil))

}
