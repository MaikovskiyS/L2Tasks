package app

import (
	"dev11/internal/app/config"
	"dev11/internal/calendar"
	"net/http"
)

func Run(cfg *config.Config) error {
	router := http.NewServeMux()

	calendar.New(router)

	err := http.ListenAndServe(cfg.Port(), router)
	if err != nil {
		return err
	}
	return nil
}
