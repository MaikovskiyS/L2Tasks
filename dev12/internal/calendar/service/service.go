package service

import (
	"context"
	"dev11/internal/calendar/controller/handler"
	"dev11/internal/calendar/domain/entity"
	"dev11/internal/calendar/domain/params"
)

type service struct {
	storage Storage
}

// NewCalendarService Init new calendar service instance
func New(s Storage) handler.Service {
	return &service{storage: s}
}

// GetEventsByPeriod Return list of events by period
func (s *service) GetEventsByPeriod(ctx context.Context, param params.Period) ([]*entity.Event, error) {

	return s.storage.GetByPeriod(ctx, param)
}

// CreateEvent Create new event and save it to storage
func (s *service) Create(ctx context.Context, e *entity.Event) error {
	if err := s.storage.Save(ctx, e); err != nil {
		return err
	}
	return nil
}

// UpdateEvent Update existing event
func (s *service) Update(ctx context.Context, p params.Fielder) error {
	return s.storage.Update(ctx, p)
}

// RemoveEvent Remove existing event
func (s *service) Remove(ctx context.Context, id string) error {
	return s.storage.Remove(ctx, id)
}
