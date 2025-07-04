package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/config"
	"github.com/samirgattas/microblog/internal/core/domain"
	inmemorystore "github.com/samirgattas/microblog/lib/in_memory_store"
	"github.com/stretchr/testify/assert"
)

var (
	c          *config.Config
	r          *gin.Engine
	userDB     inmemorystore.Store
	followedDB map[int64]domain.Followed
	tweetDB    map[int64]domain.Tweet
)

func initTest() *gin.Engine {
	userDB = inmemorystore.NewStore()
	followedDB = make(map[int64]domain.Followed)
	tweetDB = make(map[int64]domain.Tweet)
	c = &config.Config{}
	c = c.NewConfig(userDB, followedDB, tweetDB)
	handler := Container(c)
	r := gin.Default()
	Routes(r, handler)
	return r
}

func CleanDBs() {
	userDB.Drop()
	followedDB = make(map[int64]domain.Followed)
	tweetDB = make(map[int64]domain.Tweet)
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
	defer CleanDBs()

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

	_, err = userDB.Get(5)
	assert.Nil(t, err)
}

// GET USER - GET /users/:id

func TestGetUser_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

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

	req, _ := http.NewRequest("GET", "/users/1", nil)
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
	defer CleanDBs()

	userDB.SaveWithID(11, domain.User{ID: 11, Nickname: "user11"})
	userDB.SaveWithID(12, domain.User{ID: 12, Nickname: "user12"})

	// user_id:11 follows user_id:12
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
	expectedResp := domain.Followed{ID: respFollowed.ID, UserID: 11, FollowedUserID: 12, Enabled: true, CreatedAt: respFollowed.CreatedAt, UpdatedAt: respFollowed.UpdatedAt}
	assert.Equal(t, expectedResp, respFollowed)
}

// GET FOLLOWED - GET /followed/:followed_id

func TestGetFollowed_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	followedDB[1] = domain.Followed{ID: 1, UserID: 11, FollowedUserID: 12, Enabled: true}

	req, _ := http.NewRequest("GET", "/followed/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	respFollowed := domain.Followed{}
	err = json.Unmarshal(responseData, &respFollowed)
	assert.Nil(t, err)
	expectedResp := domain.Followed{ID: respFollowed.ID, UserID: 11, FollowedUserID: 12, Enabled: true, CreatedAt: respFollowed.CreatedAt, UpdatedAt: respFollowed.UpdatedAt}
	assert.Equal(t, expectedResp, respFollowed)
}

// UPDATE FOLLOWED - PATCH /followed/:followed_id

func TestUpdateFollowed_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	followedDB[1] = domain.Followed{ID: 1, UserID: 11, FollowedUserID: 12, Enabled: true}

	// user_id:11 unfollows user_id:12
	body := `{
		"enabled": false
	}`
	req, _ := http.NewRequest("PATCH", "/followed/1", bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	respFollowed := domain.Followed{}
	err = json.Unmarshal(responseData, &respFollowed)
	assert.Nil(t, err)
	expectedResp := domain.Followed{ID: respFollowed.ID, UserID: 11, FollowedUserID: 12, Enabled: false, CreatedAt: respFollowed.CreatedAt, UpdatedAt: respFollowed.UpdatedAt}
	assert.Equal(t, expectedResp, respFollowed)
}

// CREATE TWEET - POST /tweets

func TestCreateTweet_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	userDB.SaveWithID(12345, domain.User{ID: 12345, Nickname: "nickname"})

	// user_id:11 unfollows user_id:12
	body := `{
		"user_id": 12345,
		"post": "My post"
	}`
	req, _ := http.NewRequest("POST", "/tweets", bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	respTweet := domain.Tweet{}
	err = json.Unmarshal(responseData, &respTweet)
	assert.Nil(t, err)
	expectedResp := domain.Tweet{ID: respTweet.ID, UserID: 12345, Post: "My post", CreatedAt: respTweet.CreatedAt}
	assert.Equal(t, expectedResp, respTweet)
}

func TestCreateTweet_PostTooLongError(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	userDB.SaveWithID(12345, domain.User{ID: 12345, Nickname: "nickname"})

	// user_id:11 unfollows user_id:12
	post := strings.Repeat("a", 250)
	body := fmt.Sprintf(`{
		"user_id": 12345,
		"post": "%s"
	}`, post)
	req, _ := http.NewRequest("POST", "/tweets", bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"message":"Message: post too long, Status: 400","status":400}`, string(responseData))
}

// GET TWEET - GET /tweets/:tweet_id

func TestGetTweet_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	tweetDB[100] = domain.Tweet{ID: 100, UserID: 12, Post: "Tweet 100!"}

	req, _ := http.NewRequest("GET", "/tweets/100", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	expectedResp := `{"id":100,"user_id":12,"post":"Tweet 100!","created_at":null}`
	assert.Equal(t, expectedResp, string(responseData))
}

// SEARCH TWEETS - GET /tweets

func TestSearchTweets_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	followedDB[1] = domain.Followed{ID: 1, UserID: 11, FollowedUserID: 12, Enabled: true}
	followedDB[2] = domain.Followed{ID: 2, UserID: 11, FollowedUserID: 13, Enabled: true}
	tweetDB[1] = domain.Tweet{ID: 1, UserID: 12, Post: "Tweet 1 of user_id: 12"}
	tweetDB[2] = domain.Tweet{ID: 2, UserID: 12, Post: "Tweet 2 of user_id: 12"}
	tweetDB[3] = domain.Tweet{ID: 3, UserID: 13, Post: "Tweet 1 of user_id: 13"}
	tweetDB[4] = domain.Tweet{ID: 4, UserID: 13, Post: "Tweet 2 of user_id: 13"}

	req, _ := http.NewRequest("GET", "/tweets?user_id=11&limit=3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	expectedResp := `{"paging":{"total":4,"limit":3,"offset":0},"results":[{"id":4,"user_id":13,"post":"Tweet 2 of user_id: 13","created_at":null},{"id":3,"user_id":13,"post":"Tweet 1 of user_id: 13","created_at":null},{"id":2,"user_id":12,"post":"Tweet 2 of user_id: 12","created_at":null}]}`
	assert.Equal(t, expectedResp, string(responseData))
}

func TestSearchTweets_UsePaging_Ok(t *testing.T) {
	r := initTest()
	defer CleanDBs()

	followedDB[1] = domain.Followed{ID: 1, UserID: 11, FollowedUserID: 12, Enabled: true}
	followedDB[2] = domain.Followed{ID: 2, UserID: 11, FollowedUserID: 13, Enabled: true}
	tweetDB[1] = domain.Tweet{ID: 1, UserID: 12, Post: "Tweet 1 of user_id: 12"}
	tweetDB[2] = domain.Tweet{ID: 2, UserID: 12, Post: "Tweet 2 of user_id: 12"}
	tweetDB[3] = domain.Tweet{ID: 3, UserID: 13, Post: "Tweet 1 of user_id: 13"}
	tweetDB[4] = domain.Tweet{ID: 4, UserID: 13, Post: "Tweet 2 of user_id: 13"}
	// Search for user_id: 11 with offset:0 and limit:3
	{
		req, _ := http.NewRequest("GET", "/tweets?user_id=11&limit=3", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		responseData, err := io.ReadAll(w.Body)
		assert.Nil(t, err)
		expectedResp := `{"paging":{"total":4,"limit":3,"offset":0},"results":[{"id":4,"user_id":13,"post":"Tweet 2 of user_id: 13","created_at":null},{"id":3,"user_id":13,"post":"Tweet 1 of user_id: 13","created_at":null},{"id":2,"user_id":12,"post":"Tweet 2 of user_id: 12","created_at":null}]}`
		assert.Equal(t, expectedResp, string(responseData))
	}
	// Search for user_id: 11 with offset:3 and limit:3 to get the first tweet
	{
		req, _ := http.NewRequest("GET", "/tweets?user_id=11&limit=3&offset=3", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		responseData, err := io.ReadAll(w.Body)
		assert.Nil(t, err)
		expectedResp := `{"paging":{"total":4,"limit":3,"offset":3},"results":[{"id":1,"user_id":12,"post":"Tweet 1 of user_id: 12","created_at":null}]}`
		assert.Equal(t, expectedResp, string(responseData))
	}
}
