package server

import (
	"net/http"

	"restaurant/types"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Menu() ([]types.Food, error)
	AddFood(f types.Food) error 

}

type Handler struct {
	repo Repository
	table types.Table
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



//Name        string    `json:"name,omitempty"`     
// Category    string    `json:"category,omitempty"`
// Ingredients string    `json:"ingredients,omitempty"`
// Price       int       `json:"price,omitempty"`

func NewRouter(repo Repository) *gin.Engine {
	r := gin.Default()
	h := Handler{repo: repo}
	r.GET("/menu", h.Menu)
	r.GET("/menu/first-meal", )
	r.GET("/menu/second-meal", )
	r.GET("/menu/salad", )
	r.GET("/menu/dessert", )
	r.GET("/menu/drinks", )
	
	r.GET("/table/")
	r.GET("/table/buy/")

	r.GET("/table/buy/budget/")

	r.POST("/add/food")
	r.PUT("/add/food/")
	r.DELETE("/add/food/")



	return r
}
