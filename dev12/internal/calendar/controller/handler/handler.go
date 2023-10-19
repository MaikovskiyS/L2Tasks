package handler

import (
	"context"
	"dev11/internal/app/apperror"
	"dev11/internal/calendar/domain/params"
	"encoding/json"
	"net/http"
	"time"
)

type Handler interface {
	RegisterRoutes(router *http.ServeMux)
}
type handler struct {
	timeout time.Duration
	svc     Service
}

func New(svc Service) Handler {
	return &handler{
		timeout: time.Second * 2,
		svc:     svc,
	}
}
func (h *handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/create_event", Logging(h.Create))
	router.HandleFunc("/delete_event", Logging(h.Delete))
	router.HandleFunc("/update_event", Logging(h.Update))
	router.HandleFunc("/events_for_day", Logging(h.Get))
	router.HandleFunc("/events_for_week", Logging(h.Get))
	router.HandleFunc("/events_for_month", Logging(h.Get))

}
func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return apperror.BadRequestErr("handler.create.checkMethod", apperror.ErrMethod)
	}
	var dto *createEventDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return apperror.BadRequestErr("handler.create.Decode", err)
	}
	event, err := dto.Validate()
	if err != nil {
		return err

	}
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	err = h.svc.Create(ctx, event)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	resp := Resp{}
	resp.SetMsg("event created")
	w.Write(resp.Bytes())
	return nil
}
func (h *handler) Get(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return apperror.BadRequestErr("handler.get.checkMethod", apperror.ErrMethod)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	p, err := params.NewPeriod(r)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	events, err := h.svc.GetEventsByPeriod(ctx, p)
	if err != nil {
		return err
	}
	// eventsbytes, err := json.Marshal(&events)
	// if err != nil {
	// 	return apperror.InternalErr("handler.get.jsonMarshal", err)
	// }
	w.WriteHeader(http.StatusOK)
	resp := Resp{}
	resp.SetEvents(events)
	w.Write(resp.Bytes())
	return nil

}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return apperror.BadRequestErr("handler.update.checkMethod", apperror.ErrMethod)
	}
	var dto *updateEventDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return apperror.BadRequestErr("handler.update.Decode", err)
	}
	params, err := dto.Validate()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	err = h.svc.Update(ctx, params)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	resp := Resp{}
	resp.SetMsg("event updated")
	w.Write(resp.Bytes())

	return nil
}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return apperror.BadRequestErr("handler.delete.checkMethod", apperror.ErrMethod)
	}
	var dto *deleteEventDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return apperror.BadRequestErr("handler.delete.Decode", err)
	}
	id, err := dto.Validate()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	err = h.svc.Remove(ctx, id)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	resp := Resp{}
	resp.SetMsg("event deleted")
	w.Write(resp.Bytes())
	return nil
}
