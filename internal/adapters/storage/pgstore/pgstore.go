package pgstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/LightAlykard/SupportTgBot/internal/repos/info"
	"github.com/google/uuid"
)

var _ info.InfoStore = &Infos{}

type DBPgUser struct {
	ID        uuid.UUID  `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Name      string     `db:"name"`
	Data      string     `db:"data"`
}

type DBPgDeal struct {
	ID        uuid.UUID  `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	IDUser    string     `db:"iduser"`
	Data      string     `db:"data"`
}

type Infos struct {
	db *sql.DB
}

func NewInfos(dsn string) (*Infos, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS public.users (
		id uuid NOT NULL,
		created_at timestamptz NOT NULL,
		updated_at timestamptz NOT NULL,
		deleted_at timestamptz NULL,
		name varchar NOT NULL,
		"data" varchar NULL,		
		CONSTRAINT users_pk PRIMARY KEY (id)
	);`)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS public.deals (
		id uuid NOT NULL,
		created_at timestamptz NOT NULL,
		updated_at timestamptz NOT NULL,
		deleted_at timestamptz NULL,
		iduser varchar NOT NULL,
		"data" varchar NULL,		
		CONSTRAINT users_pk PRIMARY KEY (id)
	);`)

	if err != nil {
		db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	is := &Infos{
		db: db,
	}
	return is, nil
}

func (us *Infos) Close() {
	us.db.Close()
}

func (us *Infos) ReadUser(ctx context.Context, uid uuid.UUID) (*info.User, error) {
	dbu := &DBPgUser{}
	rows, err := us.db.QueryContext(ctx, `SELECT id, created_at, updated_at, deleted_at, name, data 
	FROM users WHERE id = $1`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&dbu.ID,
			&dbu.CreatedAt,
			&dbu.UpdatedAt,
			&dbu.DeletedAt,
			&dbu.Name,
			&dbu.Data,
		); err != nil {
			return nil, err
		}
	}

	return &info.User{
		ID:   dbu.ID,
		Name: dbu.Name,
		Data: dbu.Data,
	}, nil
}

func (us *Infos) ReadDeal(ctx context.Context, uid uuid.UUID) (*info.Deal, error) {
	dbu := &DBPgDeal{}
	rows, err := us.db.QueryContext(ctx, `SELECT id, created_at, updated_at, deleted_at, iduser, data 
	FROM users WHERE id = $1`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&dbu.ID,
			&dbu.CreatedAt,
			&dbu.UpdatedAt,
			&dbu.DeletedAt,
			&dbu.IDUser,
			&dbu.Data,
		); err != nil {
			return nil, err
		}
	}

	return &info.Deal{
		ID:     dbu.ID,
		IDUser: dbu.IDUser,
		Data:   dbu.Data,
	}, nil
}
