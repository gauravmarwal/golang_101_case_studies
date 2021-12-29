package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	r := gin.Default()

	r.GET("/ms2", random)

	r.Run(":8000")
}

func random(c *gin.Context) {

	fmt.Print("MS1 says: ")

	status, err := http.Get("http://localhost:8080/ms1")

	if err != nil {
		fmt.Println("This Err: ", err)
	}

	// fmt.Println("This is status: ", &status, "\nerr: ", err)
	defer status.Body.Close()
	body, err := io.ReadAll(status.Body)
	if err != nil {
		fmt.Println("This Err: ", err)
	}

	// var temp Message
	// json.Unmarshal([]byte(body), &temp)

	c.JSON(200, string(body))
}
