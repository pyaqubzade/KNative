package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pyaqubzade/knative/client"
	"github.com/pyaqubzade/knative/config"
	"github.com/pyaqubzade/knative/model"
)

type reportHandler struct {
	validate *validator.Validate
	client   client.Client
}

func NewHandler(router fiber.Router, client client.Client) {
	h := &reportHandler{validator.New(), client}
	router.Get(config.PublicRootPath+"/hello", h.SendLocalReport)
	router.Get(config.PublicRootPath+"/bye", h.SendExternalReport)
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

	//if err != nil {
	//	return err
	//}

	ctx.JSON(model.Data{Value: "Hello World"})
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
	result, err := h.client.GetBye(ctx)
	if err != nil {
		return err
	}

	ctx.JSON(result)
	return nil
}
