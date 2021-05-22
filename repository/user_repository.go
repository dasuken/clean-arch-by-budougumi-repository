package repository

import (
	"context"
	"database/sql"
	"github.com/noguchidaisuke/clean-arch-by-budougumi-repository/entity"
)

func (repo *Repo) FindUser(ctx context.Context, id int64) (*entity.User, error) {
	u := &entity.User{}
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.QueryRowContext(ctx, `
		SELECT id, name, email FROM user WHERE id = ?
	`, id).Scan(
		&u.ID, &u.Name, &u.Email,
		)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}

	return u, nil
}