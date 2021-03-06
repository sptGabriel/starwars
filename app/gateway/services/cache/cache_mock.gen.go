// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cache

import (
	"sync"
)

// Ensure, that CacheMock does implement Cache.
// If this is not the case, regenerate this file with moq.
var _ Cache = &CacheMock{}

// CacheMock is a mock implementation of Cache.
//
// 	func TestSomethingThatUsesCache(t *testing.T) {
//
// 		// make and configure a mocked Cache
// 		mockedCache := &CacheMock{
// 			GetFunc: func(key string) (interface{}, error) {
// 				panic("mock out the Get method")
// 			},
// 			SaveFunc: func(key string, value []byte) error {
// 				panic("mock out the Save method")
// 			},
// 		}
//
// 		// use mockedCache in code that requires Cache
// 		// and then make assertions.
//
// 	}
type CacheMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(key string) (interface{}, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(key string, value []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Key is the key argument value.
			Key string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value []byte
		}
	}
	lockGet  sync.RWMutex
	lockSave sync.RWMutex
}

// Get calls GetFunc.
func (mock *CacheMock) Get(key string) (interface{}, error) {
	if mock.GetFunc == nil {
		panic("CacheMock.GetFunc: method is nil but Cache.Get was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(key)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedCache.GetCalls())
func (mock *CacheMock) GetCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *CacheMock) Save(key string, value []byte) error {
	if mock.SaveFunc == nil {
		panic("CacheMock.SaveFunc: method is nil but Cache.Save was just called")
	}
	callInfo := struct {
		Key   string
		Value []byte
	}{
		Key:   key,
		Value: value,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(key, value)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedCache.SaveCalls())
func (mock *CacheMock) SaveCalls() []struct {
	Key   string
	Value []byte
} {
	var calls []struct {
		Key   string
		Value []byte
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
