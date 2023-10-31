package logger

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"time"
)

type LoggerUtils struct {
}

func InitLoggerUtils() Logger {
	return &LoggerUtils{
	}
}

func (l *LoggerUtils) BuildLogger(ctx *gin.Context, logType string, statusCode int, message string) (out LoggerPayload) {
	out.Name = logType
	out.Payload.RequestId = uuid.NewV4().String()
	out.Payload.Time = time.Now().Local().Format(constant.TimeTemplateFormatDDmmYYYhhMMss)
	out.Payload.Header = ctx.Request.Header
	out.Payload.Url = ctx.Request.Host
	out.Payload.Endpoint = ctx.Request.URL.Path
	out.Payload.HttpMethod = ctx.Request.Method
	out.Payload.Proto = ctx.Request.Proto
	out.Payload.UrlQuery = ctx.Request.URL.RawQuery
	_ = ctx.ShouldBindBodyWith(&out.Payload.Body, binding.JSON)
	out.Payload.StatusCode = statusCode
	out.Payload.Message = message
	return
}
