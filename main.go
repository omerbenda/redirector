package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omerbenda/redirector/db"
)

type CreateUrlRequestBody struct {
	Url string
}

type UpdateUrlRequestBody struct {
	Id  string
	Url string
}

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

	log.Println("Loading templates")
	r.LoadHTMLGlob("templates/*")

	log.Println("Loading static files")
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
		var body CreateUrlRequestBody

		if err := c.BindJSON(&body); err != nil {
			log.Panic("Error binding JSON: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
		}

		id := db.SetValue(body.Url)

		c.String(http.StatusOK, id)
	})

	r.PUT("", func(c *gin.Context) {
		var body UpdateUrlRequestBody

		if err := c.BindJSON(&body); err != nil {
			log.Panic("Error binding JSON: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
		}

		exists := db.UpdateValue(body.Id, body.Url)

		if exists {
			c.String(http.StatusOK, "Updated")
		} else {
			c.String(http.StatusNotFound, "Not Found")
		}
	})

	r.Run()
}
