package errhandler

type CommonError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type HTTPClientError CommonError

func (e *HTTPClientError) Error() string {
	return e.Message
}

func NewHTTPClientError(message string) error {
	return &HTTPClientError{
		Message: message,
		Code:    "client.request.failed",
	}
}

type RepositoryError CommonError

func (e *RepositoryError) Error() string {
	return e.Message
}

func NewRepositoryError(message string) error {
	return &RepositoryError{
		Message: message,
		Code:    "server.db.failed",
	}
}
