package routines

import (
	"context"
	"github.com/tgmendes/guitrainer/internal/platform/db"
)

type Routine struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

func Add(ctx context.Context, repo db.Repository, r Routine) (*Routine, error) {
	q := `INSERT INTO routines(name, description, level) VALUES($1, $2, $3) RETURNING "routine_id"`

	err := repo.QueryRow(ctx, q, r.Name, r.Description, r.Level).Scan(&r.ID)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func List(ctx context.Context, repo db.Repository) ([]Routine, error) {
	q := `SELECT * FROM routines`
	var routines []Routine

	rows, err := repo.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r Routine

		err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.Level)
		if err != nil {
			return nil, err
		}
		routines = append(routines, r)
	}

	return routines, nil
}

func Get(ctx context.Context, repo db.Repository, ID int64) (*Routine, error) {
	q := `SELECT * FROM routines WHERE routine_id = $1`
	var r Routine

	err := repo.QueryRow(ctx, q, ID).Scan(&r.ID, &r.Name, &r.Description, &r.Level)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func Delete(ctx context.Context, repo db.Repository, ID int64) error {
	q := `DELETE FROM routines WHERE routine_id = $1`
	_, err := repo.Exec(ctx, q, ID)

	return err
}
