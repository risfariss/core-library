package logger

import "github.com/gin-gonic/gin"

type Logger interface {
	BuildLogger(ctx *gin.Context, logType string, statusCode int, message string) (out LoggerPayload)
}
