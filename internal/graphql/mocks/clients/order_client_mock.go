// Code generated by mockery v1.0.0. DO NOT EDIT.

package clients

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	order "github.com/dictyBase/go-genproto/dictybaseapis/order"
)

// OrderServiceClient is an autogenerated mock type for the OrderServiceClient type
type OrderServiceClient struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) CreateOrder(ctx context.Context, in *order.NewOrder, opts ...grpc.CallOption) (*order.Order, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *order.Order
	if rf, ok := ret.Get(0).(func(context.Context, *order.NewOrder, ...grpc.CallOption) *order.Order); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *order.NewOrder, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) GetOrder(ctx context.Context, in *order.OrderId, opts ...grpc.CallOption) (*order.Order, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *order.Order
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderId, ...grpc.CallOption) *order.Order); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *order.OrderId, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrders provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) ListOrders(ctx context.Context, in *order.ListParameters, opts ...grpc.CallOption) (*order.OrderCollection, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *order.OrderCollection
	if rf, ok := ret.Get(0).(func(context.Context, *order.ListParameters, ...grpc.CallOption) *order.OrderCollection); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.OrderCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *order.ListParameters, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) LoadOrder(ctx context.Context, in *order.ExistingOrder, opts ...grpc.CallOption) (*order.Order, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *order.Order
	if rf, ok := ret.Get(0).(func(context.Context, *order.ExistingOrder, ...grpc.CallOption) *order.Order); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *order.ExistingOrder, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareForOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) PrepareForOrder(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *empty.Empty, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *empty.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) UpdateOrder(ctx context.Context, in *order.OrderUpdate, opts ...grpc.CallOption) (*order.Order, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *order.Order
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderUpdate, ...grpc.CallOption) *order.Order); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *order.OrderUpdate, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
