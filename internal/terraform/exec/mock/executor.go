// Code generated by mockery v2.3.0. DO NOT EDIT.

package mock

import (
	context "context"

	log "log"

	mock "github.com/stretchr/testify/mock"

	tfjson "github.com/hashicorp/terraform-json"

	time "time"

	version "github.com/hashicorp/go-version"
)

// Executor is an autogenerated mock type for the TerraformExecutor type
type Executor struct {
	mock.Mock
}

// Format provides a mock function with given fields: ctx, input
func (_m *Executor) Format(ctx context.Context, input []byte) ([]byte, error) {
	ret := _m.Called(ctx, input)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte) []byte); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecPath provides a mock function with given fields:
func (_m *Executor) GetExecPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Init provides a mock function with given fields: ctx
func (_m *Executor) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProviderSchemas provides a mock function with given fields: ctx
func (_m *Executor) ProviderSchemas(ctx context.Context) (*tfjson.ProviderSchemas, error) {
	ret := _m.Called(ctx)

	var r0 *tfjson.ProviderSchemas
	if rf, ok := ret.Get(0).(func(context.Context) *tfjson.ProviderSchemas); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfjson.ProviderSchemas)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetExecLogPath provides a mock function with given fields: path
func (_m *Executor) SetExecLogPath(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogger provides a mock function with given fields: logger
func (_m *Executor) SetLogger(logger *log.Logger) {
	_m.Called(logger)
}

// SetTimeout provides a mock function with given fields: duration
func (_m *Executor) SetTimeout(duration time.Duration) {
	_m.Called(duration)
}

// Version provides a mock function with given fields: ctx
func (_m *Executor) Version(ctx context.Context) (*version.Version, error) {
	ret := _m.Called(ctx)

	var r0 *version.Version
	if rf, ok := ret.Get(0).(func(context.Context) *version.Version); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
