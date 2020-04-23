// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import annotation "github.com/dictyBase/go-genproto/dictybaseapis/annotation"
import auth "github.com/dictyBase/go-genproto/dictybaseapis/auth"
import content "github.com/dictyBase/go-genproto/dictybaseapis/content"
import grpc "google.golang.org/grpc"
import identity "github.com/dictyBase/go-genproto/dictybaseapis/identity"
import mock "github.com/stretchr/testify/mock"
import order "github.com/dictyBase/go-genproto/dictybaseapis/order"

import stock "github.com/dictyBase/go-genproto/dictybaseapis/stock"
import user "github.com/dictyBase/go-genproto/dictybaseapis/user"

// Registry is an autogenerated mock type for the Registry type
type Registry struct {
	mock.Mock
}

// AddAPIConnection provides a mock function with given fields: key, conn
func (_m *Registry) AddAPIConnection(key string, conn *grpc.ClientConn) {
	_m.Called(key, conn)
}

// AddAPIEndpoint provides a mock function with given fields: key, endpoint
func (_m *Registry) AddAPIEndpoint(key string, endpoint string) {
	_m.Called(key, endpoint)
}

// GetAPIConnection provides a mock function with given fields: key
func (_m *Registry) GetAPIConnection(key string) *grpc.ClientConn {
	ret := _m.Called(key)

	var r0 *grpc.ClientConn
	if rf, ok := ret.Get(0).(func(string) *grpc.ClientConn); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.ClientConn)
		}
	}

	return r0
}

// GetAPIEndpoint provides a mock function with given fields: key
func (_m *Registry) GetAPIEndpoint(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetAnnotationClient provides a mock function with given fields: key
func (_m *Registry) GetAnnotationClient(key string) annotation.TaggedAnnotationServiceClient {
	ret := _m.Called(key)

	var r0 annotation.TaggedAnnotationServiceClient
	if rf, ok := ret.Get(0).(func(string) annotation.TaggedAnnotationServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(annotation.TaggedAnnotationServiceClient)
		}
	}

	return r0
}

// GetAuthClient provides a mock function with given fields: key
func (_m *Registry) GetAuthClient(key string) auth.AuthServiceClient {
	ret := _m.Called(key)

	var r0 auth.AuthServiceClient
	if rf, ok := ret.Get(0).(func(string) auth.AuthServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.AuthServiceClient)
		}
	}

	return r0
}

// GetContentClient provides a mock function with given fields: key
func (_m *Registry) GetContentClient(key string) content.ContentServiceClient {
	ret := _m.Called(key)

	var r0 content.ContentServiceClient
	if rf, ok := ret.Get(0).(func(string) content.ContentServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(content.ContentServiceClient)
		}
	}

	return r0
}

// GetIdentityClient provides a mock function with given fields: key
func (_m *Registry) GetIdentityClient(key string) identity.IdentityServiceClient {
	ret := _m.Called(key)

	var r0 identity.IdentityServiceClient
	if rf, ok := ret.Get(0).(func(string) identity.IdentityServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(identity.IdentityServiceClient)
		}
	}

	return r0
}

// GetOrderClient provides a mock function with given fields: key
func (_m *Registry) GetOrderClient(key string) order.OrderServiceClient {
	ret := _m.Called(key)

	var r0 order.OrderServiceClient
	if rf, ok := ret.Get(0).(func(string) order.OrderServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(order.OrderServiceClient)
		}
	}

	return r0
}

// GetPermissionClient provides a mock function with given fields: key
func (_m *Registry) GetPermissionClient(key string) user.PermissionServiceClient {
	ret := _m.Called(key)

	var r0 user.PermissionServiceClient
	if rf, ok := ret.Get(0).(func(string) user.PermissionServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.PermissionServiceClient)
		}
	}

	return r0
}

// GetRoleClient provides a mock function with given fields: key
func (_m *Registry) GetRoleClient(key string) user.RoleServiceClient {
	ret := _m.Called(key)

	var r0 user.RoleServiceClient
	if rf, ok := ret.Get(0).(func(string) user.RoleServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.RoleServiceClient)
		}
	}

	return r0
}

// GetStockClient provides a mock function with given fields: key
func (_m *Registry) GetStockClient(key string) stock.StockServiceClient {
	ret := _m.Called(key)

	var r0 stock.StockServiceClient
	if rf, ok := ret.Get(0).(func(string) stock.StockServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stock.StockServiceClient)
		}
	}

	return r0
}

// GetUserClient provides a mock function with given fields: key
func (_m *Registry) GetUserClient(key string) user.UserServiceClient {
	ret := _m.Called(key)

	var r0 user.UserServiceClient
	if rf, ok := ret.Get(0).(func(string) user.UserServiceClient); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.UserServiceClient)
		}
	}

	return r0
}
