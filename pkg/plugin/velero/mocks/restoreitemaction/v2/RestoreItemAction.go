/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0. DO NOT EDIT.

package v2

import (
	mock "github.com/stretchr/testify/mock"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"

	velero "github.com/vmware-tanzu/velero/pkg/plugin/velero"
)

// RestoreItemAction is an autogenerated mock type for the RestoreItemAction type
type RestoreItemAction struct {
	mock.Mock
}

// AppliesTo provides a mock function with given fields:
func (_m *RestoreItemAction) AppliesTo() (velero.ResourceSelector, error) {
	ret := _m.Called()

	var r0 velero.ResourceSelector
	if rf, ok := ret.Get(0).(func() velero.ResourceSelector); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(velero.ResourceSelector)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AreAdditionalItemsReady provides a mock function with given fields: AdditionalItems, restore
func (_m *RestoreItemAction) AreAdditionalItemsReady(AdditionalItems []velero.ResourceIdentifier, restore *v1.Restore) (bool, error) {
	ret := _m.Called(AdditionalItems, restore)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]velero.ResourceIdentifier, *v1.Restore) bool); ok {
		r0 = rf(AdditionalItems, restore)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]velero.ResourceIdentifier, *v1.Restore) error); ok {
		r1 = rf(AdditionalItems, restore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cancel provides a mock function with given fields: operationID, restore
func (_m *RestoreItemAction) Cancel(operationID string, restore *v1.Restore) error {
	ret := _m.Called(operationID, restore)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.Restore) error); ok {
		r0 = rf(operationID, restore)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Execute provides a mock function with given fields: input
func (_m *RestoreItemAction) Execute(input *velero.RestoreItemActionExecuteInput) (*velero.RestoreItemActionExecuteOutput, error) {
	ret := _m.Called(input)

	var r0 *velero.RestoreItemActionExecuteOutput
	if rf, ok := ret.Get(0).(func(*velero.RestoreItemActionExecuteInput) *velero.RestoreItemActionExecuteOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*velero.RestoreItemActionExecuteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*velero.RestoreItemActionExecuteInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Progress provides a mock function with given fields: operationID, restore
func (_m *RestoreItemAction) Progress(operationID string, restore *v1.Restore) (velero.OperationProgress, error) {
	ret := _m.Called(operationID, restore)

	var r0 velero.OperationProgress
	if rf, ok := ret.Get(0).(func(string, *v1.Restore) velero.OperationProgress); ok {
		r0 = rf(operationID, restore)
	} else {
		r0 = ret.Get(0).(velero.OperationProgress)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *v1.Restore) error); ok {
		r1 = rf(operationID, restore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}