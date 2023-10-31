package response

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

type Response interface {
	NewPagination(queryParam url.Values) (Pagination, error)
	SendResponseWithPagination(c *gin.Context, statusCode int, result bool, message string, data interface{}, pagination Pagination)
	SendResponseWithPaginationWithPageLimit(c *gin.Context, statusCode int, result bool, message string, data interface{}, pagination Pagination)
	SendResponse(c *gin.Context, statusCode int, result bool, message string, data interface{})
	SendResponseHtmlTemplateReset(c *gin.Context, statusCode int, result bool, message string,template string ,data interface{},baseUrl string)
	SendResponseHtmlTemplateEmailVerified(c *gin.Context, statusCode int, result bool,message string,template string ,data interface{},data2 interface{})
	SendResponseWithClientMessage(ctx *gin.Context, statusCode int, result bool, message string, data interface{}, clientMessage string)
}
