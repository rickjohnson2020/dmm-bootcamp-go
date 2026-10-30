package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	Create(ctx context.Context, status *object.Status) (*object.Status, error)

	FindByID(ctx context.Context, id object.StatusID) (*object.Status, error)
}
