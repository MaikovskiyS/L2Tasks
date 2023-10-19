package handler

import (
	"dev11/internal/app/apperror"
	"errors"
	"log"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func Logging(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *apperror.AppError
		resp := Resp{}
		log.Println("incoming Request: ", r.Method, r.URL.Path)
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {

				err := err.(*apperror.AppError)
				switch err.Code {
				case apperror.ErrBadRequest:
					w.WriteHeader(http.StatusBadRequest)
				case apperror.ErrNotFound:
					w.WriteHeader(http.StatusNotFound)
				case apperror.ErrInternal:
					w.WriteHeader(http.StatusInternalServerError)
				}
				resp.SetErr(err.Error())
				w.Write(resp.Bytes())
				log.Printf("errLocation: %s errMsg: %s", err.Location, err.Error())
				return
			}
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(err.Error()))

		}
		log.Println("Request success")
	}
}
