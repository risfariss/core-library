package response

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"bitbucket.org/kawancicil/core-library/external/rabbitMQ"
	"bitbucket.org/kawancicil/core-library/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

type ResponseUtils struct {
	/*todo fill this block if need connection with another interface*/
	logger logger.Logger
	rabbit rabbitMQ.RabbitMQ
}

func InitResponseUtils(logger logger.Logger,
	rabbit rabbitMQ.RabbitMQ) Response {
	return &ResponseUtils{
		logger: logger,
		rabbit: rabbit,
	}
}

type Pagination struct {
	Page      int    `response:"page"`
	Limit     int    `response:"limit"`
	Offset    int    `response:"offset"`
	Order     string `response:"order"`
	TotalData int    `response:"totalData"`
	TotalPage int    `response:"totalPage"`
}

func (r *ResponseUtils) NewPagination(queryParam url.Values) (Pagination, error) {
	pagination := Pagination{}
	pagination.Page = 0
	pagination.Limit = 10
	pagination.Order = "id DESC"
	if queryParam.Get("page") != "" {
		page, err := strconv.Atoi(queryParam.Get("page"))
		if err != nil {
			return pagination, errors.New("invalid page parameter")
		}
		pagination.Page = page - 1
	}

	if queryParam.Get("limit") != "" {
		limit, err := strconv.Atoi(queryParam.Get("limit"))
		if err != nil {
			return pagination, errors.New("invalid page parameter")
		}
		pagination.Limit = limit
	}

	pagination.Offset = pagination.Limit * pagination.Page
	return pagination, nil
}

//Send used to send the response
func (r *ResponseUtils) SendResponseWithPagination(ctx *gin.Context, statusCode int, result bool, message string, data interface{}, pagination Pagination) {
	if result == true {
		paginationResponse := map[string]interface{}{
			"total_data": pagination.TotalData,
			"total_page": pagination.TotalPage,
		}
		ctx.JSON(statusCode, gin.H{"message": message, "data": data, "pagination": paginationResponse})
		return
	}
	ctx.AbortWithStatusJSON(statusCode, gin.H{"error_description": message})
}

//Send used to send the response
func (r *ResponseUtils) SendResponseWithPaginationWithPageLimit(ctx *gin.Context, statusCode int, result bool, message string, data interface{}, pagination Pagination) {
	if result == true {
		paginationResponse := map[string]interface{}{
			"totalData": pagination.TotalData,
			"totalPage": pagination.TotalPage,
			"page":      pagination.Page + 1,
			"limit":     pagination.Limit,
		}
		ctx.JSON(statusCode, gin.H{"message": message, "data": data, "pagination": paginationResponse})
		return
	}
}

func (r *ResponseUtils) SendResponse(ctx *gin.Context, statusCode int, result bool, message string, data interface{}) {
	if result == true {
		ctx.JSON(statusCode, gin.H{"message": message, "data": data})
		return
	}


	rabbitMQDialUrl := fmt.Sprint(ctx.MustGet(constant.RabbitMQDialUrl))
	log := r.logger.BuildLogger(ctx, constant.RabbitMQQueueHttpRequest, statusCode, message)
	r.rabbit.PublishLogger(rabbitMQDialUrl, log)
	ctx.AbortWithStatusJSON(statusCode, gin.H{"error_description": message})
}

//Send used to send the response
func (r *ResponseUtils) SendResponseHtmlTemplateEmailVerified(c *gin.Context, statusCode int, result bool, message string, template string, data interface{}, data2 interface{}) {
	if message != "" {
		c.HTML(statusCode, template, gin.H{
			"Success": false,
			"Message": constant.ErrorInvalidParameterMsg,
		})
		return

	}

	if result == true {
		c.HTML(statusCode, template, gin.H{
			"Title":          "Congratulations!",
			"Description":    "Email verification is complete",
			"AndroidAppLink": data,
			"IOSAppLink":     data2,
		})

		return
	}
	c.HTML(statusCode, template, gin.H{
		"Title":       "Email verification failed",
		"Description": data,
	})
	return
}

//Send used to send the response
func (r *ResponseUtils) SendResponseHtmlTemplateReset(c *gin.Context, statusCode int, result bool, message string, template string, data interface{}, baseUrl string) {
	if result == true {
		c.HTML(statusCode, template, gin.H{
			"Success":     true,
			"Message":     message,
			"BaseURL":     baseUrl,
			"NewPassword": data,
		})

		return
	}
	c.HTML(statusCode, template, gin.H{
		"Success": false,
		"Message": message,
	})
	return
}

func (r *ResponseUtils) SendResponseWithClientMessage(ctx *gin.Context, statusCode int, result bool, message string, data interface{}, clientMessage string) {
	if result == true {
		ctx.JSON(statusCode, gin.H{"message": message, "data": data, "clientMessage": clientMessage})
		return
	}

	rabbitMQDialUrl := fmt.Sprint(ctx.MustGet(constant.RabbitMQDialUrl))
	log := r.logger.BuildLogger(ctx, "httpRequest", statusCode, message)
	r.rabbit.PublishLogger(rabbitMQDialUrl, log)
	ctx.AbortWithStatusJSON(statusCode, gin.H{"error_description": message, "error_description_client": clientMessage})
}
