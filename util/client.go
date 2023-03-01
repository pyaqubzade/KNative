package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func IsHTTPStatus2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

func SendRequest(ctx *fiber.Ctx, request *http.Request, response interface{}, headers map[string]string,
	method string, requestTimeout time.Duration) error {

	request.Header = GetHeader(ctx)
	for key, val := range headers {
		request.Header.Del(key)
		request.Header.Add(key, val)
	}

	client := http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if !IsHTTPStatus2xx(resp.StatusCode) {
		return fmt.Errorf("client %s returned http status %s", method, resp.Status)
	}

	if response == nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &response)
	}
	return err
}

func GetRequest(ctx *fiber.Ctx, url string, response interface{}, method string, timeout time.Duration) error {
	request, err := http.NewRequest(fiber.MethodGet, url, bytes.NewBuffer(nil))
	if err != nil {
		return err
	}

	headers := map[string]string{fiber.HeaderContentType: fiber.MIMEApplicationJSON}
	err = SendRequest(ctx, request, &response, headers, method, timeout)
	if err != nil {
		return err
	}

	return nil
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return fmt.Sprintf("Basic " + base64.StdEncoding.EncodeToString([]byte(auth)))
}

func IsStatusError(status string) bool {
	return status == "error"
}
