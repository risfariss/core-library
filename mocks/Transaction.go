// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// GenerateClientRef provides a mock function with given fields: seq
func (_m *Transaction) GenerateClientRef(seq int) string {
	ret := _m.Called(seq)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(seq)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}