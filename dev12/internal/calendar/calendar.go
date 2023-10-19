package calendar

import (
	"dev11/internal/calendar/controller/handler"
	"dev11/internal/calendar/service"
	"dev11/internal/calendar/storage"
	"net/http"
)

func New(router *http.ServeMux) handler.Service {
	storage := storage.New()
	svc := service.New(storage)
	handler := handler.New(svc)
	handler.RegisterRoutes(router)
	return svc
}
