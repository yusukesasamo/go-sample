package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PurchasePOST purchase item
func PurchasePOST(c *gin.Context) {
	// db := model.DBConnect()
	authkey := c.PostForm("authkey")
	itemID, _ := strconv.Atoi(c.PostForm("itemID"))
	user := FindUserByAuthkey(string(authkey))
	item := FindByItemID(uint(itemID))
	fmt.Println(user.Point)
	fmt.Println(item.Price)
	// userID := user.ID
	// userPoint := user.Point
	// itemID := c.PostForm("itemID")

	// /// updateitem db
	// now := time.Now()

	// _, err := db.Exec("INSERT INTO item (user_id, name, price, stock_flg, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", userID, name, price, stockFlg, now, now)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Printf("post sent. name: %s", name)
}
