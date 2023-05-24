package api

import (
	"strconv"
	"web-service/pkg/service"

	"github.com/gin-gonic/gin"
)

func RouteMovies(router *gin.Engine) {
	router.GET("/movies", movies)
}

func movies(c *gin.Context) {
	page, error := strconv.Atoi(c.Query("page"))
	if error != nil {
		page = 0
	}
	size, error := strconv.Atoi(c.Query("size"))
	if error != nil {
		size = 0
	}
	movies, err := service.FindMovies(page, size)
	if err == nil {
		c.JSON(200, movies)
	}
}
