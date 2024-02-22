package service

import (
	"errors"
	"net/http"
	"time"
)

type Event struct {
	Id    int
	Title string
	Date  time.Time
}

type Storage interface {
	Create(e Event) (Event, error)
	Update(e Event) error
	Delete(e Event) error
	GetEventsByPeriod(time.Time, time.Time) []Event
}

var ErrNotFound = errors.New("your requested item is not found")

func GetStatusCode(err error) int {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

//func (c *Calendar) createEvent(e Event) (Event, error) {
//	return
//}
//
//func (c *Calendar) updateEvent(e Event) error {
//	return
//}
//
//func (c *Calendar) deleteEvent(e Event) error {
//	return c.storage.delete()
//}
//
//func (c *Calendar) getEventByDay() []Event {
//	return c.storage.getEventsByPeriod(time.Now(), time.Now().AddDate(0, 0, 1))
//}
//
//func (c *Calendar) getEventByWeek() []Event {
//	return c.storage.getEventsByPeriod(time.Now(), time.Now().AddDate(0, 0, 7))
//}
//
//func (c *Calendar) getEventByMonth() []Event {
//	return c.storage.getEventsByPeriod(time.Now(), time.Now().AddDate(0, 1, 0))
//}
