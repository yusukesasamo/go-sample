package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// ItemsGET is getting List of data
func ItemsGET(c *gin.Context) {
	db := model.DBConnect()
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

// FindByItemID is getting data by id
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

// ItemPOST is adding user
func ItemPOST(c *gin.Context) {
	db := model.DBConnect()

	userID := c.PostForm("userID")
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

// ItemPATCH is updating user
func ItemPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.PostForm("userID")
	name := c.PostForm("name")
	now := time.Now()

	_, err := db.Exec("UPDATE item SET user_id = ?, name = ? updated_at=? WHERE id = ?", userID, name, now, id)
	if err != nil {
		panic(err.Error())
	}

	item := FindByItemID(uint(id))

	fmt.Println(item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

// ItemDELETE deletes user
func ItemDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.PostForm("userID")

	// Check if record exists
	_, err := db.Query("DELETE FROM user WHERE id = ? and user_id = ?", id, userID)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")
}
