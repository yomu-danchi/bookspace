// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/domain/repositories"
	"sync"
)

// Ensure, that RepositoryMock does implement repositories.Repository.
// If this is not the case, regenerate this file with moq.
var _ repositories.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of repositories.Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked repositories.Repository
// 		mockedRepository := &RepositoryMock{
// 			SaveUserFunc: func(userMoqParam user.User) error {
// 				panic("mock out the SaveUser method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires repositories.Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// SaveUserFunc mocks the SaveUser method.
	SaveUserFunc func(userMoqParam user.User) error

	// calls tracks calls to the methods.
	calls struct {
		// SaveUser holds details about calls to the SaveUser method.
		SaveUser []struct {
			// UserMoqParam is the userMoqParam argument value.
			UserMoqParam user.User
		}
	}
	lockSaveUser sync.RWMutex
}

// SaveUser calls SaveUserFunc.
func (mock *RepositoryMock) SaveUser(userMoqParam user.User) error {
	if mock.SaveUserFunc == nil {
		panic("RepositoryMock.SaveUserFunc: method is nil but Repository.SaveUser was just called")
	}
	callInfo := struct {
		UserMoqParam user.User
	}{
		UserMoqParam: userMoqParam,
	}
	mock.lockSaveUser.Lock()
	mock.calls.SaveUser = append(mock.calls.SaveUser, callInfo)
	mock.lockSaveUser.Unlock()
	return mock.SaveUserFunc(userMoqParam)
}

// SaveUserCalls gets all the calls that were made to SaveUser.
// Check the length with:
//     len(mockedRepository.SaveUserCalls())
func (mock *RepositoryMock) SaveUserCalls() []struct {
	UserMoqParam user.User
} {
	var calls []struct {
		UserMoqParam user.User
	}
	mock.lockSaveUser.RLock()
	calls = mock.calls.SaveUser
	mock.lockSaveUser.RUnlock()
	return calls
}