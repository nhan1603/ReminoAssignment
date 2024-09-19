package auth

import (
	"encoding/json"
	"net/http"

	"github.com/nhan1603/ReminoAssignment/api/internal/model"
)

func (h Handler) CreateUser() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user := model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		err := h.authCtrl.CreateUser(r.Context(), &user)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}
