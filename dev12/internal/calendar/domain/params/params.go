package params

import (
	"dev11/internal/app/apperror"
	"net/http"
	"strings"
	"time"
)

const shortForm = "2006-01-02"

type Fielder interface {
	Uuid() string
	GetFields() []Field
	AddField(name, value string)
}
type Period interface {
	Period() (string, []time.Time)
}
type period struct {
	name  string
	value []time.Time
}

func (p *period) Period() (string, []time.Time) {
	return p.name, p.value
}
func NewPeriod(r *http.Request) (Period, error) {
	p := &period{}
	interval := []string{}
	switch r.URL.Path {
	case "/events_for_day":
		interval = append(interval, r.FormValue("day"))
		p.name = "day"

	case "/events_for_week":
		time := r.FormValue("week")
		times := strings.Split(time, ":")
		interval = append(interval, times...)
		p.name = "week"

	case "/events_for_month":
		time := r.FormValue("month")
		time = time + "-01"
		interval = append(interval, time)
		p.name = "month"
	}
	times, err := validateInterval(interval)
	if err != nil {
		return nil, err
	}
	p.value = times

	return p, nil
}
func validateInterval(internals []string) ([]time.Time, error) {
	times := make([]time.Time, 0)
	for _, interval := range internals {
		t, err := time.Parse(shortForm, interval)
		if err != nil {
			return nil, apperror.BadRequestErr("params.validateInterval.timeParse", err)
		}
		times = append(times, t)
	}
	return times, nil
}

type fields struct {
	id   string
	data []Field
}

func NewFilder(uuid string) Fielder {
	return &fields{id: uuid, data: make([]Field, 0, 3)}
}

type Field struct {
	Name      string
	Value     string
	DateValue time.Time
}

func (f *fields) Uuid() string {
	return f.id
}
func (f *fields) GetFields() []Field {
	return f.data
}
func (f *fields) AddField(name string, value string) {
	fl := Field{Name: name}
	if name == "date" {
		date, _ := time.Parse(shortForm, value)
		fl.DateValue = date
	} else {
		fl.Value = value
	}
	f.data = append(f.data, fl)
}
