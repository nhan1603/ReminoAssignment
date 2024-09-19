package user

import (
	"context"
	"log"

	"github.com/nhan1603/ReminoAssignment/api/internal/model"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Create create user by input
func (i impl) Create(ctx context.Context, user *model.User) error {
	dbUser := dbmodel.User{
		Email:        user.Email,
		PasswordHash: user.Password,
	}
	err := dbUser.Insert(ctx, i.dbConn, boil.Infer())
	log.Println("Finish insert with values " + dbUser.Email)

	if err != nil {
		pkgerrors.WithStack(err)
	}

	return nil
}
