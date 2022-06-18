package handler

import (
	"P2/backend/infrastructure/http"
	"P2/backend/user/api"
	"errors"
	gohttp "net/http"
)

type UsersHandler struct{}

func (u UsersHandler) GetPath() string {
	return "/users"
}

func (u UsersHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET: // get a user
		id, err := http.GetQueryParameter(r, "id")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		user, err := api.GetUser(r.Context(), id)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		http.WriteResponse(w, gohttp.StatusOK, user)
	case http.POST: // login
		id, err := http.GetQueryParameter(r, "id")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		user, err := api.Login(r.Context(), id)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		http.WriteResponse(w, gohttp.StatusOK, user)

	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}
