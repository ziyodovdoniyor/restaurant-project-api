package server

import (
	"fmt"
	"net/http"

	"restaurant/menu"
	"restaurant/types"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Menu() ([]types.Food, error)
	AddFood(f types.Food) error
	GetFoodIDByName(foodName string, foods []types.Food) (string, string) 
	UpdateSecondMeal(id string, f types.Food) error 
	UpdateSaladMeal(id string, f types.Food) error
	UpdateDessertMeal(id string, f types.Food) error
	UpdateFirstMeal(id string, f types.Food) error
	UpdateBeverageMeal(id string, f types.Food) error
	GetFood(foods []types.Food, id string) (types.Food, error)
	DeleteFoodByName(foodID, cetegory string) error 
}

type Handler struct {
	repo  Repository
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

func (h *Handler) AddFood(c *gin.Context) {
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

func (h *Handler) UpdateFood(c *gin.Context)  {
	foodName, ok := c.GetQuery("name")
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid query %v", ok),
			},
		)
		return
	}

	var updateFood types.Food
	if err := c.BindJSON(&updateFood); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid json: %v", err),
			},
		)
		return
	}

	allFoods, err := h.repo.Menu()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf(" Menu(): %v", err),
			},
		)
		return
	}

	foodID, category := h.repo.GetFoodIDByName(foodName, allFoods)

	if category == types.SecondMeal {
		err = h.repo.UpdateSecondMeal(foodID, updateFood)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": fmt.Sprintf("UpdateFood(): %v", err),
				},
			)
			return
		}
	}
	if category == types.FirstMeal {
		err = h.repo.UpdateFirstMeal(foodID, updateFood)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": fmt.Sprintf("UpdateFood(): %v", err),
				},
			)
			return
		}
	}
	if category == types.Salad {
		err = h.repo.UpdateSaladMeal(foodID, updateFood)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": fmt.Sprintf("UpdateFood(): %v", err),
				},
			)
			return
		}
	}
	if category == types.Dessert {
		err = h.repo.UpdateDessertMeal(foodID, updateFood)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": fmt.Sprintf("UpdateFood(): %v", err),
				},
			)
			return
		}
	}
	if category == types.Beverage {
		err = h.repo.UpdateBeverageMeal(foodID, updateFood)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": fmt.Sprintf("UpdateFood(): %v", err),
				},
			)
			return
		}
	}

	

	c.Status(http.StatusOK)
}


func (h *Handler) GetFood(c *gin.Context)  {
	foodName, ok := c.GetQuery("name")
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid query %v", ok),
			},
		)
		return
	}

	allFoods, err := h.repo.Menu()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf(" Menu(): %v", err),
			},
		)
		return
	}

	foodID, _ := h.repo.GetFoodIDByName(foodName, allFoods)

	WantedFood, err := h.repo.GetFood(allFoods, foodID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid query %v", ok),
			},
		)
		return
	}

	c.JSON(http.StatusOK, WantedFood)

}

func (h *Handler) DeleteFood(c *gin.Context)  {
	foodName, ok := c.GetQuery("name")
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid query %v", ok),
			},
		)
		return
	}

	allFoods, err := h.repo.Menu()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf(" DeleteFood(): %v", err),
			},
		)
		return
	}

	foodID, category := h.repo.GetFoodIDByName(foodName, allFoods)
	 
	err = h.repo.DeleteFoodByName(foodID, category)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf(" DeleteFood(): %v", err),
			},
		)
		return
	}
}




func NewRouter(repo Repository) *gin.Engine {
	r := gin.Default()
	h := Handler{repo: repo}
	r.GET("/menu", h.Menu)
	r.GET("/menu/first-meal")
	r.GET("/menu/second-meal")
	r.GET("/menu/salad")
	r.GET("/menu/dessert")
	r.GET("/menu/drinks")

	r.GET("/table/")
	r.POST("/table/buy/")

	r.GET("/table/buy/budget/")

	r.POST("/add/food", h.AddFood)
	r.GET("/food/", h.GetFood)
	r.PUT("/update/food/", h.UpdateFood)
	r.DELETE("/delete/food/", h.DeleteFood)

	return r
}
