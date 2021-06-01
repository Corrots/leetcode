package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": "internal server error",
				})
			}
		}()

		name := c.Param("name")
		panic("err")
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "OK",
			"data":    name,
		})
	})

	r.Run(":8080")
}
