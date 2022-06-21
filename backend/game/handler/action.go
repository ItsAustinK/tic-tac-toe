package handler

import (
	"P2/backend/game/api"
	"P2/backend/infrastructure/http"
	"encoding/json"
	"errors"
	gohttp "net/http"
)

type ActionsHandler struct{}

func (a ActionsHandler) GetPath() string {
	return "/actions"
}

func (a ActionsHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.POST:
		a.makeBoardAction(w, r)
	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}

func (a ActionsHandler) makeBoardAction(w gohttp.ResponseWriter, r *gohttp.Request) {
	id, err := http.GetQueryParameter(r, "id")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	token, err := http.GetQueryParameter(r, "token")
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	var action api.Action
	err = json.NewDecoder(r.Body).Decode(&action)
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	err = api.MakePlayerAction(r.Context(), id, token, action)
	if err != nil {
		http.WriteError(w, gohttp.StatusBadRequest, err)
		return
	}

	// TODO: handle if this was a winning action? This should be a long poll or websocket action

	http.WriteResponse(w, gohttp.StatusOK, nil)
}
