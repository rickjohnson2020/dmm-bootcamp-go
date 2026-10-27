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

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository (constructor, factory function)
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("find by username: %w", err)
	}

	return entity, nil
}

func (r *account) Create(ctx context.Context, account *object.Account) (*object.Account, error) {
	result, err := r.db.ExecContext(ctx, "insert into account (username, password_hash, display_name, avatar, header, note) values (?, ?, ?, ?, ?, ?)",
		account.Username, account.PasswordHash, account.DisplayName, account.Avatar, account.Header, account.Note)
	if err != nil {
		return nil, fmt.Errorf("insert account: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("get inserted account ID: %w", err)
	}

	entity := new(object.Account)
	err = r.db.QueryRowxContext(ctx, "select * from account where id = ?", id).StructScan(entity)
	if err != nil {
		return nil, fmt.Errorf("fetch created account: %w", err)
	}
	return entity, nil
}
