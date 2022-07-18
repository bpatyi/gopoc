package controllers

import "github.com/gin-gonic/gin"

func GinRoute(router *gin.Engine) {
	handler := newGinHandler()

	group := router.Group("/api")
	group.GET("/", handler.getAll)
	group.POST("/", handler.create)
	group.GET("/:id", handler.getByID)
	group.PUT("/:id", handler.update)
	group.PATCH("/:id", handler.partialUpdate)
	group.DELETE("/:id", handler.delete)
}

type ginHandler struct{}

func newGinHandler() *ginHandler {
	return &ginHandler{}
}

func (g *ginHandler) getAll(context *gin.Context) {
	return
}

func (g *ginHandler) getByID(context *gin.Context) {

}

func (g *ginHandler) create(context *gin.Context) {

}

func (g *ginHandler) update(context *gin.Context) {

}

func (g *ginHandler) partialUpdate(context *gin.Context) {

}

func (g *ginHandler) delete(context *gin.Context) {

}
