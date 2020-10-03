package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// ItemsGET gets List of item
func ItemsGET(c *gin.Context) {
	db := model.DBConnect()
	// TODO we should specify limit and offest.
	result, err := db.Query("SELECT * FROM item ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	items := []model.Item{}
	for result.Next() {
		item := model.Item{}
		var id uint
		var userID uint
		var name string
		var price uint
		var stockFlg uint
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &userID, &name, &price, &stockFlg, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		item.ID = id
		item.UserID = userID
		item.Name = name
		item.Price = price
		item.StockFlg = stockFlg
		item.CreatedAt = createdAt
		item.UpdatedAt = updatedAt
		items = append(items, item)
	}
	fmt.Println(items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// FindByItemID gets item data by id
func FindByItemID(id uint) model.Item {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM item WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	item := model.Item{}
	for result.Next() {
		var id uint
		var userID uint
		var name string
		var price uint
		var stockFlg uint
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &userID, &name, &price, &stockFlg, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		item.ID = id
		item.UserID = userID
		item.Name = name
		item.Price = price
		item.StockFlg = stockFlg
		item.CreatedAt = createdAt
		item.UpdatedAt = updatedAt
	}
	return item
}

// ItemPOST adds item
func ItemPOST(c *gin.Context) {
	db := model.DBConnect()
	authkey := c.PostForm("authkey")
	user := FindUserByAuthkey(string(authkey))

	userID := user.ID
	name := c.PostForm("name")
	price := c.PostForm("price")
	stockFlg := 1
	now := time.Now()

	_, err := db.Exec("INSERT INTO item (user_id, name, price, stock_flg, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", userID, name, price, stockFlg, now, now)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. name: %s", name)
}

// ItemPATCH updates item
func ItemPATCH(c *gin.Context) {
	db := model.DBConnect()

	authkey := c.PostForm("authkey")
	user := FindUserByAuthkey(string(authkey))

	userID := user.ID
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	price := c.PostForm("price")
	now := time.Now()

	_, err := db.Exec("UPDATE item SET name = ?, price = ? updated_at=? WHERE id = ? and user_id = ?", name, price, now, id, userID)
	if err != nil {
		panic(err.Error())
	}

	item := FindByItemID(uint(id))

	fmt.Println(item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

// ItemDELETE deletes item
func ItemDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))
	authkey := c.PostForm("authkey")
	user := FindUserByAuthkey(string(authkey))

	userID := user.ID

	// Check if record exists
	_, err := db.Query("DELETE FROM item WHERE id = ? and user_id = ?", id, userID)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")
}
