// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	protos "github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
)

// UserServiceClient is an autogenerated mock type for the UserServiceClient type
type UserServiceClient struct {
	mock.Mock
}

// GetUserById provides a mock function with given fields: ctx, in, opts
func (_m *UserServiceClient) GetUserById(ctx context.Context, in *protos.GetUser, opts ...grpc.CallOption) (*protos.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser, ...grpc.CallOption) *protos.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.GetUser, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: ctx, in, opts
func (_m *UserServiceClient) GetUserByUsername(ctx context.Context, in *protos.GetUser, opts ...grpc.CallOption) (*protos.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser, ...grpc.CallOption) *protos.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.GetUser, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: ctx, in, opts
func (_m *UserServiceClient) SaveUser(ctx context.Context, in *protos.User, opts ...grpc.CallOption) (*protos.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User, ...grpc.CallOption) *protos.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.User, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, in, opts
func (_m *UserServiceClient) UpdateUser(ctx context.Context, in *protos.User, opts ...grpc.CallOption) (*protos.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User, ...grpc.CallOption) *protos.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.User, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
