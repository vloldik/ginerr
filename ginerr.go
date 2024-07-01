package ginerr

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HighLevelError struct {
	Code    int
	Message string
}

func (err HighLevelError) Error() string {
	return err.Message
}

func NewHighLevelError(code int, message string) HighLevelError {
	return HighLevelError{Code: code, Message: message}
}

// Default errors
var (
	DefaultInternalError = NewHighLevelError(http.StatusInternalServerError, "internal server error")
	BadRequest           = NewHighLevelError(http.StatusBadRequest, "missed parameter, incorrect or malformed data")
	NotFound             = NewHighLevelError(http.StatusNotFound, "not found")
	Unauthorized         = NewHighLevelError(http.StatusUnauthorized, "unauthorized")
	Forbidden            = NewHighLevelError(http.StatusForbidden, "forbidden")
	Conflict             = NewHighLevelError(http.StatusConflict, "conflict")
	TooManyRequests      = NewHighLevelError(http.StatusTooManyRequests, "too many requests")
	InternalServerError  = NewHighLevelError(http.StatusInternalServerError, "internal server error")
	NotImplemented       = NewHighLevelError(http.StatusNotImplemented, "not implemented")
	ServiceUnavailable   = NewHighLevelError(http.StatusServiceUnavailable, "service unavailable")
	GatewayTimeout       = NewHighLevelError(http.StatusGatewayTimeout, "gateway timeout")
	BadGateway           = NewHighLevelError(http.StatusBadGateway, "bad gateway")
	ProxyError           = NewHighLevelError(http.StatusBadGateway, "proxy error")
	UnknownError         = NewHighLevelError(http.StatusInternalServerError, "unknown error")
)

func ExtractError(err error) HighLevelError {
	highLevel := HighLevelError{}
	if !errors.As(err, &highLevel) {
		highLevel = DefaultInternalError
	}
	return highLevel
}

func GinErrorHandlerMiddleware(context *gin.Context) {
	context.Next()
	var message gin.H
	err := context.Errors.Last()
	if err == nil {
		context.Abort()
		return
	}
	highLevel := ExtractError(err)
	message = gin.H{"status": highLevel.Code, "error": highLevel.Message}
	context.AbortWithStatusJSON(highLevel.Code, message)
}

func AbortAndError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
