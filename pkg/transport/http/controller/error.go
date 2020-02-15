package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	HttpCode int `json:"-"`

	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("[%s] %s", err.Code, err.Message)
}

func WrapError(httpCode int, code string, format string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		HttpCode: httpCode,
		Code:     code,
		Message:  fmt.Sprintf(format, args...),
	}
}

func EchoErrorHandler(err error, ctx echo.Context) {
	var errorResponse *ErrorResponse

	switch err.(type) {
	case *ErrorResponse:
		errorResponse = err.(*ErrorResponse)

	case *echo.HTTPError:
		httpError := err.(*echo.HTTPError)

		errorResponse = &ErrorResponse{
			HttpCode: httpError.Code,
			Code:     "internal_error",
			Message:  fmt.Sprintf("%s", httpError.Message),
		}

	default:
		errorResponse = &ErrorResponse{
			HttpCode: http.StatusInternalServerError,
			Code:     "internal_error",
			Message:  err.Error(),
		}
	}

	sErr := ctx.JSON(errorResponse.HttpCode, errorResponse)
	if sErr != nil {
		ctx.Logger().Error(sErr)
	}
}
