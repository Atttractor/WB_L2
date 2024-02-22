package store

import (
	"fmt"
	"server/pkg/service"
	"time"
)

type InMemoryStore struct {
	events map[int]service.Event
}

func NewInMemoryStore() InMemoryStore {
	return InMemoryStore{events: make(map[int]service.Event)}
}

func (s *InMemoryStore) Create(e service.Event) (service.Event, error) {
	if _, ok := s.events[e.Id]; !ok {
		s.events[e.Id] = e
		return e, nil
	}

	return service.Event{}, fmt.Errorf("event with the same id already exists: %v", e.Id)
}

func (s *InMemoryStore) Update(e service.Event) error {
	if _, ok := s.events[e.Id]; !ok {
		return service.ErrNotFound
	}

	s.events[e.Id] = e
	return nil
}

func (s *InMemoryStore) Delete(e service.Event) error {
	if _, ok := s.events[e.Id]; !ok {
		return service.ErrNotFound
	}

	delete(s.events, e.Id)
	return nil
}

func (s *InMemoryStore) GetEventsByPeriod(start time.Time, end time.Time) []service.Event {
	var result []service.Event
	for _, event := range s.events {
		if event.Date.After(start) && event.Date.Before(end) {
			result = append(result, event)
		}
	}

	return result
}
