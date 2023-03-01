package util

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pyaqubzade/knative/model"
	"github.com/sirupsen/logrus"
)

// WithLogger add logger to context
func WithLogger(ctx *fiber.Ctx, logger *logrus.Entry) {
	ctx.Locals(model.ContextLogger, logger)
}

// WithHeader add header to context
func WithHeader(ctx *fiber.Ctx, header http.Header) {
	ctx.Locals(model.ContextHeader, header)
}

// GetLogger get logger from context
func GetLogger(ctx *fiber.Ctx) *logrus.Entry {
	logger, ok := ctx.Locals(model.ContextLogger).(*logrus.Entry)
	if !ok {
		return logrus.NewEntry(logrus.New())
	}
	return logger
}

// GetHeader get header from context
func GetHeader(ctx *fiber.Ctx) http.Header {
	header, ok := ctx.Locals(model.ContextHeader).(http.Header)
	if !ok {
		return http.Header{}
	}
	return header
}
