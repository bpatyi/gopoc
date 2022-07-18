package controllers

import "github.com/labstack/echo/v4"

func RouteEcho(router *echo.Echo) {
	handler := newEchoHandler()

	router.Group("/api")
	router.GET("/", handler.getAll)
	router.POST("/", handler.create)
	router.GET("/:id", handler.getByID)
	router.PUT("/:id", handler.update)
	router.PATCH("/:id", handler.updatePartial)
	router.DELETE("/:id", handler.delete)

}

type echoHandler struct {
}

func newEchoHandler() *echoHandler {
	return &echoHandler{}
}

func (e *echoHandler) getAll(c echo.Context) error {
	return nil
}

func (e *echoHandler) create(c echo.Context) error {
	return nil
}

func (e *echoHandler) getByID(c echo.Context) error {
	return nil
}

func (e *echoHandler) update(c echo.Context) error {
	return nil
}

func (e *echoHandler) updatePartial(c echo.Context) error {
	return nil
}

func (e *echoHandler) delete(c echo.Context) error {
	return nil
}
