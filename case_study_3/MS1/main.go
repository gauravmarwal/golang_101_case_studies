package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	r := gin.Default()

	r.GET("/ms1", random)

	r.Run(":8080")
}

func random(c *gin.Context) {

	Pass := &Message{Text: "Hello from MS1"}
	e, err := json.Marshal(Pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, string(e))
}
