package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type response struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Url     string            `json:"url"`
}

func home(c *gin.Context) {
	c.Data(200, "text/html", []byte("<p>Golang Echo!</p>"))
}

func get(c *gin.Context) {
	// Get query parameters
	queryParams := c.Request.URL.Query()

	// Create response map
	response := make(map[string]interface{})
	response["args"] = queryParams

	// Set headers in response
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		headers[k] = v[0]
	}
	response["headers"] = headers

	// Set other response fields
	response["url"] = c.Request.URL.String()
	response["origin"] = c.ClientIP()
	response["method"] = c.Request.Method

	// Return response as JSON
	c.JSON(http.StatusOK, response)
}

func headers(c *gin.Context) {
	// Create response map
	response := make(map[string]interface{})

	// Set headers in response
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		headers[k] = v[0]
	}
	response["headers"] = headers

	// Set other response fields
	response["url"] = c.Request.URL.String()
	response["origin"] = c.ClientIP()
	response["method"] = c.Request.Method

	// Return response as JSON
	c.JSON(http.StatusOK, response)
}

func post(c *gin.Context) {
	// Create response map
	response := make(map[string]interface{})
	var data map[string]interface{}

	if c.Request.Body == http.NoBody {
		data = make(map[string]interface{})
	} else {
		c.BindJSON(&data)
	}

	response["data"] = data

	// Set headers in response
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		headers[k] = v[0]
	}
	response["headers"] = headers

	// Set other response fields
	response["url"] = c.Request.URL.String()
	response["origin"] = c.ClientIP()
	response["method"] = c.Request.Method

	// Return response as JSON
	c.JSON(http.StatusOK, response)

}

func status(c *gin.Context) {
	response := make(map[string]interface{})

	status, _ := strconv.Atoi(c.Param("code"))

	response["status"] = status

	c.JSON(status, response)

}

func responseHeaders(c *gin.Context) {
	// Create response map
	response := make(map[string]interface{})

	// Set headers in response
	headers := make(map[string]string)
	for k, v := range c.Request.Response.Header {
		headers[k] = v[0]
	}
	response["responseHeaders"] = headers

	// Set other response fields
	response["url"] = c.Request.URL.String()
	response["origin"] = c.ClientIP()
	response["method"] = c.Request.Method

	// Return response as JSON
	c.JSON(http.StatusOK, response)

}
