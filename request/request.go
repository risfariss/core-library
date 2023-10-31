package request

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type Request interface {
	CheckRequestId(ctx *gin.Context) *gin.Context
	HasAuthenticatedUserId(ctx *gin.Context) (userId string, err error)
	HasDeviceToken(ctx *gin.Context) (deviceToken string, err error)
	HasFile(ctx *gin.Context) (file multipart.File,header *multipart.FileHeader, err error)
}
