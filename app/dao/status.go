package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type status struct {
	db *sqlx.DB
}

// constructor
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) Create(ctx context.Context, status *object.Status) (*object.Status, error) {
	result, err := r.db.ExecContext(ctx, "insert into status (account_id, content) values (?, ?)",
		status.AccountID, status.Content)
	if err != nil {
		return nil, fmt.Errorf("insert status: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("get inserted status ID: %w", err)
	}
	entity := new(object.Status)
	err = r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		return nil, fmt.Errorf("fetch created status: %w", err)
	}
	return entity, nil
}

func (r *status) FindByID(ctx context.Context, id object.StatusID) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("find by id: %w", err)
	}
	account := new(object.Account)
	err = r.db.QueryRowxContext(ctx, "select * from account where id = ?", entity.AccountID).StructScan(account)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no account found: %w", err)
		}
		return nil, fmt.Errorf("fetch account by id: %w", err)
	}
	entity.Account = account
	return entity, nil
}

func (r *status) Delete(ctx context.Context, id object.StatusID) error {
	_, err := r.db.ExecContext(ctx, "delete from status where id = ?", id)
	if err != nil {
		return fmt.Errorf("delete status: %w", err)
	}
	return nil
}
