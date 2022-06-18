package http

import (
	"encoding/json"
	"errors"
	"fmt"
	gohttp "net/http"
)

type Handler interface {
	GetPath() string
	gohttp.Handler
}

type Method string

const (
	GET    Method = "GET"
	PUT    Method = "PUT"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

// private mux instead of using http's default mux
var mux = gohttp.NewServeMux()

func ListenAndServe(addr string) error {
	return gohttp.ListenAndServe(addr, mux)
}

func RegisterHandler(handler Handler) {
	mux.Handle(handler.GetPath(), handler)
}

func GetQueryParameter(r *gohttp.Request, key string) (string, error) {
	p := r.URL.Query().Get(key)
	if p == "" {
		return "", errors.New(fmt.Sprintf("query parameter '%s' is missing", key))
	}

	return p, nil
}

func WriteResponse(w gohttp.ResponseWriter, code int, out interface{}) {
	w.WriteHeader(code)

	if out != nil {
		r, err := json.Marshal(out)
		if err != nil {
			WriteError(w, gohttp.StatusInternalServerError, err)
			return
		}

		_, _ = w.Write(r)
	}
}

func WriteError(w gohttp.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(err.Error()))
}
