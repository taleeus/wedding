package server

import (
	"errors"
	"net/http"
)

type HandlerFuncEx[T any] func(http.ResponseWriter, *http.Request, T)

func extend[T any](handler HandlerFuncEx[T], ex T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ex)
	}
}

func recoverable(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}

				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		handler(w, r)
	}
}
