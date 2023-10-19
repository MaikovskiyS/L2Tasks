package handler

import (
	"dev11/internal/app/apperror"
	"dev11/internal/calendar/domain/entity"
	"dev11/internal/calendar/domain/params"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

const shortForm = "2006-01-02"

type createEventDTO struct {
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func (d *createEventDTO) Validate() (*entity.Event, error) {
	if d.Name == "" || d.Date == "" {
		return nil, apperror.BadRequestErr("handler.validateDTO.checkParams", errors.New("name and date are required params"))
	}
	date, err := time.Parse(shortForm, d.Date)
	if err != nil {
		return nil, apperror.BadRequestErr("handler.validateDTO.timeParse", err)
	}
	return &entity.Event{Id: uuid.NewV4().String(), Name: d.Name, Date: date, Description: d.Description}, nil
}

type updateEventDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func (u *updateEventDTO) Validate() (params.Fielder, error) {
	if u.Id == "" {
		return nil, apperror.BadRequestErr("handler.validateDTO.checkParams", errors.New("id required"))
	}
	p := params.NewFilder(u.Id)
	if u.Name != "" {
		p.AddField("name", u.Name)
	}
	if u.Description != "" {
		p.AddField("desc", u.Description)
	}
	if u.Date != "" {
		_, err := time.Parse(shortForm, u.Date)
		if err != nil {
			return nil, apperror.BadRequestErr("handler.validateDTO.timeParse", err)
		}
		p.AddField("date", u.Date)
	}
	return p, nil

}

type deleteEventDTO struct {
	Id string `json:"id"`
}

func (d *deleteEventDTO) Validate() (id string, err error) {
	if d.Id == "" {
		return "", apperror.BadRequestErr("handler.validateDTO.checkParams", errors.New("id required"))
	}
	id = d.Id
	return
}
