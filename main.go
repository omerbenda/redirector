package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omerbenda/redirector/db"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}

	defer logFile.Close()

	multiwriter := io.MultiWriter(logFile, os.Stdout)
	gin.DefaultWriter = multiwriter
	gin.DefaultErrorWriter = multiwriter
	log.SetOutput(gin.DefaultWriter)

	log.Println("Starting server...")

	log.Println("Reading DB")
	db.Read()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"shortenedCount": db.GetCount(),
			},
		)
	})

	r.GET(":id", func(c *gin.Context) {
		url, ok := db.GetValue(c.Param("id"))

		if ok {
			c.Redirect(http.StatusPermanentRedirect, url)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.POST("", func(c *gin.Context) {
		id := db.SetValue(c.Query("url"))

		c.String(http.StatusOK, id)
	})

	r.Run()
}
