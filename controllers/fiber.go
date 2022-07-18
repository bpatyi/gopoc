package controllers

import "github.com/gofiber/fiber/v2"

func RouteFiber(router *fiber.App) {
	handler := newFiberHandler()

	api := router.Group("/api")

	api.Get("/", handler.getAll)
	api.Post("/", handler.create)
	api.Get("/:id", handler.get)
	api.Put("/:id", handler.update)
	api.Patch("/:id", handler.updatePartial)
	api.Delete("/:id", handler.delete)
}

type fiberHandler struct {
}

func newFiberHandler() *fiberHandler {
	return &fiberHandler{}
}

func (f *fiberHandler) getAll(ctx *fiber.Ctx) error {
	return nil
}

func (f *fiberHandler) create(ctx *fiber.Ctx) error {
	return nil
}

func (f *fiberHandler) get(ctx *fiber.Ctx) error {
	return nil
}

func (f *fiberHandler) update(ctx *fiber.Ctx) error {
	return nil
}

func (f *fiberHandler) updatePartial(ctx *fiber.Ctx) error {
	return nil
}

func (f *fiberHandler) delete(ctx *fiber.Ctx) error {
	return nil
}
