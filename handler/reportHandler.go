package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pyaqubzade/knative/config"
)

type reportHandler struct {
	validate *validator.Validate
}

func NewHandler(router fiber.Router) {
	h := &reportHandler{validator.New()}
	router.Post(config.PublicRootPath+"/hello", h.SendLocalReport)
	router.Post(config.PublicRootPath+"/bye", h.SendExternalReport)
}

func (h *reportHandler) SendLocalReport(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	//request := new(model.ReportRequest)
	//err := json.Unmarshal(ctx.Body(), request)
	//if err != nil {
	//	return err
	//}
	//
	//err = h.validate.Struct(request)
	//if err != nil {
	//	log.Error("ActionLog.ReportHandler.error", err.Error())
	//	return err
	//}

	resultMessage := "Hello World"
	//if err != nil {
	//	return err
	//}

	ctx.Response().SetBody([]byte(resultMessage))
	return nil
}

func (h *reportHandler) SendExternalReport(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	//request := new(model.ReportRequest)
	//err := json.Unmarshal(ctx.Body(), request)
	//if err != nil {
	//	return err
	//}
	//
	//err = h.validate.Struct(request)
	//if err != nil {
	//	log.Error("ActionLog.ReportHandler.error", err.Error())
	//	return err
	//}
	//
	resultMessage := "Bye"
	//if err != nil {
	//	return err
	//}

	ctx.Response().SetBody([]byte(resultMessage))
	return nil
}
