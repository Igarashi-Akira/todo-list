package controller

import (
	"fmt"
  "github.com/gin-gonic/gin"
	"github.com/Igarashi-Akira/todo-list/service"
)

type Controller struct {}

func (pc Controller) Index(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) Create(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(c)

	if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
	} else {
			c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
	} else {
			c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
	} else {
			c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service

	if err := s.DeleteByID(id); err != nil {
			c.AbortWithStatus(403)
			fmt.Println(err)
	} else {
			c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}