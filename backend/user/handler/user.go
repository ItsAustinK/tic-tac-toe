package handler

import (
	"P2/backend/infrastructure/http"
	"P2/backend/user/api"
	"errors"
	"fmt"
	gohttp "net/http"
)

type UsersHandler struct{}

func (u UsersHandler) GetPath() string {
	return "/users"
}

func (u UsersHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET:
		u.getUser(w, r)
	case http.POST:
		u.userLogin(w, r)
	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}

func (u UsersHandler) getUser(w gohttp.ResponseWriter, r *gohttp.Request) {
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
}

func (u UsersHandler) userLogin(w gohttp.ResponseWriter, r *gohttp.Request) {
	id, err := http.GetQueryParameter(r, "id")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}
	fmt.Println(fmt.Sprintf("logging in user %s", id))

	user, err := api.Login(r.Context(), id)
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	http.WriteResponse(w, gohttp.StatusOK, user)
}
