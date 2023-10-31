// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// CheckRequestId provides a mock function with given fields: ctx
func (_m *Request) CheckRequestId(ctx *gin.Context) *gin.Context {
	ret := _m.Called(ctx)

	var r0 *gin.Context
	if rf, ok := ret.Get(0).(func(*gin.Context) *gin.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gin.Context)
		}
	}

	return r0
}

// HasAuthenticatedUserId provides a mock function with given fields: ctx
func (_m *Request) HasAuthenticatedUserId(ctx *gin.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasDeviceToken provides a mock function with given fields: ctx
func (_m *Request) HasDeviceToken(ctx *gin.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasFile provides a mock function with given fields: ctx
func (_m *Request) HasFile(ctx *gin.Context) (multipart.File, *multipart.FileHeader, error) {
	ret := _m.Called(ctx)

	var r0 multipart.File
	if rf, ok := ret.Get(0).(func(*gin.Context) multipart.File); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multipart.File)
		}
	}

	var r1 *multipart.FileHeader
	if rf, ok := ret.Get(1).(func(*gin.Context) *multipart.FileHeader); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*multipart.FileHeader)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gin.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
