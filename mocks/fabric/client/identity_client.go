// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockfabric

import (
	msp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	mock "github.com/stretchr/testify/mock"
)

// IdentityClient is an autogenerated mock type for the IdentityClient type
type IdentityClient struct {
	mock.Mock
}

// GetClientOrg provides a mock function with given fields:
func (_m *IdentityClient) GetClientOrg() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSigningIdentity provides a mock function with given fields: name
func (_m *IdentityClient) GetSigningIdentity(name string) (msp.SigningIdentity, error) {
	ret := _m.Called(name)

	var r0 msp.SigningIdentity
	if rf, ok := ret.Get(0).(func(string) msp.SigningIdentity); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msp.SigningIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
