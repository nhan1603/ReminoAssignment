package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/ReminoAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/ReminoAssignment/api/internal/controller/auth"
	authHandler "github.com/nhan1603/ReminoAssignment/api/internal/handler/rest/public/v1/auth"
)

type router struct {
	ctx      context.Context
	authCtrl auth.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.authenticated)
	r.Group(rtr.public)
}

func (rtr router) authenticated(r chi.Router) {
	prefix := "/api/authenticated"

	r.Group(func(r chi.Router) {
		r.Use(httpserver.AuthenticateUserMiddleware())
		prefix = prefix + "/v1"

	})
}

func (rtr router) public(r chi.Router) {
	prefix := "/api/public"

	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"

		r.Group(func(r chi.Router) {
			authH := authHandler.New(rtr.authCtrl)
			r.Post(prefix+"/auth/ou", authH.AuthenticateOperationUser())
		})
	})
}
