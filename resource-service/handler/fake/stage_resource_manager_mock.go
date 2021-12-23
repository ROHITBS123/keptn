// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler_mock

import (
	"github.com/keptn/keptn/resource-service/models"
	"sync"
)

// IStageResourceManagerMock is a mock implementation of handler.IStageResourceManager.
//
// 	func TestSomethingThatUsesIStageResourceManager(t *testing.T) {
//
// 		// make and configure a mocked handler.IStageResourceManager
// 		mockedIStageResourceManager := &IStageResourceManagerMock{
// 			CreateStageResourcesFunc: func(projectName string, stageName string, params models.CreateResourcesParams)  {
// 				panic("mock out the CreateStageResources method")
// 			},
// 			DeleteStageResourceFunc: func(projectName string, stageName string, resourceURI string) error {
// 				panic("mock out the DeleteStageResource method")
// 			},
// 			GetStageResourceFunc: func(projectName string, stageName string, resourceURI string) (models.GetResourceResponse, error) {
// 				panic("mock out the GetStageResource method")
// 			},
// 			GetStageResourcesFunc: func(projectName string, stageName string, gitCommitID string) (models.GetResourcesResponse, error) {
// 				panic("mock out the GetStageResources method")
// 			},
// 			UpdateStageResourceFunc: func(projectName string, stageName string, params models.UpdateResourceParams) error {
// 				panic("mock out the UpdateStageResource method")
// 			},
// 			UpdateStageResourcesFunc: func(projectName string, stageName string, params models.UpdateResourcesParams) error {
// 				panic("mock out the UpdateStageResources method")
// 			},
// 		}
//
// 		// use mockedIStageResourceManager in code that requires handler.IStageResourceManager
// 		// and then make assertions.
//
// 	}
type IStageResourceManagerMock struct {
	// CreateStageResourcesFunc mocks the CreateStageResources method.
	CreateStageResourcesFunc func(projectName string, stageName string, params models.CreateResourcesParams)

	// DeleteStageResourceFunc mocks the DeleteStageResource method.
	DeleteStageResourceFunc func(projectName string, stageName string, resourceURI string) error

	// GetStageResourceFunc mocks the GetStageResource method.
	GetStageResourceFunc func(projectName string, stageName string, resourceURI string) (models.GetResourceResponse, error)

	// GetStageResourcesFunc mocks the GetStageResources method.
	GetStageResourcesFunc func(projectName string, stageName string, gitCommitID string) (models.GetResourcesResponse, error)

	// UpdateStageResourceFunc mocks the UpdateStageResource method.
	UpdateStageResourceFunc func(projectName string, stageName string, params models.UpdateResourceParams) error

	// UpdateStageResourcesFunc mocks the UpdateStageResources method.
	UpdateStageResourcesFunc func(projectName string, stageName string, params models.UpdateResourcesParams) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateStageResources holds details about calls to the CreateStageResources method.
		CreateStageResources []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// Params is the params argument value.
			Params models.CreateResourcesParams
		}
		// DeleteStageResource holds details about calls to the DeleteStageResource method.
		DeleteStageResource []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// ResourceURI is the resourceURI argument value.
			ResourceURI string
		}
		// GetStageResource holds details about calls to the GetStageResource method.
		GetStageResource []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// ResourceURI is the resourceURI argument value.
			ResourceURI string
		}
		// GetStageResources holds details about calls to the GetStageResources method.
		GetStageResources []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// GitCommitID is the gitCommitID argument value.
			GitCommitID string
		}
		// UpdateStageResource holds details about calls to the UpdateStageResource method.
		UpdateStageResource []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// Params is the params argument value.
			Params models.UpdateResourceParams
		}
		// UpdateStageResources holds details about calls to the UpdateStageResources method.
		UpdateStageResources []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// Params is the params argument value.
			Params models.UpdateResourcesParams
		}
	}
	lockCreateStageResources sync.RWMutex
	lockDeleteStageResource  sync.RWMutex
	lockGetStageResource     sync.RWMutex
	lockGetStageResources    sync.RWMutex
	lockUpdateStageResource  sync.RWMutex
	lockUpdateStageResources sync.RWMutex
}

// CreateStageResources calls CreateStageResourcesFunc.
func (mock *IStageResourceManagerMock) CreateStageResources(projectName string, stageName string, params models.CreateResourcesParams) {
	if mock.CreateStageResourcesFunc == nil {
		panic("IStageResourceManagerMock.CreateStageResourcesFunc: method is nil but IStageResourceManager.CreateStageResources was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		Params      models.CreateResourcesParams
	}{
		ProjectName: projectName,
		StageName:   stageName,
		Params:      params,
	}
	mock.lockCreateStageResources.Lock()
	mock.calls.CreateStageResources = append(mock.calls.CreateStageResources, callInfo)
	mock.lockCreateStageResources.Unlock()
	mock.CreateStageResourcesFunc(projectName, stageName, params)
}

// CreateStageResourcesCalls gets all the calls that were made to CreateStageResources.
// Check the length with:
//     len(mockedIStageResourceManager.CreateStageResourcesCalls())
func (mock *IStageResourceManagerMock) CreateStageResourcesCalls() []struct {
	ProjectName string
	StageName   string
	Params      models.CreateResourcesParams
} {
	var calls []struct {
		ProjectName string
		StageName   string
		Params      models.CreateResourcesParams
	}
	mock.lockCreateStageResources.RLock()
	calls = mock.calls.CreateStageResources
	mock.lockCreateStageResources.RUnlock()
	return calls
}

// DeleteStageResource calls DeleteStageResourceFunc.
func (mock *IStageResourceManagerMock) DeleteStageResource(projectName string, stageName string, resourceURI string) error {
	if mock.DeleteStageResourceFunc == nil {
		panic("IStageResourceManagerMock.DeleteStageResourceFunc: method is nil but IStageResourceManager.DeleteStageResource was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		ResourceURI string
	}{
		ProjectName: projectName,
		StageName:   stageName,
		ResourceURI: resourceURI,
	}
	mock.lockDeleteStageResource.Lock()
	mock.calls.DeleteStageResource = append(mock.calls.DeleteStageResource, callInfo)
	mock.lockDeleteStageResource.Unlock()
	return mock.DeleteStageResourceFunc(projectName, stageName, resourceURI)
}

// DeleteStageResourceCalls gets all the calls that were made to DeleteStageResource.
// Check the length with:
//     len(mockedIStageResourceManager.DeleteStageResourceCalls())
func (mock *IStageResourceManagerMock) DeleteStageResourceCalls() []struct {
	ProjectName string
	StageName   string
	ResourceURI string
} {
	var calls []struct {
		ProjectName string
		StageName   string
		ResourceURI string
	}
	mock.lockDeleteStageResource.RLock()
	calls = mock.calls.DeleteStageResource
	mock.lockDeleteStageResource.RUnlock()
	return calls
}

// GetStageResource calls GetStageResourceFunc.
func (mock *IStageResourceManagerMock) GetStageResource(projectName string, stageName string, resourceURI string) (models.GetResourceResponse, error) {
	if mock.GetStageResourceFunc == nil {
		panic("IStageResourceManagerMock.GetStageResourceFunc: method is nil but IStageResourceManager.GetStageResource was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		ResourceURI string
	}{
		ProjectName: projectName,
		StageName:   stageName,
		ResourceURI: resourceURI,
	}
	mock.lockGetStageResource.Lock()
	mock.calls.GetStageResource = append(mock.calls.GetStageResource, callInfo)
	mock.lockGetStageResource.Unlock()
	return mock.GetStageResourceFunc(projectName, stageName, resourceURI)
}

// GetStageResourceCalls gets all the calls that were made to GetStageResource.
// Check the length with:
//     len(mockedIStageResourceManager.GetStageResourceCalls())
func (mock *IStageResourceManagerMock) GetStageResourceCalls() []struct {
	ProjectName string
	StageName   string
	ResourceURI string
} {
	var calls []struct {
		ProjectName string
		StageName   string
		ResourceURI string
	}
	mock.lockGetStageResource.RLock()
	calls = mock.calls.GetStageResource
	mock.lockGetStageResource.RUnlock()
	return calls
}

// GetStageResources calls GetStageResourcesFunc.
func (mock *IStageResourceManagerMock) GetStageResources(projectName string, stageName string, gitCommitID string) (models.GetResourcesResponse, error) {
	if mock.GetStageResourcesFunc == nil {
		panic("IStageResourceManagerMock.GetStageResourcesFunc: method is nil but IStageResourceManager.GetStageResources was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		GitCommitID string
	}{
		ProjectName: projectName,
		StageName:   stageName,
		GitCommitID: gitCommitID,
	}
	mock.lockGetStageResources.Lock()
	mock.calls.GetStageResources = append(mock.calls.GetStageResources, callInfo)
	mock.lockGetStageResources.Unlock()
	return mock.GetStageResourcesFunc(projectName, stageName, gitCommitID)
}

// GetStageResourcesCalls gets all the calls that were made to GetStageResources.
// Check the length with:
//     len(mockedIStageResourceManager.GetStageResourcesCalls())
func (mock *IStageResourceManagerMock) GetStageResourcesCalls() []struct {
	ProjectName string
	StageName   string
	GitCommitID string
} {
	var calls []struct {
		ProjectName string
		StageName   string
		GitCommitID string
	}
	mock.lockGetStageResources.RLock()
	calls = mock.calls.GetStageResources
	mock.lockGetStageResources.RUnlock()
	return calls
}

// UpdateStageResource calls UpdateStageResourceFunc.
func (mock *IStageResourceManagerMock) UpdateStageResource(projectName string, stageName string, params models.UpdateResourceParams) error {
	if mock.UpdateStageResourceFunc == nil {
		panic("IStageResourceManagerMock.UpdateStageResourceFunc: method is nil but IStageResourceManager.UpdateStageResource was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		Params      models.UpdateResourceParams
	}{
		ProjectName: projectName,
		StageName:   stageName,
		Params:      params,
	}
	mock.lockUpdateStageResource.Lock()
	mock.calls.UpdateStageResource = append(mock.calls.UpdateStageResource, callInfo)
	mock.lockUpdateStageResource.Unlock()
	return mock.UpdateStageResourceFunc(projectName, stageName, params)
}

// UpdateStageResourceCalls gets all the calls that were made to UpdateStageResource.
// Check the length with:
//     len(mockedIStageResourceManager.UpdateStageResourceCalls())
func (mock *IStageResourceManagerMock) UpdateStageResourceCalls() []struct {
	ProjectName string
	StageName   string
	Params      models.UpdateResourceParams
} {
	var calls []struct {
		ProjectName string
		StageName   string
		Params      models.UpdateResourceParams
	}
	mock.lockUpdateStageResource.RLock()
	calls = mock.calls.UpdateStageResource
	mock.lockUpdateStageResource.RUnlock()
	return calls
}

// UpdateStageResources calls UpdateStageResourcesFunc.
func (mock *IStageResourceManagerMock) UpdateStageResources(projectName string, stageName string, params models.UpdateResourcesParams) error {
	if mock.UpdateStageResourcesFunc == nil {
		panic("IStageResourceManagerMock.UpdateStageResourcesFunc: method is nil but IStageResourceManager.UpdateStageResources was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		Params      models.UpdateResourcesParams
	}{
		ProjectName: projectName,
		StageName:   stageName,
		Params:      params,
	}
	mock.lockUpdateStageResources.Lock()
	mock.calls.UpdateStageResources = append(mock.calls.UpdateStageResources, callInfo)
	mock.lockUpdateStageResources.Unlock()
	return mock.UpdateStageResourcesFunc(projectName, stageName, params)
}

// UpdateStageResourcesCalls gets all the calls that were made to UpdateStageResources.
// Check the length with:
//     len(mockedIStageResourceManager.UpdateStageResourcesCalls())
func (mock *IStageResourceManagerMock) UpdateStageResourcesCalls() []struct {
	ProjectName string
	StageName   string
	Params      models.UpdateResourcesParams
} {
	var calls []struct {
		ProjectName string
		StageName   string
		Params      models.UpdateResourcesParams
	}
	mock.lockUpdateStageResources.RLock()
	calls = mock.calls.UpdateStageResources
	mock.lockUpdateStageResources.RUnlock()
	return calls
}