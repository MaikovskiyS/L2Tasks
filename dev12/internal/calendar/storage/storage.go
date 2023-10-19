package storage

import (
	"context"
	"dev11/internal/app/apperror"
	"dev11/internal/calendar/domain/entity"
	"dev11/internal/calendar/domain/params"
	"dev11/internal/calendar/service"
	"errors"
	"fmt"
	"sync"
)

type period string

var (
	day   period = "day"
	week  period = "week"
	month period = "month"
)

// Storage struct
type storage struct {
	sync.Mutex
	events map[string]*entity.Event
}

// NewStorage Create new in-memory storage
func New() service.Storage {
	eventStorage := make(map[string]*entity.Event, 50)
	return &storage{events: eventStorage}
}

// GetByPeriod Return list of events by period
func (s *storage) GetByPeriod(ctx context.Context, p params.Period) ([]*entity.Event, error) {
	s.Lock()
	defer s.Unlock()
	events := make([]*entity.Event, 0)
	interval, value := p.Period()
	switch period(interval) {
	case day:
		for _, e := range s.events {
			if e.Date == value[0] {
				events = append(events, e)
			}
		}
	case week:
		for _, event := range s.events {
			if event.Date.Before(value[1]) && event.Date.After(value[0]) {
				events = append(events, event)
			}
		}
	case month:
		for _, e := range s.events {

			if e.Date.Month() == value[0].Month() {
				events = append(events, e)
			}
		}
	default:
		return nil, apperror.NotFoundErr("storage.GetByPeriod.Switch", fmt.Errorf("wrong interval: %s", interval))

	}
	if len(events) == 0 {
		return nil, apperror.NotFoundErr("storage.GetByPeriod.CheckLen", errors.New("event notFound"))
	}
	return events, nil
}

// Save Create or update event in storage
func (s *storage) Save(ctx context.Context, e *entity.Event) error {
	s.Lock()
	defer s.Unlock()
	s.events[e.Id] = e
	return nil
}

// Remove event from storage
func (s *storage) Remove(ctx context.Context, id string) error {
	s.Lock()
	defer s.Unlock()
	_, ok := s.events[id]
	if ok {
		delete(s.events, id)
		return nil
	}
	return apperror.NotFoundErr("storage.Remove", errors.New("id notFound"))
}
func (s *storage) Update(ctx context.Context, p params.Fielder) error {
	s.Lock()
	defer s.Unlock()
	event, ok := s.events[p.Uuid()]
	if ok {
		for _, f := range p.GetFields() {
			switch f.Name {
			case "name":
				event.Name = f.Value
			case "desc":
				event.Description = f.Value
			case "date":
				event.Date = f.DateValue
			}
		}
		s.events[p.Uuid()] = event
		return nil
	}

	return apperror.NotFoundErr("storage.Update", errors.New("id notFound"))
}
