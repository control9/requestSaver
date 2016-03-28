package main

import "fmt"
import "os"
import "github.com/gin-gonic/gin"

const path = "./logs/"

func main() {
	initFolder()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/logger", func(c *gin.Context) {
		id := c.Query("id")
		write(path+"/access.log", fmt.Sprintf("request received: %s\n", id))
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

func initFolder() {
	os.MkdirAll(path, 0755)
}
func write(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
