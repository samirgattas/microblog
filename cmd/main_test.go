package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/config"
	"github.com/samirgattas/microblog/internal/core/domain"
	inmemorystorage "github.com/samirgattas/microblog/internal/core/lib/customerror/in_memory_storage"
	"github.com/stretchr/testify/assert"
)

var (
	c *config.Config
	r *gin.Engine
)

func initTest() *gin.Engine {
	userDB := inmemorystorage.NewStore()
	c = &config.Config{}
	c = c.NewConfig(userDB)
	handler := Container(c)
	r := gin.Default()
	Routes(r, handler)
	return r
}

// PING - GET /ping

func TestPing_Ok(t *testing.T) {
	r := initTest()

	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, `{"message":"pong"}`, string(responseData))
}

// SAVE USER - POST /users

func TestSaveUser_Ok(t *testing.T) {
	r := initTest()

	body := `{
		"id": 5,
		"nickname": "myNickname"
	}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	responseUser := domain.User{}
	err = json.Unmarshal(responseData, &responseUser)
	assert.Nil(t, err)
	expectedResp := domain.User{ID: 5, Nickname: "myNickname", CreatedAt: responseUser.CreatedAt}
	assert.Equal(t, expectedResp, responseUser)
}

// GET USER - GET /users/:id

func TestGetUser_Ok(t *testing.T) {
	r := initTest()

	body := `{
		"id": 1,
		"nickname": "samir"
	}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	responseUser := domain.User{}
	err = json.Unmarshal(responseData, &responseUser)
	assert.Nil(t, err)
	expectedResp := domain.User{ID: 1, Nickname: "samir", CreatedAt: responseUser.CreatedAt}
	assert.Equal(t, expectedResp, responseUser)
}

func TestGetUser_NotFoundError(t *testing.T) {
	r := initTest()

	req, _ := http.NewRequest("GET", "/users/100", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	responseData, _ := io.ReadAll(w.Body)
	expected := `{"message":"Message: user not found, Status: 404","status":404}`
	assert.Equal(t, expected, string(responseData))
}

// CREATE FOLLOWED - POST /followed

func TestSaveFollowed_Ok(t *testing.T) {
	r := initTest()

	// Create user_id:11
	{
		body := `{
			"id": 11,
			"nickname": "user11"
		}`
		req, _ := http.NewRequest("POST", "/users", bytes.NewReader([]byte(body)))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

	// Create user_id:12
	{
		body := `{
			"id": 12,
			"nickname": "user12"
		}`
		req, _ := http.NewRequest("POST", "/users", bytes.NewReader([]byte(body)))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

	// user_id:11 follows user_id:12
	followedID := int64(0)
	{
		body := `{
			"user_id": 11,
			"followed_user_id": 12
		}`
		req, _ := http.NewRequest("POST", "/followed", bytes.NewReader([]byte(body)))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		responseData, err := io.ReadAll(w.Body)
		assert.Nil(t, err)
		respFollowed := domain.Followed{}
		err = json.Unmarshal(responseData, &respFollowed)
		assert.Nil(t, err)
		followedID = respFollowed.ID
		expectedResp := domain.Followed{ID: followedID, UserID: 11, FollowedUserID: 12, Enabled: true, CreatedAt: respFollowed.CreatedAt, UpdatedAt: respFollowed.UpdatedAt}
		assert.Equal(t, expectedResp, respFollowed)
	}

	// Get followed
	{
		req, _ := http.NewRequest("GET", fmt.Sprintf("/followed/%d", followedID), nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		responseData, err := io.ReadAll(w.Body)
		assert.Nil(t, err)
		respFollowed := domain.Followed{}
		err = json.Unmarshal(responseData, &respFollowed)
		assert.Nil(t, err)
		followedID = respFollowed.ID
		expectedResp := domain.Followed{ID: followedID, UserID: 11, FollowedUserID: 12, Enabled: true, CreatedAt: respFollowed.CreatedAt, UpdatedAt: respFollowed.UpdatedAt}
		assert.Equal(t, expectedResp, respFollowed)
	}

	// user_id:11 unfollows user_id:12
	{

		body := `{
			"enabled": false
		}`
		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/followed/%d", followedID), bytes.NewReader([]byte(body)))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		responseData, err := io.ReadAll(w.Body)
		assert.Nil(t, err)
		respFollowed := domain.Followed{}
		err = json.Unmarshal(responseData, &respFollowed)
		assert.Nil(t, err)
		followedID = respFollowed.ID
		expectedResp := domain.Followed{ID: followedID, UserID: 11, FollowedUserID: 12, Enabled: false, CreatedAt: respFollowed.CreatedAt, UpdatedAt: respFollowed.UpdatedAt}
		assert.Equal(t, expectedResp, respFollowed)
	}
}
