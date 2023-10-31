// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Loan is an autogenerated mock type for the Loan type
type Loan struct {
	mock.Mock
}

// CheckIsFunded provides a mock function with given fields: loanStatus
func (_m *Loan) CheckIsFunded(loanStatus int) (bool, int, string) {
	ret := _m.Called(loanStatus)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(loanStatus)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(loanStatus)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(int) string); ok {
		r2 = rf(loanStatus)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// CheckIsLate provides a mock function with given fields: loanStatus, dueDate
func (_m *Loan) CheckIsLate(loanStatus int, dueDate int64) (bool, int, string) {
	ret := _m.Called(loanStatus, dueDate)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int64) bool); ok {
		r0 = rf(loanStatus, dueDate)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int64) int); ok {
		r1 = rf(loanStatus, dueDate)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(int, int64) string); ok {
		r2 = rf(loanStatus, dueDate)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// CheckTenor provides a mock function with given fields: tenor, tenorPeriod
func (_m *Loan) CheckTenor(tenor int, tenorPeriod int) (int, int, string) {
	ret := _m.Called(tenor, tenorPeriod)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(tenor, tenorPeriod)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(tenor, tenorPeriod)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(int, int) string); ok {
		r2 = rf(tenor, tenorPeriod)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// GenerateLoanNo provides a mock function with given fields: seq, loanCode, year, month, day
func (_m *Loan) GenerateLoanNo(seq int, loanCode string, year int, month time.Month, day int) string {
	ret := _m.Called(seq, loanCode, year, month, day)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, string, int, time.Month, int) string); ok {
		r0 = rf(seq, loanCode, year, month, day)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
