package server

import (
	"fmt"
	"net/http"

	"restaurant/types"
	"restaurant/menu"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Menu() ([]types.Food, error)
	AddFood(f types.Food) error 

}

type Handler struct {
	repo Repository
	user types.Client
}


func (h *Handler) Menu(c *gin.Context) {
	data, err := h.repo.Menu()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "data couldn't be shown",
			},
		)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) AddFood(c *gin.Context)  {
	if !h.user.IsAdmin {
		c.AbortWithStatusJSON(
			http.StatusMethodNotAllowed,
			gin.H{
				"error": "this method is only allowed to admins",
			},
		)
		return
	}

	var food types.Food
	if err := c.BindJSON(&food); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid json: %v", err),
			},
		)
		return
	}
	newFood := menu.NewFood(food.Category, food.Name, food.Ingredients, food.Price)
	err := h.repo.AddFood(*newFood)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("error in writing to the database: %v", err),
			},
		)
		return
	}
	c.Status(http.StatusCreated)
}

func NewRouter(repo Repository) *gin.Engine {
	r := gin.Default()
	h := Handler{repo: repo}
	r.GET("/menu", h.Menu)
	r.POST("/add/food", h.AddFood)
	

	return r
}
