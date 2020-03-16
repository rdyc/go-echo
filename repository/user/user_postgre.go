package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	models "github.com/rdyc/go-echo/entities"
	pRepo "github.com/rdyc/go-echo/repository"
)

// NewSQLUserRepo retunrs implement of post repository interface
func NewSQLUserRepo(Conn *sql.DB) pRepo.UserRepo {
	return &postgreUserRepo{
		Conn: Conn,
	}
}

type postgreUserRepo struct {
	Conn *sql.DB
}

func (m *postgreUserRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.Id,
			&data.UserName,
			&data.Email,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *postgreUserRepo) Fetch(ctx context.Context, num int64) ([]*models.User, error) {
	query := `select "Id", "UserName", "Email" from "Users" limit $1;`

	return m.fetch(ctx, query, num)
}

func (m *postgreUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `select "Id", "UserName", "Email" from "Users" where "Id"=$1`

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *postgreUserRepo) Create(ctx context.Context, p *models.User) (uuid.UUID, error) {
	query := `insert into "Users" values ($1, $2, $3)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = stmt.ExecContext(ctx, p.Id, p.UserName, p.Email)
	defer stmt.Close()

	if err != nil {
		return p.Id, err
	}

	return uuid.Nil, nil
}

func (m *postgreUserRepo) Update(ctx context.Context, p *models.User) (*models.User, error) {
	query := `update "Users" set "UserName"=$2, "Email"=$3 where "Id"=$1`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Id,
		p.UserName,
		p.Email,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *postgreUserRepo) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `delete from "Users" where "Id"=$1`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
