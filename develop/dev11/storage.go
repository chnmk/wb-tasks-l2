package main

import "time"

type AbstractStorage interface {
	NewStorage() *Storage
}

type Storage struct {
	Events []Event
}

type Event struct {
	Id          int
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
	e.Id = max(e.Id) + 1

	s.Events = append(s.Events, e)
}
