package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// UsersGET is getting List of users
func UsersGET(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	users := []model.User{}
	for result.Next() {
		user := model.User{}
		var id uint
		var mail string
		var password string
		var authkey string
		var point uint
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &mail, &password, &authkey, &point, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Mail = mail
		user.Authkey = authkey
		user.Point = point
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
		users = append(users, user)
	}
	fmt.Println(users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// FindByID is getting data by id
func FindByID(id uint) model.User {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for result.Next() {
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	return user
}

// FindByAuthkey is getting data by id
func FindByAuthkey(authkey string) model.User {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user WHERE authkey = ?", authkey)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for result.Next() {
		var id uint
		var createdAt, updatedAt time.Time
		err = result.Scan(&id, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	return user
}

// UserPOST is adding user
func UserPOST(c *gin.Context) {
	db := model.DBConnect()

	mail := c.PostForm("mail")
	password := c.PostForm("password")
	// TODO eventually we have to activate logic which will make authkey.
	// s := []string{mail, password}
	// joinedString := strings.Join(s, "")
	// b := []byte(joinedString)
	// authkey := sha512.Sum512(b)
	authkey := "hoge"
	point := 10000
	now := time.Now()

	_, err := db.Exec("INSERT INTO user (mail, password, authkey, point, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", mail, password, authkey, point, now, now)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. mail: %s", mail)
}

// UserPATCH is updating user
func UserPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	password := c.PostForm("password")
	now := time.Now()

	_, err := db.Exec("UPDATE user SET password = ?, updated_at=? WHERE id = ?", password, now, id)
	if err != nil {
		panic(err.Error())
	}

	user := FindByID(uint(id))

	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UserDELETE deletes user
func UserDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	// Check if record exists
	_, err := db.Query("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")
}

// UserAuth is getting user information by mail and password
func UserAuth(c *gin.Context) {
	db := model.DBConnect()

	mail := c.PostForm("mail")
	password := c.PostForm("password")
	result, err := db.Query("SELECT * FROM user WHERE mail = ? and password = ?", mail, password)
	if err != nil {
		panic(err.Error())
	}

	user := model.User{}
	for result.Next() {
		var id uint
		var mail string
		var password string
		var authkey string
		var point uint
		var createdAt, updatedAt time.Time

		err = result.Scan(&id, &mail, &password, &authkey, &point, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Mail = mail
		user.Authkey = authkey
		user.Point = point
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
