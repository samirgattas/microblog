package user

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/internal/core/domain"
	mocks "github.com/samirgattas/microblog/internal/mock"
	"github.com/stretchr/testify/assert"
)

// SAVE USER

func TestSave_Ok(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	user := &domain.User{ID: 1, Nickname: "nickname"}

	userRepository.EXPECT().Save(ctx, user).Return(nil)

	s := NewUserService(userRepository)

	err := s.Create(ctx, user)
	assert.Nil(t, err)
	userRepository.AssertExpectations(t)
}

func TestSave_SaveUserError(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	user := &domain.User{ID: 1, Nickname: "nickname"}
	userRepository.EXPECT().Save(ctx, user).Return(errors.New("save_user_error"))

	s := NewUserService(userRepository)
	err := s.Create(ctx, user)

	assert.NotNil(t, err)
	assert.Equal(t, "save_user_error", err.Error())
	userRepository.AssertExpectations(t)
}

// GET USER

func TestGet_Ok(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	now := time.Now()
	userID := int64(1)
	userRepository.EXPECT().Get(ctx, userID).Return(&domain.User{ID: userID, Nickname: "nickname", CreatedAt: &now}, nil)

	s := NewUserService(userRepository)
	user, err := s.Get(ctx, userID)

	assert.Nil(t, err)
	expecterUser := &domain.User{ID: userID, Nickname: "nickname", CreatedAt: &now}
	assert.Equal(t, expecterUser, user)
	userRepository.AssertExpectations(t)
}

func TestGet_GetUserError(t *testing.T) {
	userRepository := &mocks.MockUserRepository{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	userID := int64(1)
	userRepository.EXPECT().Get(ctx, userID).Return(&domain.User{}, errors.New("get_user_error"))

	s := NewUserService(userRepository)
	_, err := s.Get(ctx, userID)

	assert.NotNil(t, err)
	assert.Equal(t, "get_user_error", err.Error())
	userRepository.AssertExpectations(t)
}
