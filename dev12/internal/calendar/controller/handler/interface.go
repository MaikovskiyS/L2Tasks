package handler

import (
	"context"
	"dev11/internal/calendar/domain/entity"
	"dev11/internal/calendar/domain/params"
)

type Service interface {
	GetEventsByPeriod(ctx context.Context, param params.Period) ([]*entity.Event, error)
	Create(ctx context.Context, e *entity.Event) error
	Update(ctx context.Context, p params.Fielder) error
	Remove(ctx context.Context, id string) error
}
