package video

import (
	"encoding/json"
	"net/http"

	"github.com/nhan1603/ReminoAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/ReminoAssignment/api/internal/controller/videos"
)

func (h Handler) ShareVideo() http.HandlerFunc {
	type request struct {
		VideoUrl   string `json:"videoUrl"`
		VideoTitle string `json:"videoTitle"`
	}

	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return err
		}

		err := h.videoCtrl.ShareVideo(r.Context(), req.VideoUrl, req.VideoTitle)
		if err != nil {
			switch err {
			case videos.ErrInvalidVideoUrl:
				http.Error(w, "Invalid YouTube video URL", http.StatusBadRequest)
			case videos.ErrUserNotFound:
				http.Error(w, "Target user does not exist", http.StatusBadRequest)
			default:
				http.Error(w, "Failed to share video", http.StatusInternalServerError)
			}
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	})
}
