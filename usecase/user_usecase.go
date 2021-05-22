package usecase

import (
	"context"
	"github.com/noguchidaisuke/clean-arch-by-budougumi-repository/entity"
	"github.com/noguchidaisuke/clean-arch-by-budougumi-repository/usecase/port"
)

type UserCase struct {
	ua port.UserAccessor
}

func NewUserCase(ua port.UserAccessor) *UserCase {
	return &UserCase {
		ua: ua,
	}
}

func (uc UserCase) FindUser(ctx context.Context, id int64) (user *entity.User, err error) {
	user, err = uc.ua.FindUser(ctx, id)
	return
}