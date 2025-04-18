package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerbenda/redirector/db"
)

func RedirectUrl(c *gin.Context) {
	url, ok := db.GetValue(c.Param("id"))

	if ok {
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		c.Status(http.StatusNotFound)
	}
}

type AddUrlRequestBody struct {
	Url string
}

func AddUrl(c *gin.Context) {
	var body AddUrlRequestBody

	if err := c.BindJSON(&body); err != nil {
		log.Panic("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
	}

	id := db.SetValue(body.Url)

	c.String(http.StatusOK, id)
}

type UpdateUrlRequestBody struct {
	Id  string
	Url string
}

func UpdateUrl(c *gin.Context) {
	var body UpdateUrlRequestBody

	if err := c.BindJSON(&body); err != nil {
		log.Panic("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
	}

	exists := db.UpdateUrl(body.Id, body.Url)

	if exists {
		c.String(http.StatusOK, "Updated")
	} else {
		c.String(http.StatusNotFound, "Not Found")
	}
}

type DeleteUrlRequestBody struct {
	Id string
}

func DeleteUrl(c *gin.Context) {
	var body UpdateUrlRequestBody

	if err := c.BindJSON(&body); err != nil {
		log.Panic("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
	}

	exists := db.DeleteUrl(body.Id)

	if exists {
		c.String(http.StatusOK, "Deleted")
	} else {
		c.String(http.StatusNotFound, "Not Found")
	}
}
