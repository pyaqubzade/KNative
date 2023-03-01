package handler

import "github.com/gofiber/fiber/v2"

type healthHandler struct{}

func NewHealthHandler(router fiber.Router) {
	h := &healthHandler{}
	router.Get("health/liveness", h.Health)
	router.Get("health/readiness", h.Health)
}

func (h *healthHandler) Health(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusOK)
	return nil
}
