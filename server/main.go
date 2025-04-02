package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerbenda/redirector/db"
)

func main() {
	db.Read()

	r := gin.Default()

	r.GET(":hash", func(c *gin.Context) {
		url, ok := db.GetValue(c.Param("hash"))

		if ok {
			c.Redirect(http.StatusPermanentRedirect, url)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.POST("", func(c *gin.Context) {
		hash := db.SetValue(c.Query("url"))

		c.String(http.StatusOK, hash)
	})

	r.Run()
}
