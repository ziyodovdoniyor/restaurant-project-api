package types

import "time"

type Food struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Category    string    `json:"category,omitempty"`
	Ingredients string    `json:"ingredients,omitempty"`
	Price       float32   `json:"price,omitempty"`
	Quantity    int       `json:"quantity,omitempty"`
	CookedAt    time.Time `json:"cooked_at,omitempty"`
}
type PreEnterFood struct {
	Name        string  `json:"name,omitempty"`
	Category    string  `json:"category,omitempty"`
	Ingredients string  `json:"ingredients,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
}

type UpdateFood struct {
	Name        string  `json:"name,omitempty"`
	Ingredients string  `json:"ingredients,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
}

const (
	FirstMeal  = "first_meal"
	SecondMeal = "second_meal"
	Dessert    = "dessert"
	Salad      = "salad"
	Beverage   = "beverage"
)

type Purchase struct {
	TableID      string    `json:"table_id,omitempty"`
	FirstMealID  string    `json:"first_meal_id,omitempty"`
	SecondMealID string    `json:"second_meal_id,omitempty"`
	DessertID    string    `json:"dessert_id,omitempty"`
	SaladID      string    `json:"salad_id,omitempty"`
	BeverageID   string    `json:"beverage_id,omitempty"`
	Total        int       `json:"total,omitempty"`
	PurchasedAt  time.Time `json:"purchased_at,omitempty"`
}

type Table struct {
	ID      string `json:"id,omitempty"`
	Number  int    `json:"number,omitempty"`
	IsTaken bool   `json:"is_taken,omitempty"`
}
