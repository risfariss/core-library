package request

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

type RequestUtils struct {
	/*todo fill this block if need connection with another interface*/
}

func InitRequestUtils() Request {
	return &RequestUtils{
	}
}

func (r *RequestUtils) CheckRequestId(ctx *gin.Context) *gin.Context {
	requestId := ctx.GetHeader(constant.HeaderXRequestId)
	if requestId == "" {
		ctx.Request.Header.Set(constant.HeaderXRequestId, uuid.NewV4().String())
	}
	return ctx
}

func (r *RequestUtils) HasAuthenticatedUserId(ctx *gin.Context) (userId string, err error) {
	userId = ctx.GetHeader(constant.HeaderXAuthenticatedUserId)
	if userId != "" {
		return userId, nil
	}
	return "", errors.New(constant.ErrorNotSendUserId)
}

func (r *RequestUtils) HasDeviceToken(ctx *gin.Context) (deviceToken string, err error) {
	deviceToken = ctx.GetHeader("device_token")
	if deviceToken != "" {
		return deviceToken, nil
	}
	return "", errors.New(constant.ErrorNotSendUserId)
}

func (r *RequestUtils) HasFile(ctx *gin.Context) (file multipart.File,header *multipart.FileHeader, err error) {
	file, header, err = ctx.Request.FormFile("file")
	if err != nil {
		return
	}
	return
}
