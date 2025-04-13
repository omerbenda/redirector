package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omerbenda/redirector/db"
	"github.com/omerbenda/redirector/handlers"
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

	r.GET(":id", handlers.RedirectUrl)
	r.POST("", handlers.AddUrl)
	r.PUT("", handlers.UpdateUrl)
	r.DELETE("", handlers.DeleteUrl)

	r.Run()
}
