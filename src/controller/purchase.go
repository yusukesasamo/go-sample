package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// PurchasePOST purchase item
func PurchasePOST(c *gin.Context) {
	db := model.DBConnect()
	authkey := c.PostForm("authkey")
	itemID, _ := strconv.Atoi(c.PostForm("itemID"))
	user := FindUserByAuthkey(string(authkey))
	item := FindByItemID(uint(itemID))
	now := time.Now()

	fmt.Println(user.Point)
	fmt.Println(item.Price)

	tx, beginError := db.Begin()
	if beginError != nil {
		panic(beginError.Error())
	}
	_, updateItemErr := db.Exec("UPDATE item SET stock_flg = ?, updated_at=? WHERE id = ?", 0, now, item.ID)
	if updateItemErr != nil {
		panic(updateItemErr.Error())
	}

	remainPoint := user.Point - item.Price
	//TODO We will implement logc if remainPoint less than 0
	_, updateUserErr := db.Exec("UPDATE user SET point = ?, updated_at=? WHERE id = ?", remainPoint, now, user.ID)
	if updateUserErr != nil {
		panic(updateUserErr.Error())
	}
	tx.Commit()
}
