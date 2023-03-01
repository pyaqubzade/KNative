package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestLiveness(t *testing.T) {
	app := fiber.New()
	req := httptest.NewRequest("GET", "/health/liveness", nil)
	NewHealthHandler(app)
	res, _ := app.Test(req, 1)
	assert.Equal(t, 200, res.StatusCode)
}

func TestReadiness(t *testing.T) {
	app := fiber.New()
	req := httptest.NewRequest("GET", "/health/readiness", nil)
	NewHealthHandler(app)
	res, _ := app.Test(req, 1)
	assert.Equal(t, 200, res.StatusCode)
}
