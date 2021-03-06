// Code generated by mockery v1.0.0. DO NOT EDIT.

package clients

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	stock "github.com/dictyBase/go-genproto/dictybaseapis/stock"
)

// StockServiceClient is an autogenerated mock type for the StockServiceClient type
type StockServiceClient struct {
	mock.Mock
}

// CreatePlasmid provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) CreatePlasmid(ctx context.Context, in *stock.NewPlasmid, opts ...grpc.CallOption) (*stock.Plasmid, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Plasmid
	if rf, ok := ret.Get(0).(func(context.Context, *stock.NewPlasmid, ...grpc.CallOption) *stock.Plasmid); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Plasmid)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.NewPlasmid, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStrain provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) CreateStrain(ctx context.Context, in *stock.NewStrain, opts ...grpc.CallOption) (*stock.Strain, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Strain
	if rf, ok := ret.Get(0).(func(context.Context, *stock.NewStrain, ...grpc.CallOption) *stock.Strain); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Strain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.NewStrain, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlasmid provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) GetPlasmid(ctx context.Context, in *stock.StockId, opts ...grpc.CallOption) (*stock.Plasmid, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Plasmid
	if rf, ok := ret.Get(0).(func(context.Context, *stock.StockId, ...grpc.CallOption) *stock.Plasmid); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Plasmid)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.StockId, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStrain provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) GetStrain(ctx context.Context, in *stock.StockId, opts ...grpc.CallOption) (*stock.Strain, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Strain
	if rf, ok := ret.Get(0).(func(context.Context, *stock.StockId, ...grpc.CallOption) *stock.Strain); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Strain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.StockId, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlasmids provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) ListPlasmids(ctx context.Context, in *stock.StockParameters, opts ...grpc.CallOption) (*stock.PlasmidCollection, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.PlasmidCollection
	if rf, ok := ret.Get(0).(func(context.Context, *stock.StockParameters, ...grpc.CallOption) *stock.PlasmidCollection); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.PlasmidCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.StockParameters, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStrains provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) ListStrains(ctx context.Context, in *stock.StockParameters, opts ...grpc.CallOption) (*stock.StrainCollection, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.StrainCollection
	if rf, ok := ret.Get(0).(func(context.Context, *stock.StockParameters, ...grpc.CallOption) *stock.StrainCollection); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.StrainCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.StockParameters, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadPlasmid provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) LoadPlasmid(ctx context.Context, in *stock.ExistingPlasmid, opts ...grpc.CallOption) (*stock.Plasmid, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Plasmid
	if rf, ok := ret.Get(0).(func(context.Context, *stock.ExistingPlasmid, ...grpc.CallOption) *stock.Plasmid); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Plasmid)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.ExistingPlasmid, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadStrain provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) LoadStrain(ctx context.Context, in *stock.ExistingStrain, opts ...grpc.CallOption) (*stock.Strain, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Strain
	if rf, ok := ret.Get(0).(func(context.Context, *stock.ExistingStrain, ...grpc.CallOption) *stock.Strain); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Strain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.ExistingStrain, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveStock provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) RemoveStock(ctx context.Context, in *stock.StockId, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *stock.StockId, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.StockId, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlasmid provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) UpdatePlasmid(ctx context.Context, in *stock.PlasmidUpdate, opts ...grpc.CallOption) (*stock.Plasmid, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Plasmid
	if rf, ok := ret.Get(0).(func(context.Context, *stock.PlasmidUpdate, ...grpc.CallOption) *stock.Plasmid); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Plasmid)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.PlasmidUpdate, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStrain provides a mock function with given fields: ctx, in, opts
func (_m *StockServiceClient) UpdateStrain(ctx context.Context, in *stock.StrainUpdate, opts ...grpc.CallOption) (*stock.Strain, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stock.Strain
	if rf, ok := ret.Get(0).(func(context.Context, *stock.StrainUpdate, ...grpc.CallOption) *stock.Strain); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stock.Strain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stock.StrainUpdate, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
