package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// PurchasePOST purchase item
func PurchasePOST(c *gin.Context) {
	authkey := c.PostForm("authkey")
	itemID, _ := strconv.Atoi(c.PostForm("itemID"))
	user := FindUserByAuthkey(string(authkey))
	now := time.Now()

	db := model.DBConnect()
	tx, beginError := db.Begin()
	if beginError != nil {
		panic(beginError.Error())
	}
	result, err := db.Query("SELECT * FROM item WHERE id = ? FOR UPDATE", itemID)
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

	if item.StockFlg != 1 {
		c.JSON(http.StatusBadRequest, "Alreadypurchased")
	} else {
		_, updateItemErr := db.Exec("UPDATE item SET stock_flg = ?, updated_at=? WHERE id = ?", 0, now, item.ID)
		if updateItemErr != nil {
			tx.Rollback()
			panic(updateItemErr.Error())
		}

		remainPoint := user.Point - item.Price
		//TODO We will implement logc if remainPoint less than 0
		_, updateUserErr := db.Exec("UPDATE user SET point = ?, updated_at=? WHERE id = ?", remainPoint, now, user.ID)
		if updateUserErr != nil {
			tx.Rollback()
			panic(updateUserErr.Error())
		}
		tx.Commit()
	}
}
