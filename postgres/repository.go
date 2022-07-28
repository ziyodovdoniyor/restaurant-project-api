package postgres

import (
	"database/sql"
	"fmt"
	"restaurant/types"
	"time"
)

// PSQL

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p PostgresRepository) Buy(purchase *types.Purchase) (int, error) {
	price := 0

	row, er := p.db.Query("select price from first_meal where id = $1", purchase.FirstMealID)
	if er != nil {
		return 0, er
	}
	if er = row.Scan(&price); er != nil {
		return 0, er
	}
	purchase.Total += price

	row, er = p.db.Query("select price from second_meal where id = $1", purchase.SecondMealID)
	if er != nil {
		return 0, er
	}
	if er = row.Scan(&price); er != nil {
		return 0, er
	}
	purchase.Total += price

	row, er = p.db.Query("select price from beverage where id = $1", purchase.BeverageID)
	if er != nil {
		return 0, er
	}
	if er = row.Scan(&price); er != nil {
		return 0, er
	}
	purchase.Total += price

	row, er = p.db.Query("select price from salad where id = $1", purchase.SaladID)
	if er != nil {
		return 0, er
	}
	if er = row.Scan(&price); er != nil {
		return 0, er
	}
	purchase.Total += price

	row, er = p.db.Query("select price from dessert where id = $1", purchase.DessertID)
	if er != nil {
		return 0, er
	}
	if er = row.Scan(&price); er != nil {
		return 0, er
	}
	purchase.Total += price

	return purchase.Total, nil
}

func (p *PostgresRepository) TakeTable(num int) (types.Table, bool, error) {
	var t types.Table
	row := p.db.QueryRow("select table_number, is_taken from tables where table_number = $1", num)
	if er := row.Scan(&t.Number, &t.IsTaken); er != nil {
		return types.Table{}, false, er
	}
	if t.Number != num {
		return types.Table{}, true, nil
	}
	if _, er := p.db.Exec("update tables set is_taken=$1 where table_number=$2", true, num); er != nil {
		return types.Table{}, true, er
	}
	return t, true, nil
}

func (p *PostgresRepository) GetTables() ([]int, error) {
	rows, er := p.db.Query("select table_number from tables where is_taken = $1", false)
	if er != nil {
		return nil, er
	}
	var tableNumbers []int
	for rows.Next() {
		var num int
		if er = rows.Scan(&num); er != nil {
			return nil, er
		}
		tableNumbers = append(tableNumbers, num)
	}
	return tableNumbers, nil
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
