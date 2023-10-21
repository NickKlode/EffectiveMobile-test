package postgres

import (
	"context"
	"emobletest/internal/storage/model"
	"emobletest/lib/api"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(connStr string) (*DB, error) {

	c, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	s := DB{
		pool: c,
	}
	return &s, nil
}

func (db *DB) CreateUser(u model.User) (int, error) {
	age, err := api.GetAge(u.Name)
	if err != nil {
		return 0, err
	}
	gender, err := api.GetGender(u.Name)
	if err != nil {
		return 0, err
	}
	nat, err := api.GetNationality(u.Name)
	if err != nil {
		return 0, err
	}
	var id int

	if u.Patronymic == "" {
		query := `INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		err = db.pool.QueryRow(context.Background(), query, u.Name, u.Surname, "", age, gender, nat).Scan(&id)
		if err != nil {
			return 0, err
		}
	} else {
		query := `INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		err = db.pool.QueryRow(context.Background(), query, u.Name, u.Surname, u.Patronymic, age, gender, nat).Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (db *DB) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := db.pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateUser(id int, ui model.UpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if ui.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *ui.Name)
		argId++
	}
	if ui.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("suranme=$%d", argId))
		args = append(args, *ui.Surname)
		argId++
	}
	if ui.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *ui.Patronymic)
		argId++
	}
	if ui.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *ui.Age)
		argId++
	}
	if ui.Gender != nil {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *ui.Gender)
		argId++
	}
	if ui.Nationality != nil {
		setValues = append(setValues, fmt.Sprintf("nationality=$%d", argId))
		args = append(args, *ui.Nationality)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)
	_, err := db.pool.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetUser(gi model.GetInput) ([]model.User, error) {
	var su []model.User
	if gi == (model.GetInput{}) {
		query := fmt.Sprintf("SELECT * FROM users")
		rows, err := db.pool.Query(context.Background(), query)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var s model.User
			err = rows.Scan(&s.ID, &s.Name, &s.Surname, &s.Patronymic, &s.Age, &s.Gender, &s.Nationality)
			if err != nil {
				return nil, err
			}
			su = append(su, s)
		}
	} else {
		setValues := make([]string, 0)
		args := make([]interface{}, 0)
		argId := 1
		if gi.ID != nil {
			setValues = append(setValues, fmt.Sprintf("id=$%d", argId))
			args = append(args, *gi.ID)
			argId++
		}
		if gi.Name != nil {
			setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
			args = append(args, *gi.Name)
			argId++
		}
		if gi.Surname != nil {
			setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
			args = append(args, *gi.Surname)
			argId++
		}
		if gi.Patronymic != nil {
			setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
			args = append(args, *gi.Patronymic)
			argId++
		}
		if gi.Age != nil {
			setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
			args = append(args, *gi.Age)
			argId++
		}
		if gi.Gender != nil {
			setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
			args = append(args, *gi.Gender)
			argId++
		}
		if gi.Nationality != nil {
			setValues = append(setValues, fmt.Sprintf("nationality=$%d", argId))
			args = append(args, *gi.Nationality)
			argId++
		}
		setQuery := strings.Join(setValues, " AND ")
		query := fmt.Sprintf("SELECT * FROM users WHERE %s", setQuery)
		rows, err := db.pool.Query(context.Background(), query, args...)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var s model.User
			err = rows.Scan(&s.ID, &s.Name, &s.Surname, &s.Patronymic, &s.Age, &s.Gender, &s.Nationality)
			if err != nil {
				return nil, err
			}
			su = append(su, s)
		}
	}
	return su, nil
}
