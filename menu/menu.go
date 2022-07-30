package menu

import (
	"restaurant/types"
	"time"

	"github.com/google/uuid"
)

func NewFood(category, name, ing string, price float32, quantity int) *types.Food {
	id := uuid.New()
	return &types.Food{
		ID:          id.String(),
		Name:        name,
		Category:    category,
		Ingredients: ing,
		Price:       price,
		Quantity:    quantity,
		CookedAt:    time.Now(),
	}
}

func NewPurchase(tableID, firstMealID, secondMealID, dessertID, saladID, beverageID string, total int) types.Purchase {
	return types.Purchase{
		TableID:      tableID,
		FirstMealID:  firstMealID,
		SecondMealID: secondMealID,
		BeverageID:   beverageID,
		SaladID:      saladID,
		DessertID:    dessertID,
		Total:        total,
		PurchasedAt:  time.Now(),
	}
}
