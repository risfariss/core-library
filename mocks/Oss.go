// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	mock "github.com/stretchr/testify/mock"

	request "bitbucket.org/kawancicil/core-library/external/oss/request"
)

// Oss is an autogenerated mock type for the Oss type
type Oss struct {
	mock.Mock
}

// DownloadFile provides a mock function with given fields: bucketConfig, file
func (_m *Oss) DownloadFile(bucketConfig *request.OssBucketConfig, file string) (*bytes.Buffer, error) {
	ret := _m.Called(bucketConfig, file)

	var r0 *bytes.Buffer
	if rf, ok := ret.Get(0).(func(*request.OssBucketConfig, string) *bytes.Buffer); ok {
		r0 = rf(bucketConfig, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bytes.Buffer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.OssBucketConfig, string) error); ok {
		r1 = rf(bucketConfig, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFile provides a mock function with given fields: bucketConfig, fileServerPath
func (_m *Oss) UploadFile(bucketConfig *request.OssBucketConfig, fileServerPath string) (string, error) {
	ret := _m.Called(bucketConfig, fileServerPath)

	var r0 string
	if rf, ok := ret.Get(0).(func(*request.OssBucketConfig, string) string); ok {
		r0 = rf(bucketConfig, fileServerPath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.OssBucketConfig, string) error); ok {
		r1 = rf(bucketConfig, fileServerPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}