package handler

import (
	"P2/backend/game/api"
	"P2/backend/infrastructure/http"
	"errors"
	gohttp "net/http"
	"strconv"
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
	case http.POST: // create a custom game - (not used)
		sRow, err := http.GetQueryParameter(r, "r")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}
		row, err := strconv.Atoi(sRow)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		sCol, err := http.GetQueryParameter(r, "c")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}
		col, err := strconv.Atoi(sCol)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		skVal, err := http.GetQueryParameter(r, "k")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}
		kVal, err := strconv.Atoi(skVal)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		ng, err := api.CreateGame(r.Context(), row, col, kVal)
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
