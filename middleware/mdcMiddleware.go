package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pyaqubzade/knative/model"
	"github.com/pyaqubzade/knative/util"
	log "github.com/sirupsen/logrus"
)

var headers = []string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",
	"x-ot-span-context",
	"User-Agent",
	"X-Forwarded-For",
	"Request_ID",
}

type Config struct{}

var ConfigDefault = Config{}

func configDefault(config ...Config) Config {
	return ConfigDefault
}

func NewMDC(config ...Config) fiber.Handler {
	_ = configDefault(config...)

	return func(c *fiber.Ctx) (err error) {
		requestID := c.GetReqHeaders()[model.HeaderKeyRequestID]
		operation := c.OriginalURL()
		userAgent := c.GetReqHeaders()[model.HeaderKeyUserAgent]
		userIP := c.GetReqHeaders()[model.HeaderKeyUserIP]

		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}
		fields := log.Fields{}
		addLoggerParam(fields, model.LoggerKeyRequestID, requestID)
		addLoggerParam(fields, model.LoggerKeyOperation, operation)
		addLoggerParam(fields, model.LoggerKeyUserAgent, userAgent)
		addLoggerParam(fields, model.LoggerKeyUserIP, userIP)

		logger := log.WithFields(fields)
		header := http.Header{}

		for _, v := range headers {
			header.Add(v, c.GetReqHeaders()[v])
		}

		util.WithLogger(c, logger)
		util.WithHeader(c, header)

		return c.Next()
	}
}

func addLoggerParam(fields log.Fields, field string, value string) {
	if len(value) > 0 {
		fields[field] = value
	}
}
