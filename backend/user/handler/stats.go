package handler

import (
	"P2/backend/infrastructure/http"
	"P2/backend/user/api"
	"errors"
	gohttp "net/http"
)

type StatsHandler struct{}

func (s StatsHandler) GetPath() string {
	return "/stats"
}

func (s StatsHandler) ServeHTTP(w gohttp.ResponseWriter, r *gohttp.Request) {
	switch http.Method(r.Method) {
	case http.GET: // get a user's stats - (not used)
		id, err := http.GetQueryParameter(r, "id")
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		stats, err := api.GetStats(r.Context(), id)
		if err != nil {
			http.WriteError(w, gohttp.StatusBadRequest, err)
			return
		}

		http.WriteResponse(w, gohttp.StatusOK, stats)
	default:
		http.WriteError(w, 404, errors.New("unsupported method type"))
		return
	}
}
