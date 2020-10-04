package controller

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yusukesasamo/go-sample/src/model"
)

var mockItem = model.Item{
	ID:        1,
	Name:      "Sasamon",
	UserID:    1,
	Price:     1000,
	StockFlg:  1,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestItemsGET(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	ItemsGET(c)
	assert.Equal(t, 200, w.Code)
}

func TestFindByItemID(t *testing.T) {
	item := FindByItemID(mockItem.ID)
	assert.NotNil(t, item)
	assert.IsType(t, mockItem, item)
}
