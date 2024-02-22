package handler

import (
	"fmt"
	"net/http"
	"server/http/middleware"
	"server/http/render"
	"server/pkg/service"
	"strconv"
	"time"
)

type Handler struct {
	Store service.Storage
}

func (h *Handler) InitRoutes() error {
	http.HandleFunc("/create_event", middleware.Logger(h.create))
	http.HandleFunc("/update_event", middleware.Logger(h.update))
	http.HandleFunc("/delete_event", middleware.Logger(h.update))
	http.HandleFunc("/events_for_day", middleware.Logger(h.update))
	http.HandleFunc("/events_for_week", middleware.Logger(h.update))
	http.HandleFunc("/events_for_month", middleware.Logger(h.update))

	return nil
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "bad method")
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	date := r.FormValue("date")
	t, err := time.Parse(time.DateOnly, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use RFC3339 format")
		return
	}

	title := r.FormValue("title")
	if title == "" {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("empty title"), "no title provided")
		return
	}

	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}

	e := service.Event{
		Id:    id,
		Title: title,
		Date:  t,
	}

	result, err := h.Store.Create(e)
	if err != nil {
		render.ErrorJSON(w, r, service.GetStatusCode(err), err, "can't create event")
		return
	}

	render.RenderResponse(w, r, http.StatusOK, result)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "bad method")
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	date := r.FormValue("date")
	t, err := time.Parse(time.DateOnly, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use RFC3339 format")
		return
	}

	title := r.FormValue("title")
	if title == "" {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("empty title"), "no title provided")
		return
	}

	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}

	e := service.Event{
		Id:    id,
		Title: title,
		Date:  t,
	}

	err = h.Store.Update(e)
	if err != nil {
		render.ErrorJSON(w, r, service.GetStatusCode(err), err, "can't create event")
		return
	}

	render.NoContent(w, r)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "bad method")
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}

	e := service.Event{
		Id: id,
	}

	err = h.Store.Update(e)
	if err != nil {
		render.ErrorJSON(w, r, service.GetStatusCode(err), err, "can't create event")
		return
	}

	render.NoContent(w, r)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "bad method")
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	events := make([]service.Event, 0)

	switch r.URL.Path {
	case "/events_for_day":
		events = h.Store.GetEventsByPeriod(time.Now(), time.Now().AddDate(0, 0, 1))
	case "/events_for_week":
		events = h.Store.GetEventsByPeriod(time.Now(), time.Now().AddDate(0, 0, 7))
	case "/events_for_month":
		events = h.Store.GetEventsByPeriod(time.Now(), time.Now().AddDate(0, 1, 0))
	}

	if len(events) == 0 {
		render.NoContent(w, r)
		return
	}
	render.RenderResponse(w, r, http.StatusOK, events)
}
