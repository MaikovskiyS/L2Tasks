package service

import (
	"context"
	"dev11/internal/calendar/domain/entity"
	"dev11/internal/calendar/domain/params"
)

type Storage interface {
	GetByPeriod(ctx context.Context, param params.Period) ([]*entity.Event, error)
	Save(ctx context.Context, event *entity.Event) error
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, p params.Fielder) error
}
