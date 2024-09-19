package videos

import (
	"context"
	"errors"
	"fmt"

	"github.com/nhan1603/ReminoAssignment/api/internal/appconfig/iam"
	"github.com/nhan1603/ReminoAssignment/api/internal/model"
	"github.com/nhan1603/ReminoAssignment/api/internal/pkg/validator"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository/user"
)

func (i impl) ShareVideo(ctx context.Context, videoUrl, videoTitle string) error {
	// Check if video url is valid youtube url
	if !validator.IsValidYouTubeURL(videoUrl) {
		return ErrInvalidVideoUrl
	}

	userProfile := iam.UserProfileFromContext(ctx)

	// Check if user email is valid user
	shareUser, err := i.repo.User().GetByEmail(ctx, user.GetUserInput{
		Email: userProfile.Email,
	})
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	videoShare := &model.VideoShare{
		VideoUrl:   videoUrl,
		VideoTitle: videoTitle,
		UserID:     int(shareUser.ID),
	}

	err = i.repo.Video().ShareVideo(ctx, *videoShare)
	if err != nil {
		return fmt.Errorf("failed to share video: %w", err)
	}

	return nil
}
