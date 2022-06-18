package handler

import (
	"P2/backend/game/api"
	"P2/backend/infrastructure/http"
	"encoding/json"
	"errors"
	gohttp "net/http"
)

type GamesHandler struct{}

func (g GamesHandler) GetPath() string {
	return "/games"
}

func (g GamesHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET: // get a game
		id, err := http.GetQueryParameter(r, "id")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		game, err := api.GetGame(r.Context(), id)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		http.WriteResponse(w, gohttp.StatusOK, game)
	case http.POST: // create a game - (not used)
		var board api.Board
		err := json.NewDecoder(r.Body).Decode(&board)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		ng, err := api.CreateGame(r.Context(), board)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		http.WriteResponse(w, gohttp.StatusOK, ng)

	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}
