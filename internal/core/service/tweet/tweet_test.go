package tweet

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/internal/core/domain"
	mocks "github.com/samirgattas/microblog/internal/mock"
	"github.com/samirgattas/microblog/lib/customerror"
	"github.com/stretchr/testify/assert"
)

// CREATE FOLLOWED

func TestCreate_Ok(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	followedRepository := &mocks.MockFollowedRepository{}
	tweetRepository := &mocks.MockTweetRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tweet := &domain.Tweet{UserID: 1, Post: "my post"}

	userRepository.EXPECT().Get(ctx, tweet.UserID).Return(&domain.User{ID: tweet.UserID}, nil)
	tweetRepository.EXPECT().Save(ctx, tweet).Return(nil)

	s := NewTweetService(tweetRepository, userRepository, followedRepository)
	err := s.Create(ctx, tweet)
	assert.Nil(t, err)
	userRepository.AssertExpectations(t)
	followedRepository.AssertExpectations(t)
	tweetRepository.AssertExpectations(t)
}

func TestCreate_PostTooLongError(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	followedRepository := &mocks.MockFollowedRepository{}
	tweetRepository := &mocks.MockTweetRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tweet := &domain.Tweet{UserID: 1, Post: "my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post my post"}

	s := NewTweetService(tweetRepository, userRepository, followedRepository)
	err := s.Create(ctx, tweet)
	assert.NotNil(t, err)
	assert.Equal(t, "Message: post too long, Status: 400", err.Error())
	userRepository.AssertExpectations(t)
	followedRepository.AssertExpectations(t)
	tweetRepository.AssertExpectations(t)
}

func TestCreate_GetUserError(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	followedRepository := &mocks.MockFollowedRepository{}
	tweetRepository := &mocks.MockTweetRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tweet := &domain.Tweet{UserID: 1, Post: "my post"}

	userRepository.EXPECT().Get(ctx, tweet.UserID).Return(&domain.User{}, errors.New("get_user_error"))

	s := NewTweetService(tweetRepository, userRepository, followedRepository)
	err := s.Create(ctx, tweet)
	assert.NotNil(t, err)
	assert.Equal(t, "get_user_error", err.Error())
	userRepository.AssertExpectations(t)
	followedRepository.AssertExpectations(t)
	tweetRepository.AssertExpectations(t)
}

func TestCreate_UserNotFoundError(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	followedRepository := &mocks.MockFollowedRepository{}
	tweetRepository := &mocks.MockTweetRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tweet := &domain.Tweet{UserID: 1, Post: "my post"}

	userRepository.EXPECT().Get(ctx, tweet.UserID).Return(&domain.User{}, customerror.NewNotFoundError("user"))

	s := NewTweetService(tweetRepository, userRepository, followedRepository)
	err := s.Create(ctx, tweet)
	assert.NotNil(t, err)
	assert.Equal(t, "Message: invalid user_id, Status: 400", err.Error())
	userRepository.AssertExpectations(t)
	followedRepository.AssertExpectations(t)
	tweetRepository.AssertExpectations(t)
}

func TestCreate_SaveTweetError(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	followedRepository := &mocks.MockFollowedRepository{}
	tweetRepository := &mocks.MockTweetRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tweet := &domain.Tweet{UserID: 1, Post: "my post"}

	userRepository.EXPECT().Get(ctx, tweet.UserID).Return(&domain.User{ID: tweet.UserID}, nil)
	tweetRepository.EXPECT().Save(ctx, tweet).Return(errors.New("save_tweet_error"))

	s := NewTweetService(tweetRepository, userRepository, followedRepository)
	err := s.Create(ctx, tweet)
	assert.NotNil(t, err)
	assert.Equal(t, "save_tweet_error", err.Error())
	userRepository.AssertExpectations(t)
	followedRepository.AssertExpectations(t)
	tweetRepository.AssertExpectations(t)
}
