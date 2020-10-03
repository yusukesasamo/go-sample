package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// UserPurchaseHistoriesGET gets List of user_purchase_history
func UserPurchaseHistoriesGET(c *gin.Context) {
	db := model.DBConnect()
	// TODO we should specify limit and offest.
	result, err := db.Query("SELECT * FROM user_purchase_history ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	userPurchaseHistories := []model.UserPurchaseHistory{}
	for result.Next() {
		userPurchaseHistory := model.UserPurchaseHistory{}
		var id uint
		var userID uint
		var itemID uint
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &userID, &itemID, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		userPurchaseHistory.ID = id
		userPurchaseHistory.UserID = userID
		userPurchaseHistory.ItemID = itemID
		userPurchaseHistory.CreatedAt = createdAt
		userPurchaseHistory.UpdatedAt = updatedAt
		userPurchaseHistories = append(userPurchaseHistories, userPurchaseHistory)
	}
	fmt.Println(userPurchaseHistories)
	c.JSON(http.StatusOK, gin.H{"userPurchaseHistories": userPurchaseHistories})
}

// FindByUserPurchaseHistoryID is getting data by id
func FindByUserPurchaseHistoryID(id uint) model.UserPurchaseHistory {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user_purchase_history WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	userPurchaseHistory := model.UserPurchaseHistory{}
	for result.Next() {
		var id uint
		var userID uint
		var itemID uint
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &userID, &itemID, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		userPurchaseHistory.ID = id
		userPurchaseHistory.ItemID = itemID
		userPurchaseHistory.CreatedAt = createdAt
		userPurchaseHistory.UpdatedAt = updatedAt
	}
	return userPurchaseHistory
}

// UserPurchaseHistoryPOST is adding data
func UserPurchaseHistoryPOST(c *gin.Context) {
	db := model.DBConnect()
	authkey := c.PostForm("authkey")
	user := FindUserByAuthkey(string(authkey))
	userID := user.ID
	itemID := c.PostForm("itemID")
	now := time.Now()

	_, err := db.Exec("INSERT INTO user (user_id, item_id, created_at, updated_at) VALUES(?, ?, ?, ?)", userID, itemID, now, now)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. userID: %s", itemID)
}
