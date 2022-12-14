package api

import (
	"github.com/DadaxonOlimov/services"
	"github.com/gin-gonic/gin"
)

func NewServer(sm services.ServiceManager) *gin.Engine {

	r := gin.Default()

	c := NewController(sm)

	todos := r.Group("/api/v1")
	{
		todos.GET("/todos", c.GetAll)
		todos.POST("/todos", c.Create)
		todos.PUT("/todos/:id", c.Update)
		todos.DELETE("/todos/:id", c.Delete)
		todos.GET("/todos/:id", c.GetByID)
	}

	return r
}