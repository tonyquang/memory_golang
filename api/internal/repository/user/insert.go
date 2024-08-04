package user

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"memory_golang/api/internal/model"
	"memory_golang/api/internal/repository/orm"
)

func (i impl) Insert(ctx context.Context, user model.User) (model.User, error) {
	ormModel := orm.User{Email: user.Email}

	if err := ormModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.User{}, err
	}

	user.ID = ormModel.ID
	return user, nil
}
