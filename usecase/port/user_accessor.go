package port

import (
	"context"
	"github.com/noguchidaisuke/clean-arch-by-budougumi-repository/entity"
)

type UserAccessor interface {
	UserReader
}

// UserReader retrieves User data from a data store.
type UserReader interface {
	FindUser(context.Context, int64) (*entity.User, error)
}

