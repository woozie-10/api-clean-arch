// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/woozie-10/api-clean-arch/domain"
)

// StudentService is an autogenerated mock type for the StudentService type
type StudentService struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, student
func (_m *StudentService) Add(ctx context.Context, student *domain.Student) error {
	ret := _m.Called(ctx, student)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Student) error); ok {
		r0 = rf(ctx, student)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, username
func (_m *StudentService) Delete(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx
func (_m *StudentService) Get(ctx context.Context) ([]*domain.Student, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []*domain.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Student, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Student); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCourse provides a mock function with given fields: ctx, course
func (_m *StudentService) GetByCourse(ctx context.Context, course string) ([]*domain.Student, error) {
	ret := _m.Called(ctx, course)

	if len(ret) == 0 {
		panic("no return value specified for GetByCourse")
	}

	var r0 []*domain.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*domain.Student, error)); ok {
		return rf(ctx, course)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.Student); ok {
		r0 = rf(ctx, course)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByGroup provides a mock function with given fields: ctx, group
func (_m *StudentService) GetByGroup(ctx context.Context, group string) ([]*domain.Student, error) {
	ret := _m.Called(ctx, group)

	if len(ret) == 0 {
		panic("no return value specified for GetByGroup")
	}

	var r0 []*domain.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*domain.Student, error)); ok {
		return rf(ctx, group)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.Student); ok {
		r0 = rf(ctx, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *StudentService) GetByUsername(ctx context.Context, username string) (*domain.Student, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 *domain.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Student, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Student); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, username, newStudent
func (_m *StudentService) Update(ctx context.Context, username string, newStudent *domain.Student) error {
	ret := _m.Called(ctx, username, newStudent)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.Student) error); ok {
		r0 = rf(ctx, username, newStudent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStudentService creates a new instance of StudentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStudentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *StudentService {
	mock := &StudentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
