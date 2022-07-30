package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"restaurant/types"
	"time"
)

// PSQL

type PostgresRepository struct {
	db *sql.DB
}

type PurchaseR struct {
	TableID      string `json:"table_id,omitempty"`
	FirstMealID  string `json:"first_meal_id,omitempty"`
	SecondMealID string `json:"second_meal_id,omitempty"`
	DessertID    string `json:"dessert_id,omitempty"`
	SaladID      string `json:"salad_id,omitempty"`
	BeverageID   string `json:"beverage_id,omitempty"`
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p *PostgresRepository) Buy(purchase PurchaseR) (int, error) {
	var total, price int
	var row *sql.Row

	if purchase.FirstMealID != "" {
		row := p.db.QueryRow("select price from first_meal where id = $1", purchase.FirstMealID)
		if er := row.Scan(&price); er != nil {
			return 0, er
		}
		total += price
	}

	if purchase.SecondMealID != "" {
		row := p.db.QueryRow("select price from second_meal where id = $1", purchase.SecondMealID)
		if er := row.Scan(&price); er != nil {
			return 0, er
		}
		total += price
	}

	if purchase.BeverageID != "" {
		row = p.db.QueryRow("select price from beverage where id = $1", purchase.BeverageID)
		if er := row.Scan(&price); er != nil {
			return 0, er
		}
		total += price
	}

	if purchase.SaladID != "" {
		row = p.db.QueryRow("select price from salad where id = $1", purchase.SaladID)
		if er := row.Scan(&price); er != nil {
			return 0, er
		}
		total += price
	}

	if purchase.DessertID != "" {
		row = p.db.QueryRow("select price from dessert where id = $1", purchase.DessertID)
		if er := row.Scan(&price); er != nil {
			return 0, er
		}
		total += price
	}
	return total, nil
}

func (p *PostgresRepository) TakeTable(num int) (types.Table, bool, error) {
	var table types.Table
	row := p.db.QueryRow("select * from tables where table_number = $1", num)
	if er := row.Scan(&table.ID, &table.Number, &table.IsTaken); er == sql.ErrNoRows {
		return types.Table{}, true, er
	} else if er != nil {
		return types.Table{}, false, er
	} else if table.IsTaken {
		return table, false, nil
	}
	if _, er := p.db.Exec("update tables set is_taken=$1 where table_number=$2", true, num); er != nil {
		return types.Table{}, true, er
	}
	return table, true, nil
}

func (p *PostgresRepository) GetTables() ([]types.Table, error) {
	rows, er := p.db.Query("select * from tables where is_taken = $1", false)
	if er != nil {
		return nil, er
	}
	tables := []types.Table{}

	for rows.Next() {
		table := types.Table{}
		if er = rows.Scan(&table.ID, &table.Number, &table.IsTaken); er != nil {
			return nil, er
		}
		tables = append(tables, table)
	}
	return tables, nil
}

//************************************************

func (ps *PostgresRepository) First() ([]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var foods []types.Food

	rows, err := tx.Query(`SELECT id, name, ingredients, price, cooked_at FROM first_meal`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foods = append(foods, f)
	}
	tx.Commit()
	return foods, err
}

func (ps *PostgresRepository) Second() ([]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var foods []types.Food

	rows, err := tx.Query(`SELECT id, name, ingredients, price, cooked_at FROM second_meal`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foods = append(foods, f)
	}
	tx.Commit()
	return foods, err
}

func (ps *PostgresRepository) Salad() ([]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var foods []types.Food

	rows, err := tx.Query(`SELECT id, name, ingredients, price, cooked_at FROM salad`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foods = append(foods, f)
	}
	tx.Commit()
	return foods, err
}

func (ps *PostgresRepository) Dessert() ([]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var foods []types.Food

	rows, err := tx.Query(`SELECT id, name, ingredients, price, cooked_at FROM dessert`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foods = append(foods, f)
	}
	tx.Commit()
	return foods, err
}

func (ps *PostgresRepository) Drink() ([]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var foods []types.Food

	rows, err := tx.Query(`SELECT id, name, ingredients, price, cooked_at FROM beverage`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foods = append(foods, f)
	}
	tx.Commit()
	return foods, err
}

// sunbula **************************************************************************************************************
func (ps *PostgresRepository) Menu() ([]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var foods []types.Food

	rows, err := tx.Query(`SELECT * FROM first_meal`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		f.Category = types.FirstMeal
		foods = append(foods, f)
	}

	rows, err = tx.Query(`SELECT * FROM second_meal`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		f.Category = types.SecondMeal
		foods = append(foods, f)
	}

	rows, err = tx.Query(`SELECT * FROM dessert`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		f.Category = types.Dessert
		foods = append(foods, f)
	}

	rows, err = tx.Query(`SELECT * FROM salad`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		f.Category = types.Salad
		foods = append(foods, f)
	}

	rows, err = tx.Query(`SELECT * FROM beverage`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		f := types.Food{}
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		f.Category = types.Beverage
		foods = append(foods, f)
	}

	tx.Commit()
	return foods, err
}

func (ps *PostgresRepository) AddFood(f types.Food) error {
	if f.Category == types.FirstMeal {
		_, err := ps.db.Exec(`
			INSERT INTO first_meal (id, name, ingredients, price, quantity, cooked_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, f.ID, f.Name, f.Ingredients, f.Price, f.Quantity, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.SecondMeal {
		_, err := ps.db.Exec(`
			INSERT INTO second_meal (id, name, ingredients, price, qauntity, cooked_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, f.ID, f.Name, f.Ingredients, f.Price, f.Quantity, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.Salad {
		_, err := ps.db.Exec(`
		INSERT INTO salad (id, name, ingredients, price, quantity, cooked_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, f.ID, f.Name, f.Ingredients, f.Price, f.Quantity, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.Dessert {
		_, err := ps.db.Exec(`
		INSERT INTO dessert (id, name, ingredients, price, quantity, cooked_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, f.ID, f.Name, f.Ingredients, f.Price, f.Quantity, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.Beverage {
		_, err := ps.db.Exec(`
		INSERT INTO beverage(id, name, ingredients, price, qauntity, cooked_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, f.ID, f.Name, f.Ingredients, f.Price, f.Quantity, f.CookedAt)

		if err != nil {
			return err
		}
	}

	return nil
}

func (ps *PostgresRepository) GetFoodIDByName(foodName string, foods []types.Food) (string, string) {
	id := ""
	category := ""
	for _, v := range foods {
		if v.Name == foodName {
			id = v.ID
			category = v.Category
			break
		}
	}

	return id, category
}

func (ps *PostgresRepository) UpdateSecondMeal(id string, f types.Food) error {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if f.Name != "" {
		_, err := tx.Exec(`
			UPDATE second_meal SET name = $1 WHERE id = $2
		`, f.Name, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Ingredients != "" {
		_, err := tx.Exec(`
			UPDATE second_meal SET ingredients = $1 WHERE id = $2
		`, f.Ingredients, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Price != 0 {
		_, err := tx.Exec(`
			UPDATE second_meal SET price = $1 WHERE id = $2
		`, f.Price, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Quantity != 0 {
		_, err := tx.Exec(`
			UPDATE second_meal SET qauntity = $1 WHERE id = $2
		`, f.Quantity, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`
		UPDATE second_meal SET cooked_at = $1 WHERE id = $2
	`, time.Now(), id)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ps *PostgresRepository) UpdateFirstMeal(id string, f types.Food) error {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if f.Name != "" {
		_, err := tx.Exec(`
			UPDATE first_meal SET name = $1 WHERE id = $2
		`, f.Name, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Ingredients != "" {
		_, err := tx.Exec(`
			UPDATE first_meal SET ingredients = $1 WHERE id = $2
		`, f.Ingredients, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Price != 0 {
		_, err := tx.Exec(`
			UPDATE first_meal SET price = $1 WHERE id = $2
		`, f.Price, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Quantity != 0 {
		_, err := tx.Exec(`
			UPDATE first_meal SET qauntity = $1 WHERE id = $2
		`, f.Quantity, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`
		UPDATE first_meal SET cooked_at = $1 WHERE id = $2
	`, time.Now(), id)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ps *PostgresRepository) UpdateSaladMeal(id string, f types.Food) error {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if f.Name != "" {
		_, err := tx.Exec(`
			UPDATE salad SET name = $1 WHERE id = $2
		`, f.Name, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Ingredients != "" {
		_, err := tx.Exec(`
			UPDATE salad SET ingredients = $1 WHERE id = $2
		`, f.Ingredients, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Price != 0 {
		_, err := tx.Exec(`
			UPDATE salad SET price = $1 WHERE id = $2
		`, f.Price, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Quantity != 0 {
		_, err := tx.Exec(`
			UPDATE salad SET qauntity = $1 WHERE id = $2
		`, f.Quantity, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`
		UPDATE salad SET cooked_at = $1 WHERE id = $2
	`, time.Now(), id)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ps *PostgresRepository) UpdateDessertMeal(id string, f types.Food) error {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if f.Name != "" {
		_, err := tx.Exec(`
			UPDATE dessert SET name = $1 WHERE id = $2
		`, f.Name, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Ingredients != "" {
		_, err := tx.Exec(`
			UPDATE dessert SET ingredients = $1 WHERE id = $2
		`, f.Ingredients, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Price != 0 {
		_, err := tx.Exec(`
			UPDATE dessert SET price = $1 WHERE id = $2
		`, f.Price, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Quantity != 0 {
		_, err := tx.Exec(`
			UPDATE dessert SET qauntity = $1 WHERE id = $2
		`, f.Quantity, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`
		UPDATE dessert SET cooked_at = $1 WHERE id = $2
	`, time.Now(), id)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ps *PostgresRepository) UpdateBeverageMeal(id string, f types.Food) error {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if f.Name != "" {
		_, err := tx.Exec(`
			UPDATE beverage SET name = $1 WHERE id = $2
		`, f.Name, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Ingredients != "" {
		_, err := tx.Exec(`
			UPDATE beverage SET ingredients = $1 WHERE id = $2
		`, f.Ingredients, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Price != 0 {
		_, err := tx.Exec(`
			UPDATE beverage SET price = $1 WHERE id = $2
		`, f.Price, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if f.Quantity != 0 {
		_, err := tx.Exec(`
			UPDATE beverage SET qauntity = $1 WHERE id = $2
		`, f.Quantity, id)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`
		UPDATE beverage SET cooked_at = $1 WHERE id = $2
	`, time.Now(), id)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ps *PostgresRepository) GetFood(foods []types.Food, id string) (types.Food, error) {
	var f types.Food
	exist := false
	for _, v := range foods {
		if id == v.ID {
			f = v
			exist = true
		}
	}

	if !exist {
		return types.Food{}, fmt.Errorf("product doesn't exist")
	}

	return f, nil
}

func (ps *PostgresRepository) DeleteFoodByName(foodID, cetegory string) error {
	if cetegory == types.FirstMeal {
		_, err := ps.db.Exec(`
			DELETE FROM first_meal WHERE id = $1
		`, foodID)
		if err != nil {
			return err
		}
	}
	if cetegory == types.SecondMeal {
		_, err := ps.db.Exec(`
			DELETE FROM second_meal WHERE id = $1
		`, foodID)
		if err != nil {
			return err
		}
	}
	if cetegory == types.Salad {
		_, err := ps.db.Exec(`
			DELETE FROM salad WHERE id = $1
		`, foodID)
		if err != nil {
			return err
		}
	}
	if cetegory == types.Dessert {
		_, err := ps.db.Exec(`
			DELETE FROM dessert WHERE id = $1
		`, foodID)
		if err != nil {
			return err
		}
	}
	if cetegory == types.Beverage {
		_, err := ps.db.Exec(`
			DELETE FROM beverage WHERE id = $1
		`, foodID)
		if err != nil {
			return err
		}
	}

	return nil
}

// sunbula *****************************************************************************************************************

// doniyor *****************************************************************************************************************

func (ps *PostgresRepository) Sets(cash float64) ([][]types.Food, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var sets [][]types.Food
	meal1 := make([]types.Food, 1, 1)
	meal2 := make([]types.Food, 1, 1)
	desert := make([]types.Food, 1, 1)
	salad := make([]types.Food, 1, 1)
	cashmeal1 := cash * 0.5
	cashmeal2 := cash * 0.5
	cashdesert := cash * 0.2
	cashsalad := cash * 0.3

	firstmeal, err := tx.Query(`SELECT * FROM first_meal WHERE price < $1`, cashmeal1)
	if errors.Is(err, sql.ErrNoRows) {
		meal1 = nil
	} else if err != nil {
		tx.Rollback()
		return nil, err
	} else {
		for firstmeal.Next() {
			f := types.Food{}
			if err := firstmeal.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt); err != nil {
				tx.Rollback()
				return nil, err
			}
			f.Category = types.FirstMeal
			meal1 = append(meal1, f)
		}
	}

	if meal1 != nil {
		goto desert
	} else {
		secondmeal, err := tx.Query(`SELECT * FROM second_meal WHERE price < $1`, cashmeal2)
		if errors.Is(err, sql.ErrNoRows) {
			meal2 = nil
		} else if err != nil {
			tx.Rollback()
			return nil, err
		} else {
			for secondmeal.Next() {
				f := types.Food{}
				if err := secondmeal.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt); err != nil {
					tx.Rollback()
					return nil, err
				}
				f.Category = types.SecondMeal
				meal2 = append(meal2, f)
			}
		}
	}

desert:
	if meal2 == nil {
		cashdesert = cash * 0.5
	}
	deserts, err := tx.Query(`SELECT * FROM dessert WHERE price < $1`, cashdesert)
	if errors.Is(err, sql.ErrNoRows) {
		desert = nil
	} else if err != nil {
		tx.Rollback()
		return nil, err
	} else {
		for deserts.Next() {
			f := types.Food{}
			if err := deserts.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt); err != nil {
				tx.Rollback()
				return nil, err
			}
			f.Category = types.Dessert
			desert = append(desert, f)
		}
	}
	if meal1 == nil && meal2 == nil {
		cashsalad = cash * 0.5
		if desert == nil {
			cashsalad = cash
		}
	}
	salads, err := tx.Query(`SELECT * FROM salad WHERE price < $1`, cashsalad)
	if errors.Is(err, sql.ErrNoRows) {
		salad = nil
	} else if err != nil {
		tx.Rollback()
		return nil, err
	} else {
		for salads.Next() {
			f := types.Food{}
			if err := salads.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.Quantity, &f.CookedAt); err != nil {
				tx.Rollback()
				return nil, err
			}
			f.Category = types.Salad
			salad = append(salad, f)
		}
	}
	sets = append(sets, meal1, meal2, salad, desert)
	return sets, nil
}

// doniyor *****************************************************************************************************************
