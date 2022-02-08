// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockreceipt

import mock "github.com/stretchr/testify/mock"

// ReceiptStorePersistence is an autogenerated mock type for the ReceiptStorePersistence type
type ReceiptStorePersistence struct {
	mock.Mock
}

// AddReceipt provides a mock function with given fields: requestID, _a1
func (_m *ReceiptStorePersistence) AddReceipt(requestID string, _a1 *map[string]interface{}) error {
	ret := _m.Called(requestID, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *map[string]interface{}) error); ok {
		r0 = rf(requestID, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ReceiptStorePersistence) Close() {
	_m.Called()
}

// GetReceipt provides a mock function with given fields: requestID
func (_m *ReceiptStorePersistence) GetReceipt(requestID string) (*map[string]interface{}, error) {
	ret := _m.Called(requestID)

	var r0 *map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) *map[string]interface{}); ok {
		r0 = rf(requestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(requestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReceipts provides a mock function with given fields: skip, limit, ids, sinceEpochMS, from, to, start
func (_m *ReceiptStorePersistence) GetReceipts(skip int, limit int, ids []string, sinceEpochMS int64, from string, to string, start string) (*[]map[string]interface{}, error) {
	ret := _m.Called(skip, limit, ids, sinceEpochMS, from, to, start)

	var r0 *[]map[string]interface{}
	if rf, ok := ret.Get(0).(func(int, int, []string, int64, string, string, string) *[]map[string]interface{}); ok {
		r0 = rf(skip, limit, ids, sinceEpochMS, from, to, start)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, []string, int64, string, string, string) error); ok {
		r1 = rf(skip, limit, ids, sinceEpochMS, from, to, start)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *ReceiptStorePersistence) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateConf provides a mock function with given fields:
func (_m *ReceiptStorePersistence) ValidateConf() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
