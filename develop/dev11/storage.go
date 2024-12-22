package main

import (
	"errors"
	"time"
)

type AbstractStorage interface {
	CreateEvent(uid int, desc string, date time.Time)
	UpdateEvent(id int, uid int, desc string, date time.Time) error
	DeleteEvent(id int) error
	GetEventsForDate(uid int, month int, day int) ([]Event, error)
}

type Storage struct {
	Events []Event
}

type Event struct {
	UserId      int
	Description string
	Date        time.Time
}

func ReturnStorage() *Storage {
	return &Storage{}
}

// POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month

func (s *Storage) CreateEvent(uid int, desc string, date time.Time) {
	var e Event

	e.UserId = uid
	e.Description = desc
	e.Date = date

	s.Events = append(s.Events, e)
}

func (s *Storage) UpdateEvent(id int, uid int, desc string, date time.Time) error {
	if id >= len(s.Events) {
		return errors.New("event with this id doesn't exist")
	}

	s.Events[id].UserId = uid
	s.Events[id].Description = desc
	s.Events[id].Date = date

	return nil
}

func (s *Storage) DeleteEvent(id int) error {
	if id >= len(s.Events) {
		return errors.New("event with this id doesn't exist")
	}

	s.Events = append(s.Events[:id], s.Events[id+1:]...)

	return nil
}

func (s *Storage) GetEventsForDate(uid int, month int, day int) ([]Event, error) {
	now := time.Now()
	d := now.AddDate(0, month, day)

	var result []Event
	for _, e := range s.Events {

		if e.UserId == uid && e.Date.Before(d) && now.Before(e.Date) {
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return result, errors.New("events not found")
	}

	return result, nil
}
