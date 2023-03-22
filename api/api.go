package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {

	productRouter(app.Group("/product"))

	serialRouter(app.Group("/serial"))

	clientRouter(app.Group("/client"))

	serverRouter(app.Group("/server"))

	linkRouter(app.Group("/link"))

}
