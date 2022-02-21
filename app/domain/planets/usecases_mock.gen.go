// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package planets

import (
	"context"
	"sync"
)

// Ensure, that UseCasesMock does implement UseCases.
// If this is not the case, regenerate this file with moq.
var _ UseCases = &UseCasesMock{}

// UseCasesMock is a mock implementation of UseCases.
//
// 	func TestSomethingThatUsesUseCases(t *testing.T) {
//
// 		// make and configure a mocked UseCases
// 		mockedUseCases := &UseCasesMock{
// 			CreateFunc: func(contextMoqParam context.Context, planet Planet) error {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(contextMoqParam context.Context, iD ID) error {
// 				panic("mock out the Delete method")
// 			},
// 			GetByIDFunc: func(contextMoqParam context.Context, iD ID) (Planet, error) {
// 				panic("mock out the GetByID method")
// 			},
// 			GetByNameFunc: func(ctx context.Context, name string) (Planet, error) {
// 				panic("mock out the GetByName method")
// 			},
// 			ListFunc: func(contextMoqParam context.Context) ([]Planet, error) {
// 				panic("mock out the List method")
// 			},
// 		}
//
// 		// use mockedUseCases in code that requires UseCases
// 		// and then make assertions.
//
// 	}
type UseCasesMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(contextMoqParam context.Context, planet Planet) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(contextMoqParam context.Context, iD ID) error

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(contextMoqParam context.Context, iD ID) (Planet, error)

	// GetByNameFunc mocks the GetByName method.
	GetByNameFunc func(ctx context.Context, name string) (Planet, error)

	// ListFunc mocks the List method.
	ListFunc func(contextMoqParam context.Context) ([]Planet, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Planet is the planet argument value.
			Planet Planet
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// ID is the iD argument value.
			ID ID
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// ID is the iD argument value.
			ID ID
		}
		// GetByName holds details about calls to the GetByName method.
		GetByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// List holds details about calls to the List method.
		List []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockCreate    sync.RWMutex
	lockDelete    sync.RWMutex
	lockGetByID   sync.RWMutex
	lockGetByName sync.RWMutex
	lockList      sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UseCasesMock) Create(contextMoqParam context.Context, planet Planet) error {
	if mock.CreateFunc == nil {
		panic("UseCasesMock.CreateFunc: method is nil but UseCases.Create was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Planet          Planet
	}{
		ContextMoqParam: contextMoqParam,
		Planet:          planet,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(contextMoqParam, planet)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedUseCases.CreateCalls())
func (mock *UseCasesMock) CreateCalls() []struct {
	ContextMoqParam context.Context
	Planet          Planet
} {
	var calls []struct {
		ContextMoqParam context.Context
		Planet          Planet
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *UseCasesMock) Delete(contextMoqParam context.Context, iD ID) error {
	if mock.DeleteFunc == nil {
		panic("UseCasesMock.DeleteFunc: method is nil but UseCases.Delete was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		ID              ID
	}{
		ContextMoqParam: contextMoqParam,
		ID:              iD,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(contextMoqParam, iD)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedUseCases.DeleteCalls())
func (mock *UseCasesMock) DeleteCalls() []struct {
	ContextMoqParam context.Context
	ID              ID
} {
	var calls []struct {
		ContextMoqParam context.Context
		ID              ID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *UseCasesMock) GetByID(contextMoqParam context.Context, iD ID) (Planet, error) {
	if mock.GetByIDFunc == nil {
		panic("UseCasesMock.GetByIDFunc: method is nil but UseCases.GetByID was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		ID              ID
	}{
		ContextMoqParam: contextMoqParam,
		ID:              iD,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(contextMoqParam, iD)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//     len(mockedUseCases.GetByIDCalls())
func (mock *UseCasesMock) GetByIDCalls() []struct {
	ContextMoqParam context.Context
	ID              ID
} {
	var calls []struct {
		ContextMoqParam context.Context
		ID              ID
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// GetByName calls GetByNameFunc.
func (mock *UseCasesMock) GetByName(ctx context.Context, name string) (Planet, error) {
	if mock.GetByNameFunc == nil {
		panic("UseCasesMock.GetByNameFunc: method is nil but UseCases.GetByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockGetByName.Lock()
	mock.calls.GetByName = append(mock.calls.GetByName, callInfo)
	mock.lockGetByName.Unlock()
	return mock.GetByNameFunc(ctx, name)
}

// GetByNameCalls gets all the calls that were made to GetByName.
// Check the length with:
//     len(mockedUseCases.GetByNameCalls())
func (mock *UseCasesMock) GetByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockGetByName.RLock()
	calls = mock.calls.GetByName
	mock.lockGetByName.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *UseCasesMock) List(contextMoqParam context.Context) ([]Planet, error) {
	if mock.ListFunc == nil {
		panic("UseCasesMock.ListFunc: method is nil but UseCases.List was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(contextMoqParam)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedUseCases.ListCalls())
func (mock *UseCasesMock) ListCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}
