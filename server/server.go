package server

import (
	"database/sql"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"restaurant/postgres"
	"strconv"

	_ "restaurant/docs"
	"restaurant/menu"
	"restaurant/types"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//API and REPO

type Repository interface {
	// sunbula
	Menu() ([]types.Food, error)
	AddFood(f types.Food) error
	GetFoodIDByName(foodName string, foods []types.Food) (string, string)
	UpdateSecondMeal(id string, f types.UpdateFood) error
	UpdateSaladMeal(id string, f types.UpdateFood) error
	UpdateDessertMeal(id string, f types.UpdateFood) error
	UpdateFirstMeal(id string, f types.UpdateFood) error
	UpdateBeverageMeal(id string, f types.UpdateFood) error
	GetFood(foods []types.Food, id string) (types.Food, error)
	DeleteFoodByName(foodID, cetegory string) error
	// sunbula

	GetTables() ([]types.Table, error)
	TakeTable(num int) (types.Table, bool, error)
	Buy(purchase postgres.PurchaseR) (int, error)

	// ibrohimjon
	First() ([]types.Food, error)
	Second() ([]types.Food, error)
	Salad() ([]types.Food, error)
	Dessert() ([]types.Food, error)
	Drink() ([]types.Food, error)
	// doniyor
	Sets(cash float64) ([][]types.Food, error)
}

type Handler struct {
	repo  Repository
	table types.Table
}

// NewRouter
// @title           Swagger Restaurant API
// @version         1.0
// @description     This is a restaurant project.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

func NewRouter(repo Repository) *gin.Engine {
	r := gin.Default()
	h := Handler{repo: repo}
	r.GET("/menu/first-meal", h.First)   // done
	r.GET("/menu/second-meal", h.Second) // done
	r.GET("/menu/salad", h.Salad)        // done
	r.GET("/menu/dessert", h.Dessert)    // done
	r.GET("/menu/drinks", h.Drink)       // done

	r.GET("/tables", h.GetTables) // done
	r.POST("/table", h.TakeTable) // done
	r.POST("/table/buy", h.Buy)   // done

	// sunbula
	r.GET("/menu", h.Menu)                  // done
	r.POST("/add/food", h.AddFood)          //
	r.GET("/food/", h.GetFood)              //
	r.PUT("/update/food/", h.UpdateFood)    //
	r.DELETE("/delete/food/", h.DeleteFood) //
	// sunbula

	//doniyor
	r.GET("/set", h.Sets) // problem with query

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// Buy
// @Summary      Buy a product
// @Description  you can buy any food
// @Tags         tables
// @Accept       json
// @Produce      json
// @Param        request   body postgres.PurchaseR  true  "order info"
// @Success      200 {object} types.Purchase
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /table/buy [post]
func (h *Handler) Buy(c *gin.Context) {
	var request postgres.PurchaseR
	if er := c.ShouldBindJSON(&request); er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.Error(),
		})
		return
	}
	sum, er := h.repo.Buy(request)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.Error(),
		})
		return
	}
	pch := menu.NewPurchase(request.TableID, request.FirstMealID, request.SecondMealID, request.DessertID, request.SaladID, request.BeverageID, sum)
	c.JSON(500, pch)
}

// TakeTable
// @Summary      Order a table
// @Description  you can choose one of the free tables
// @Tags         tables
// @Accept       json
// @Produce      json
// @Param        num   query      int  true  "table number"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /table [post]
func (h *Handler) TakeTable(c *gin.Context) {
	table := h.table
	var er error
	var b bool
	num := c.Query("num")
	n, er := strconv.Atoi(num)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "please enter a number not string",
		})
		return
	}
	table, b, er = h.repo.TakeTable(n)
	if er == sql.ErrNoRows && b {
		c.JSON(http.StatusInternalServerError, gin.H{
			":(": fmt.Sprintf("we don't have table with № %d", n),
		})
		return
	} else if er != nil && !b {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.Error(),
		})
		return
	} else if !b && er == nil {
		c.JSON(http.StatusOK, gin.H{
			":(": fmt.Sprintf("table with № %d is taken", table.Number),
		})
		return
	} else if er != nil && b {
		c.JSON(http.StatusOK, gin.H{
			":(": fmt.Sprintf("could not give you your table № %d", table.Number),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		":)": fmt.Sprintf("your table's id is %s", table.ID),
	})
	return

}

// GetTables
// @Summary      GetTables
// @Description  shows all the free tables
// @Tags         tables
// @Accept       json
// @Produce      json
// @Success      200  {array}  []int
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /tables [get]
func (h *Handler) GetTables(c *gin.Context) {
	tables, er := h.repo.GetTables()
	if er != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tables)
}

func (h *Handler) First(c *gin.Context) {
	firstFoods, err := h.repo.First()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "data couldn't be shown",
			})
	}
	c.JSON(http.StatusOK, firstFoods)
}

func (h *Handler) Second(c *gin.Context) {
	secondFoods, err := h.repo.Second()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "data couldn't be shown",
			})
	}
	c.JSON(http.StatusOK, secondFoods)
}

func (h *Handler) Salad(c *gin.Context) {
	salad, err := h.repo.Salad()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "data couldn't be shown",
			})
	}
	c.JSON(http.StatusOK, salad)
}

func (h *Handler) Dessert(c *gin.Context) {
	dessert, err := h.repo.Dessert()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "data couldn't be shown",
			})
	}
	c.JSON(http.StatusOK, dessert)
}

func (h *Handler) Drink(c *gin.Context) {
	drink, err := h.repo.Drink()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "data couldn't be shown",
			})
	}
	c.JSON(http.StatusOK, drink)
}

// Menu
// @Summary      Menu
// @Description  shows all items in the menu
// @Tags         sunbula
// @Produce      json
// @Success      200  {array}  []types.Food
// @Failure      500
// @Router       /menu [GET]
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

// AddFood
// @Summary      Add food
// @Description  Add food to the menu, food name must be unique
// @Tags         sunbula
// @Accept       json
// @Param        request body types.PreEnterFood  true  "Food info"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /add/food [POST]
func (h *Handler) AddFood(c *gin.Context) {
	var food types.PreEnterFood
	if err := c.BindJSON(&food); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid json: %v", err),
			},
		)
		return
	}
	newFood := menu.NewFood(food.Category, food.Name, food.Ingredients, food.Price, food.Quantity)
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

// UpdateFood
// @Summary      Update food
// @Description  Update food in the menu
// @Tags         sunbula
// @Accept       json
// @Param        name    query     string  true  "update food by name"
// @Param        request body types.UpdateFood true "Food info"
// @Success      200
// @Failure      400
// @Failure 	 404
// @Failure      500
// @Router       /update/food/ [PUT]
func (h *Handler) UpdateFood(c *gin.Context) {
	foodName, ok := c.GetQuery("name")
	if !ok {
		message := fmt.Sprintf("name not spicified: %t", ok)
		c.String(http.StatusBadRequest, message)
		return
	}

	var updateFood types.UpdateFood
	if err := c.BindJSON(&updateFood); err != nil {
		message := fmt.Sprintf("invalid json: %v", err)
		c.String(http.StatusBadRequest, message)
		return
	}

	allFoods, err := h.repo.Menu()
	if err != nil {
		message := fmt.Sprintf("couldn't fetch data from menu: %v", err)
		c.String(http.StatusInternalServerError, message)
		return
	}

	foodID, category := h.repo.GetFoodIDByName(foodName, allFoods)

	if category == types.SecondMeal {
		err = h.repo.UpdateSecondMeal(foodID, updateFood)
		if err != nil {
			if err == sql.ErrNoRows {
				message := fmt.Sprintf("item not found from menu: %v", err)
				c.String(http.StatusNotFound, message)
				return
			} else {
				message := fmt.Sprintf("couldn't update item %v", err)
				c.String(http.StatusInternalServerError, message)
				return
			}
		}
	}
	if category == types.FirstMeal {
		err = h.repo.UpdateFirstMeal(foodID, updateFood)
		if err != nil {
			if err == sql.ErrNoRows {
				message := fmt.Sprintf("item not found from menu: %v", err)
				c.String(http.StatusNotFound, message)
				return
			} else {
				message := fmt.Sprintf("couldn't update item %v", err)
				c.String(http.StatusInternalServerError, message)
				return
			}
		}
	}
	if category == types.Salad {
		err = h.repo.UpdateSaladMeal(foodID, updateFood)
		if err != nil {
			if err == sql.ErrNoRows {
				message := fmt.Sprintf("item not found from menu: %v", err)
				c.String(http.StatusNotFound, message)
				return
			} else {
				message := fmt.Sprintf("couldn't update item %v", err)
				c.String(http.StatusInternalServerError, message)
				return
			}
		}
	}
	if category == types.Dessert {
		err = h.repo.UpdateDessertMeal(foodID, updateFood)
		if err != nil {
			if err == sql.ErrNoRows {
				message := fmt.Sprintf("item not found from menu: %v", err)
				c.String(http.StatusNotFound, message)
				return
			} else {
				message := fmt.Sprintf("couldn't update item %v", err)
				c.String(http.StatusInternalServerError, message)
				return
			}
		}
	}
	if category == types.Beverage {
		err = h.repo.UpdateBeverageMeal(foodID, updateFood)
		if err != nil {
			if err == sql.ErrNoRows {
				message := fmt.Sprintf("item not found from menu: %v", err)
				c.String(http.StatusNotFound, message)
				return
			} else {
				message := fmt.Sprintf("couldn't update item %v", err)
				c.String(http.StatusInternalServerError, message)
				return
			}
		}
	}

	message := "succesfuly updated"
	c.String(http.StatusOK, message)
}

// GetFood
// @Summary      Get food
// @Description  it gets all information about the asked food
// @Tags         sunbula
// @Accept       json
// @Param        name    query     string  true  "search food by name"
// @Success      200  {object} types.Food
// @Failure      400
// @Failure 	 404
// @Failure      500
// @Router       /food/ [GET]
func (h *Handler) GetFood(c *gin.Context) {
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
			http.StatusNotFound,
			gin.H{
				"error": fmt.Sprintf("item not found %v", err),
			},
		)
		return
	}

	c.JSON(http.StatusOK, WantedFood)

}

// DeleteFood
// @Summary      Delete food
// @Description  deletes food by its name
// @Tags         sunbula
// @Accept       json
// @Param        name    query     string  true  "delete food by name"
// @Success      200
// @Failure      400
// @Failure 	 404
// @Failure      500
// @Router       /delete/food/ [DELETE]
func (h *Handler) DeleteFood(c *gin.Context) {
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
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"error": fmt.Sprintf(" DeleteFood(): %v", err),
				},
			)
			return
		} else {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": fmt.Sprintf(" DeleteFood(): %v", err),
				},
			)
			return
		}

	}

	c.String(http.StatusOK, "successfully deleted")
}

//Sets metodi userga set yaratib beradi
func (h Handler) Sets(c *gin.Context) {
	cashstr, ok := c.GetQuery("cash")
	cash, err := strconv.ParseFloat(cashstr, 64)
	if err != nil {
		panic(err)
	}
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("invalid query %v", ok),
			},
		)
		return
	}

	sets, err := h.repo.Sets(cash)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("Sets(): %v", err),
			},
		)
		return
	}

	c.JSON(http.StatusOK, sets)
}
