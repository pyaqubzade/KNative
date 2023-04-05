package client

import (
	"github.com/pyaqubzade/knative/config"
	"github.com/pyaqubzade/knative/model"
	"github.com/pyaqubzade/knative/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Client is interface representing incidents client
type Client interface {
	GetBye(ctx *fiber.Ctx) (*model.Data, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetBye(ctx *fiber.Ctx) (*model.Data, error) {
	logger := util.GetLogger(ctx)
	logger.Debug("ActionLog.GetToken.start")

	response := new(model.Data)
	endpoint := "http://knative-go-hello.default.svc.cluster.local/api/v1/public/knative/hello"
	//endpoint := "http://knative-go-hello.default.knative.live/api/v1/public/knative/hello"
	request, err := http.NewRequest(fiber.MethodGet, endpoint, nil)
	if err != nil {
		logger.Error("Couldn't create a new request GetToken")
		return nil, err
	}

	err = util.SendRequest(ctx, request, &response, nil, "GetToken", config.DefaultRequestTimeout)
	if err != nil {
		logger.Error("GetToken request failed: ", err.Error())
		return nil, err
	}

	logger.Debug("ActionLog.GetToken.end")
	return response, nil
}
