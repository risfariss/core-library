// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	response "bitbucket.org/kawancicil/core-library/response"

	url "net/url"
)

// Response is an autogenerated mock type for the Response type
type Response struct {
	mock.Mock
}

// NewPagination provides a mock function with given fields: queryParam
func (_m *Response) NewPagination(queryParam url.Values) (response.Pagination, error) {
	ret := _m.Called(queryParam)

	var r0 response.Pagination
	if rf, ok := ret.Get(0).(func(url.Values) response.Pagination); ok {
		r0 = rf(queryParam)
	} else {
		r0 = ret.Get(0).(response.Pagination)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(url.Values) error); ok {
		r1 = rf(queryParam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendResponse provides a mock function with given fields: c, statusCode, result, message, data
func (_m *Response) SendResponse(c *gin.Context, statusCode int, result bool, message string, data interface{}) {
	_m.Called(c, statusCode, result, message, data)
}

// SendResponseHtmlTemplateEmailVerified provides a mock function with given fields: c, statusCode, result, message, template, data, data2
func (_m *Response) SendResponseHtmlTemplateEmailVerified(c *gin.Context, statusCode int, result bool, message string, template string, data interface{}, data2 interface{}) {
	_m.Called(c, statusCode, result, message, template, data, data2)
}

// SendResponseHtmlTemplateReset provides a mock function with given fields: c, statusCode, result, message, template, data, baseUrl
func (_m *Response) SendResponseHtmlTemplateReset(c *gin.Context, statusCode int, result bool, message string, template string, data interface{}, baseUrl string) {
	_m.Called(c, statusCode, result, message, template, data, baseUrl)
}

// SendResponseWithPagination provides a mock function with given fields: c, statusCode, result, message, data, pagination
func (_m *Response) SendResponseWithPagination(c *gin.Context, statusCode int, result bool, message string, data interface{}, pagination response.Pagination) {
	_m.Called(c, statusCode, result, message, data, pagination)
}

// SendResponseWithPaginationWithPageLimit provides a mock function with given fields: c, statusCode, result, message, data, pagination
func (_m *Response) SendResponseWithPaginationWithPageLimit(c *gin.Context, statusCode int, result bool, message string, data interface{}, pagination response.Pagination) {
	_m.Called(c, statusCode, result, message, data, pagination)
}