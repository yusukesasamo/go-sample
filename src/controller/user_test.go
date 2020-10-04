package controller

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yusukesasamo/go-sample/src/model"
)

var mockUser = model.User{
	ID:        1,
	Mail:      "sasamo@sasamo.com",
	Password:  "xxxxxxxxx",
	Authkey:   "xxxxxxxxxxxx",
	Point:     10000,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestUsersGET(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	UsersGET(c)
	assert.Equal(t, 200, w.Code)
}

func TestFindUserByID(t *testing.T) {
	user := FindUserByID(mockUser.ID)
	assert.NotNil(t, user)
	assert.IsType(t, mockUser, user)
}
