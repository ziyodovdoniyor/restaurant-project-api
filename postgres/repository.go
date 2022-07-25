package postgres

import (
	"database/sql"
	"restaurant/types"

)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

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
		err := rows.Scan(&f.ID, &f.Name, &f.Ingredients, &f.Price, &f.CookedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foods = append(foods, f)
	}

	rows, err = tx.Query(`SELECT * FROM second_meal`)
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

	rows, err = tx.Query(`SELECT * FROM dessert`)
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

	rows, err = tx.Query(`SELECT * FROM salad`)
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

	rows, err = tx.Query(`SELECT * FROM beverage`)
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


func (ps *PostgresRepository) AddFood(f types.Food) error {
	if f.Category == types.FirstMeal {
		_, err := ps.db.Exec(`
			INSERT INTO first_meal (id, name, ingredients, is_speacil, price, cooked_at)
			VALUES ($1, $2, $3, $4, $5)
		`, f.ID, f.Name, f.Ingredients,  f.Price, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.SecondMeal {
		_, err := ps.db.Exec(`
			INSERT INTO second_meal (id, name, ingredients, is_speacil, price, cooked_at)
			VALUES ($1, $2, $3, $4, $5)
		`, f.ID, f.Name, f.Ingredients, f.Price, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.Salad {
		_, err := ps.db.Exec(`
		INSERT INTO salad (id, name, ingredients, is_speacil, price, cooked_at)
		VALUES ($1, $2, $3, $4, $5)
	`, f.ID, f.Name, f.Ingredients,  f.Price, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.Dessert {
		_, err := ps.db.Exec(`
		INSERT INTO dessert (id, name, ingredients, is_speacil, price, cooked_at)
		VALUES ($1, $2, $3, $4, $5)
	`, f.ID, f.Name, f.Ingredients, f.Price, f.CookedAt)

		if err != nil {
			return err
		}
	} else if f.Category == types.Beverage {
		_, err := ps.db.Exec(`
		INSERT INTO beverage(id, name, ingredients, is_speacial, price, cooked_at)
		VALUES ($1, $2, $3, $4, $5)
	`, f.ID, f.Name, f.Ingredients,  f.Price, f.CookedAt)

		if err != nil {
			return err
		}
	}

	return nil
}
