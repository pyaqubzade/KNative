package errhandler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case validator.ValidationErrors:
		ctx.Status(fiber.StatusBadRequest)
		_ = ctx.JSON(e.Error())
	case *HTTPClientError:
		//TODO: Change status code to 502
		ctx.Status(fiber.StatusInternalServerError)
		_ = ctx.JSON(e)
	default:
		_ = ctx.Status(fiber.StatusInternalServerError).
			SendString(http.StatusText(fiber.StatusInternalServerError))
	}
	return nil
}
