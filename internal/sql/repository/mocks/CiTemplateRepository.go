// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	pipelineConfig "github.com/devtron-labs/devtron/internal/sql/repository/pipelineConfig"
	mock "github.com/stretchr/testify/mock"
)

// CiTemplateRepository is an autogenerated mock type for the CiTemplateRepository type
type CiTemplateRepository struct {
	mock.Mock
}

// FindByAppId provides a mock function with given fields: appId
func (_m *CiTemplateRepository) FindByAppId(appId int) (*pipelineConfig.CiTemplate, error) {
	ret := _m.Called(appId)

	var r0 *pipelineConfig.CiTemplate
	if rf, ok := ret.Get(0).(func(int) *pipelineConfig.CiTemplate); ok {
		r0 = rf(appId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipelineConfig.CiTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(appId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByDockerRegistryId provides a mock function with given fields: dockerRegistryId
func (_m *CiTemplateRepository) FindByDockerRegistryId(dockerRegistryId string) ([]*pipelineConfig.CiTemplate, error) {
	ret := _m.Called(dockerRegistryId)

	var r0 []*pipelineConfig.CiTemplate
	if rf, ok := ret.Get(0).(func(string) []*pipelineConfig.CiTemplate); ok {
		r0 = rf(dockerRegistryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.CiTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dockerRegistryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNumberOfAppsWithDockerConfigured provides a mock function with given fields: appIds
func (_m *CiTemplateRepository) FindNumberOfAppsWithDockerConfigured(appIds []int) (int, error) {
	ret := _m.Called(appIds)

	var r0 int
	if rf, ok := ret.Get(0).(func([]int) int); ok {
		r0 = rf(appIds)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(appIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: material
func (_m *CiTemplateRepository) Save(material *pipelineConfig.CiTemplate) error {
	ret := _m.Called(material)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pipelineConfig.CiTemplate) error); ok {
		r0 = rf(material)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: material
func (_m *CiTemplateRepository) Update(material *pipelineConfig.CiTemplate) error {
	ret := _m.Called(material)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pipelineConfig.CiTemplate) error); ok {
		r0 = rf(material)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCiTemplateRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCiTemplateRepository creates a new instance of CiTemplateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCiTemplateRepository(t mockConstructorTestingTNewCiTemplateRepository) *CiTemplateRepository {
	mock := &CiTemplateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
