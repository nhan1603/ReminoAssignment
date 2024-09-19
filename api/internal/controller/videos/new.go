package videos

import (
	"context"

	"github.com/nhan1603/ReminoAssignment/api/internal/model"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	ShareVideo(ctx context.Context, videoUrl, videoTitle string) error
	GetSharedVideo(ctx context.Context) (model.ListSharedVideo, error)
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
