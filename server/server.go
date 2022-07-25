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
	user types.Table
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



func NewRouter(repo Repository) *gin.Engine {
	r := gin.Default()
	h := Handler{repo: repo}
	r.GET("/menu", h.Menu)
	

	return r
}
