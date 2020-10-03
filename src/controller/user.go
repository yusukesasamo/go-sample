package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/model"
)

// UsersGET gets List of users
func UsersGET(c *gin.Context) {
	db := model.DBConnect()
	// TODO we should specify limit and offest.
	result, err := db.Query("SELECT * FROM user ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	users := []model.User{}
	for result.Next() {
		user := model.User{}
		var id, point uint
		var mail, password, authkey string
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

// FindUserByID gets user by id
func FindUserByID(id uint) model.User {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for result.Next() {
		var id, point uint
		var mail, password, authkey string
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
	return user
}

// FindUserByAuthkey gets user by authkey
func FindUserByAuthkey(authkey string) model.User {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user WHERE authkey = ?", authkey)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for result.Next() {
		var id, point uint
		var mail, password, authkey string
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
	return user
}

// UserPOST adds user
func UserPOST(c *gin.Context) {
	db := model.DBConnect()

	mail := c.PostForm("mail")
	password := c.PostForm("password")
	// TODO eventually we have to activate this logic which will make authkey.
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

// UserPATCH updates user
func UserPATCH(c *gin.Context) {
	db := model.DBConnect()

	authkey := c.PostForm("authkey")
	user := FindUserByAuthkey(string(authkey))
	id := user.ID
	password := c.PostForm("password")
	now := time.Now()

	_, err := db.Exec("UPDATE user SET password = ?, updated_at=? WHERE id = ?", password, now, id)
	if err != nil {
		panic(err.Error())
	}

	updatedUser := FindUserByID(uint(id))

	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

// UserAuth gets user information by mail and password
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
		var id, point uint
		var mail, password, authkey string
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
