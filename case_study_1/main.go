package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}

var db *gorm.DB

func main() {
	initDB()
	// "Signin" and "Signup" are handler that we will implement
	r := gin.Default()
	r.POST("/signin", Signin)
	r.POST("/signup", Signup)

	// start the server on port 8080
	r.Run(":8080")
}

func initDB() {
	var err error
	db, _ = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/sign_in?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
}

func Signup(c *gin.Context) {
	// Parse and decode the request body into a new `Credentials` instance
	var cred Credentials
	c.BindJSON(&cred)

	var result Credentials

	db.Raw("select * from credentials where username=?", cred.Username).Scan(&result)

	if result.Username != "" {
		c.JSON(401, "User already registered")
	} else {
		db.Create(&cred)
		c.JSON(200, "User Registered")
	}
}

func Signin(c *gin.Context) {
	// Parse and decode the request body into a new `Credentials` instance
	var given_cred Credentials
	c.BindJSON(&given_cred)

	var result Credentials
	db.Raw("select * from credentials where username=?", given_cred.Username).Scan(&result)

	if result.Username == "" {
		c.JSON(401, "Username not registered")
	} else {
		if given_cred.Password != result.Password {
			c.JSON(401, "Wrong Credentials Inserted")
		} else {
			c.JSON(200, "Login Successful")
		}
	}

}
