// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/aliml92/anor"
	"sync"
)

// Ensure, that UserServiceMock does implement anor.UserService.
// If this is not the case, regenerate this file with moq.
var _ anor.UserService = &UserServiceMock{}

// UserServiceMock is a mock implementation of anor.UserService.
//
//	func TestSomethingThatUsesUserService(t *testing.T) {
//
//		// make and configure a mocked anor.UserService
//		mockedUserService := &UserServiceMock{
//			CreateUserFunc: func(ctx context.Context, user anor.User) error {
//				panic("mock out the Create method")
//			},
//			GetUserFunc: func(ctx context.Context, id int64) (anor.User, error) {
//				panic("mock out the getUser method")
//			},
//			GetUserActivityCountsFunc: func(ctx context.Context, id int64) (anor.UserActivityCounts, error) {
//				panic("mock out the GetActivityCounts method")
//			},
//			GetUserByEmailFunc: func(ctx context.Context, email string) (anor.User, error) {
//				panic("mock out the GetByEmail method")
//			},
//			UpdateUserPasswordFunc: func(ctx context.Context, id int64, password string) error {
//				panic("mock out the UpdatePassword method")
//			},
//			UpdateUserStatusByEmailFunc: func(ctx context.Context, status anor.UserStatus, email string) error {
//				panic("mock out the UpdateStatusByEmail method")
//			},
//		}
//
//		// use mockedUserService in code that requires anor.UserService
//		// and then make assertions.
//
//	}
type UserServiceMock struct {
	// CreateUserFunc mocks the Create method.
	CreateUserFunc func(ctx context.Context, user anor.User) error

	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context, id int64) (anor.User, error)

	// GetUserActivityCountsFunc mocks the GetUserActivityCounts method.
	GetUserActivityCountsFunc func(ctx context.Context, id int64) (anor.UserActivityCounts, error)

	// GetUserByEmailFunc mocks the GetUserByEmail method.
	GetUserByEmailFunc func(ctx context.Context, email string) (anor.User, error)

	// UpdateUserPasswordFunc mocks the UpdateUserPassword method.
	UpdateUserPasswordFunc func(ctx context.Context, id int64, password string) error

	// UpdateUserStatusByEmailFunc mocks the UpdateUserStatusByEmail method.
	UpdateUserStatusByEmailFunc func(ctx context.Context, status anor.UserStatus, email string) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		CreateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User anor.User
		}
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// GetUserActivityCounts holds details about calls to the GetUserActivityCounts method.
		GetUserActivityCounts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// GetUserByEmail holds details about calls to the GetUserByEmail method.
		GetUserByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
		// UpdateUserPassword holds details about calls to the UpdateUserPassword method.
		UpdateUserPassword []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
			// Password is the password argument value.
			Password string
		}
		// UpdateUserStatusByEmail holds details about calls to the UpdateUserStatusByEmail method.
		UpdateUserStatusByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Status is the status argument value.
			Status anor.UserStatus
			// Email is the email argument value.
			Email string
		}
	}
	lockCreateUser              sync.RWMutex
	lockGetUser                 sync.RWMutex
	lockGetUserActivityCounts   sync.RWMutex
	lockGetUserByEmail          sync.RWMutex
	lockUpdateUserPassword      sync.RWMutex
	lockUpdateUserStatusByEmail sync.RWMutex
}

// Create calls CreateUserFunc.
func (mock *UserServiceMock) Create(ctx context.Context, user anor.User) error {
	if mock.CreateUserFunc == nil {
		panic("UserServiceMock.CreateUserFunc: method is nil but UserService.Create was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User anor.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	return mock.CreateUserFunc(ctx, user)
}

// CreateUserCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedUserService.CreateUserCalls())
func (mock *UserServiceMock) CreateUserCalls() []struct {
	Ctx  context.Context
	User anor.User
} {
	var calls []struct {
		Ctx  context.Context
		User anor.User
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}

// GetUser calls GetUserFunc.
func (mock *UserServiceMock) GetUser(ctx context.Context, id int64) (anor.User, error) {
	if mock.GetUserFunc == nil {
		panic("UserServiceMock.GetUserFunc: method is nil but UserService.getUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	mock.lockGetUser.Unlock()
	return mock.GetUserFunc(ctx, id)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//
//	len(mockedUserService.GetUserCalls())
func (mock *UserServiceMock) GetUserCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	mock.lockGetUser.RLock()
	calls = mock.calls.GetUser
	mock.lockGetUser.RUnlock()
	return calls
}

// GetUserActivityCounts calls GetUserActivityCountsFunc.
func (mock *UserServiceMock) GetUserActivityCounts(ctx context.Context, id int64) (anor.UserActivityCounts, error) {
	if mock.GetUserActivityCountsFunc == nil {
		panic("UserServiceMock.GetUserActivityCountsFunc: method is nil but UserService.GetActivityCounts was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetUserActivityCounts.Lock()
	mock.calls.GetUserActivityCounts = append(mock.calls.GetUserActivityCounts, callInfo)
	mock.lockGetUserActivityCounts.Unlock()
	return mock.GetUserActivityCountsFunc(ctx, id)
}

// GetUserActivityCountsCalls gets all the calls that were made to GetUserActivityCounts.
// Check the length with:
//
//	len(mockedUserService.GetUserActivityCountsCalls())
func (mock *UserServiceMock) GetUserActivityCountsCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	mock.lockGetUserActivityCounts.RLock()
	calls = mock.calls.GetUserActivityCounts
	mock.lockGetUserActivityCounts.RUnlock()
	return calls
}

// GetUserByEmail calls GetUserByEmailFunc.
func (mock *UserServiceMock) GetUserByEmail(ctx context.Context, email string) (anor.User, error) {
	if mock.GetUserByEmailFunc == nil {
		panic("UserServiceMock.GetUserByEmailFunc: method is nil but UserService.GetByEmail was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	mock.lockGetUserByEmail.Lock()
	mock.calls.GetUserByEmail = append(mock.calls.GetUserByEmail, callInfo)
	mock.lockGetUserByEmail.Unlock()
	return mock.GetUserByEmailFunc(ctx, email)
}

// GetUserByEmailCalls gets all the calls that were made to GetUserByEmail.
// Check the length with:
//
//	len(mockedUserService.GetUserByEmailCalls())
func (mock *UserServiceMock) GetUserByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	mock.lockGetUserByEmail.RLock()
	calls = mock.calls.GetUserByEmail
	mock.lockGetUserByEmail.RUnlock()
	return calls
}

// UpdateUserPassword calls UpdateUserPasswordFunc.
func (mock *UserServiceMock) UpdateUserPassword(ctx context.Context, id int64, password string) error {
	if mock.UpdateUserPasswordFunc == nil {
		panic("UserServiceMock.UpdateUserPasswordFunc: method is nil but UserService.UpdatePassword was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ID       int64
		Password string
	}{
		Ctx:      ctx,
		ID:       id,
		Password: password,
	}
	mock.lockUpdateUserPassword.Lock()
	mock.calls.UpdateUserPassword = append(mock.calls.UpdateUserPassword, callInfo)
	mock.lockUpdateUserPassword.Unlock()
	return mock.UpdateUserPasswordFunc(ctx, id, password)
}

// UpdateUserPasswordCalls gets all the calls that were made to UpdateUserPassword.
// Check the length with:
//
//	len(mockedUserService.UpdateUserPasswordCalls())
func (mock *UserServiceMock) UpdateUserPasswordCalls() []struct {
	Ctx      context.Context
	ID       int64
	Password string
} {
	var calls []struct {
		Ctx      context.Context
		ID       int64
		Password string
	}
	mock.lockUpdateUserPassword.RLock()
	calls = mock.calls.UpdateUserPassword
	mock.lockUpdateUserPassword.RUnlock()
	return calls
}

// UpdateUserStatusByEmail calls UpdateUserStatusByEmailFunc.
func (mock *UserServiceMock) UpdateUserStatusByEmail(ctx context.Context, status anor.UserStatus, email string) error {
	if mock.UpdateUserStatusByEmailFunc == nil {
		panic("UserServiceMock.UpdateUserStatusByEmailFunc: method is nil but UserService.UpdateStatusByEmail was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Status anor.UserStatus
		Email  string
	}{
		Ctx:    ctx,
		Status: status,
		Email:  email,
	}
	mock.lockUpdateUserStatusByEmail.Lock()
	mock.calls.UpdateUserStatusByEmail = append(mock.calls.UpdateUserStatusByEmail, callInfo)
	mock.lockUpdateUserStatusByEmail.Unlock()
	return mock.UpdateUserStatusByEmailFunc(ctx, status, email)
}

// UpdateUserStatusByEmailCalls gets all the calls that were made to UpdateUserStatusByEmail.
// Check the length with:
//
//	len(mockedUserService.UpdateUserStatusByEmailCalls())
func (mock *UserServiceMock) UpdateUserStatusByEmailCalls() []struct {
	Ctx    context.Context
	Status anor.UserStatus
	Email  string
} {
	var calls []struct {
		Ctx    context.Context
		Status anor.UserStatus
		Email  string
	}
	mock.lockUpdateUserStatusByEmail.RLock()
	calls = mock.calls.UpdateUserStatusByEmail
	mock.lockUpdateUserStatusByEmail.RUnlock()
	return calls
}
