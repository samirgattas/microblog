// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"

	"github.com/samirgattas/microblog/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// NewMockTweetService creates a new instance of MockTweetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTweetService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTweetService {
	mock := &MockTweetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockTweetService is an autogenerated mock type for the TweetService type
type MockTweetService struct {
	mock.Mock
}

type MockTweetService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTweetService) EXPECT() *MockTweetService_Expecter {
	return &MockTweetService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function for the type MockTweetService
func (_mock *MockTweetService) Create(ctx context.Context, tweet *domain.Tweet) error {
	ret := _mock.Called(ctx, tweet)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *domain.Tweet) error); ok {
		r0 = returnFunc(ctx, tweet)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockTweetService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTweetService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - tweet *domain.Tweet
func (_e *MockTweetService_Expecter) Create(ctx interface{}, tweet interface{}) *MockTweetService_Create_Call {
	return &MockTweetService_Create_Call{Call: _e.mock.On("Create", ctx, tweet)}
}

func (_c *MockTweetService_Create_Call) Run(run func(ctx context.Context, tweet *domain.Tweet)) *MockTweetService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 *domain.Tweet
		if args[1] != nil {
			arg1 = args[1].(*domain.Tweet)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockTweetService_Create_Call) Return(err error) *MockTweetService_Create_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTweetService_Create_Call) RunAndReturn(run func(ctx context.Context, tweet *domain.Tweet) error) *MockTweetService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function for the type MockTweetService
func (_mock *MockTweetService) Get(ctx context.Context, tweetID int64) (*domain.Tweet, error) {
	ret := _mock.Called(ctx, tweetID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *domain.Tweet
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) (*domain.Tweet, error)); ok {
		return returnFunc(ctx, tweetID)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) *domain.Tweet); ok {
		r0 = returnFunc(ctx, tweetID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Tweet)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = returnFunc(ctx, tweetID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockTweetService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockTweetService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - tweetID int64
func (_e *MockTweetService_Expecter) Get(ctx interface{}, tweetID interface{}) *MockTweetService_Get_Call {
	return &MockTweetService_Get_Call{Call: _e.mock.On("Get", ctx, tweetID)}
}

func (_c *MockTweetService_Get_Call) Run(run func(ctx context.Context, tweetID int64)) *MockTweetService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 int64
		if args[1] != nil {
			arg1 = args[1].(int64)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockTweetService_Get_Call) Return(tweet *domain.Tweet, err error) *MockTweetService_Get_Call {
	_c.Call.Return(tweet, err)
	return _c
}

func (_c *MockTweetService_Get_Call) RunAndReturn(run func(ctx context.Context, tweetID int64) (*domain.Tweet, error)) *MockTweetService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function for the type MockTweetService
func (_mock *MockTweetService) Search(ctx context.Context, followerUserID int64, limit int64, offset int64) (domain.TweetsSearchResult, error) {
	ret := _mock.Called(ctx, followerUserID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 domain.TweetsSearchResult
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64, int64, int64) (domain.TweetsSearchResult, error)); ok {
		return returnFunc(ctx, followerUserID, limit, offset)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64, int64, int64) domain.TweetsSearchResult); ok {
		r0 = returnFunc(ctx, followerUserID, limit, offset)
	} else {
		r0 = ret.Get(0).(domain.TweetsSearchResult)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, int64, int64, int64) error); ok {
		r1 = returnFunc(ctx, followerUserID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockTweetService_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockTweetService_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - followerUserID int64
//   - limit int64
//   - offset int64
func (_e *MockTweetService_Expecter) Search(ctx interface{}, followerUserID interface{}, limit interface{}, offset interface{}) *MockTweetService_Search_Call {
	return &MockTweetService_Search_Call{Call: _e.mock.On("Search", ctx, followerUserID, limit, offset)}
}

func (_c *MockTweetService_Search_Call) Run(run func(ctx context.Context, followerUserID int64, limit int64, offset int64)) *MockTweetService_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 int64
		if args[1] != nil {
			arg1 = args[1].(int64)
		}
		var arg2 int64
		if args[2] != nil {
			arg2 = args[2].(int64)
		}
		var arg3 int64
		if args[3] != nil {
			arg3 = args[3].(int64)
		}
		run(
			arg0,
			arg1,
			arg2,
			arg3,
		)
	})
	return _c
}

func (_c *MockTweetService_Search_Call) Return(tweetsSearchResult domain.TweetsSearchResult, err error) *MockTweetService_Search_Call {
	_c.Call.Return(tweetsSearchResult, err)
	return _c
}

func (_c *MockTweetService_Search_Call) RunAndReturn(run func(ctx context.Context, followerUserID int64, limit int64, offset int64) (domain.TweetsSearchResult, error)) *MockTweetService_Search_Call {
	_c.Call.Return(run)
	return _c
}
