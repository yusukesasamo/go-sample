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
		var createdAt, updatedAt time.Time
		var title string

		err = result.Scan(&id, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
		users = append(users, user)
	}
	fmt.Println(users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// FindByID is getting user by id
func FindByID(id uint) model.User {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for result.Next() {
		var createdAt, updatedAt time.Time
		var title string

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
	now := time.Now()

	_, err := db.Exec("INSERT INTO task (mail, password, created_at, updated_at) VALUES(?, ?, ?, ?)", mail, password, now, now)
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

	_, err := db.Exec("UPDATE task SET password = ?, updated_at=? WHERE id = ?", password, now, id)
	if err != nil {
		panic(err.Error())
	}

	task := FindByID(uint(id))

	fmt.Println(task)
	c.JSON(http.StatusOK, gin.H{"task": task})
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