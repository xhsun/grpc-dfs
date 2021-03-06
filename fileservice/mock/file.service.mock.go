// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// FileServiceMock is an autogenerated mock type for the IFileService type
type FileServiceMock struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *FileServiceMock) Close() {
	_m.Called()
}

// FileName provides a mock function with given fields:
func (_m *FileServiceMock) FileName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FileSize provides a mock function with given fields:
func (_m *FileServiceMock) FileSize() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *FileServiceMock) List() (map[string]uint64, error) {
	ret := _m.Called()

	var r0 map[string]uint64
	if rf, ok := ret.Get(0).(func() map[string]uint64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]uint64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields:
func (_m *FileServiceMock) Read() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields:
func (_m *FileServiceMock) Remove() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sync provides a mock function with given fields:
func (_m *FileServiceMock) Sync() {
	_m.Called()
}

// Write provides a mock function with given fields: data
func (_m *FileServiceMock) Write(data []byte) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewFileServiceMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewFileServiceMock creates a new instance of FileServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFileServiceMock(t NewFileServiceMockT) *FileServiceMock {
	mock := &FileServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
