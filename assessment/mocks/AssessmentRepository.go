// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/woozie-10/api-clean-arch/domain"
)

// AssessmentRepository is an autogenerated mock type for the AssessmentRepository type
type AssessmentRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, _a1
func (_m *AssessmentRepository) Add(ctx context.Context, _a1 *domain.Assessment) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Assessment) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, username
func (_m *AssessmentRepository) Get(ctx context.Context, username string) (*domain.Assessment, error) {
	ret := _m.Called(ctx, username)

	var r0 *domain.Assessment
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Assessment); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Assessment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAssessmentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAssessmentRepository creates a new instance of AssessmentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAssessmentRepository(t mockConstructorTestingTNewAssessmentRepository) *AssessmentRepository {
	mock := &AssessmentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
