package auth

import (
	"context"

	"github.com/nhan1603/ReminoAssignment/api/internal/model"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	CheckAuth(ctx context.Context, inp LoginInput) (model.User, string, error)
	CreateUser(ctx context.Context, user *model.User) error
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
