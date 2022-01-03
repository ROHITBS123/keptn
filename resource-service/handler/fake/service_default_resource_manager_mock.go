// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler_mock

import (
	"sync"
)

// IServiceDefaultResourceManagerMock is a mock implementation of handler.IServiceDefaultResourceManager.
//
// 	func TestSomethingThatUsesIServiceDefaultResourceManager(t *testing.T) {
//
// 		// make and configure a mocked handler.IServiceDefaultResourceManager
// 		mockedIServiceDefaultResourceManager := &IServiceDefaultResourceManagerMock{
// 			CreateServiceDefaultResourcesFunc: func()  {
// 				panic("mock out the CreateServiceDefaultResources method")
// 			},
// 		}
//
// 		// use mockedIServiceDefaultResourceManager in code that requires handler.IServiceDefaultResourceManager
// 		// and then make assertions.
//
// 	}
type IServiceDefaultResourceManagerMock struct {
	// CreateServiceDefaultResourcesFunc mocks the CreateServiceDefaultResources method.
	CreateServiceDefaultResourcesFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// CreateServiceDefaultResources holds details about calls to the CreateServiceDefaultResources method.
		CreateServiceDefaultResources []struct {
		}
	}
	lockCreateServiceDefaultResources sync.RWMutex
}

// CreateServiceDefaultResources calls CreateServiceDefaultResourcesFunc.
func (mock *IServiceDefaultResourceManagerMock) CreateServiceDefaultResources() {
	if mock.CreateServiceDefaultResourcesFunc == nil {
		panic("IServiceDefaultResourceManagerMock.CreateServiceDefaultResourcesFunc: method is nil but IServiceDefaultResourceManager.CreateServiceDefaultResources was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCreateServiceDefaultResources.Lock()
	mock.calls.CreateServiceDefaultResources = append(mock.calls.CreateServiceDefaultResources, callInfo)
	mock.lockCreateServiceDefaultResources.Unlock()
	mock.CreateServiceDefaultResourcesFunc()
}

// CreateServiceDefaultResourcesCalls gets all the calls that were made to CreateServiceDefaultResources.
// Check the length with:
//     len(mockedIServiceDefaultResourceManager.CreateServiceDefaultResourcesCalls())
func (mock *IServiceDefaultResourceManagerMock) CreateServiceDefaultResourcesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCreateServiceDefaultResources.RLock()
	calls = mock.calls.CreateServiceDefaultResources
	mock.lockCreateServiceDefaultResources.RUnlock()
	return calls
}