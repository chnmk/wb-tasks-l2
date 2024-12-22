package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type result struct {
	Events []Event
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", "wrong method")))
		return
	}

	r.ParseForm()

	uid, err := strconv.Atoi(r.Form["uid"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	date, err := time.Parse(time.DateTime, r.Form["date"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	desc := r.Form["desc"][0]

	storage.CreateEvent(uid, desc, date)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"result\": \"%s\"}", "success")))
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", "wrong method")))
		return
	}

	r.ParseForm()

	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	uid, err := strconv.Atoi(r.Form["uid"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	date, err := time.Parse(time.DateTime, r.Form["date"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	desc := r.Form["desc"][0]

	err = storage.UpdateEvent(id, uid, desc, date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"result\": \"%s\"}", "success")))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", "wrong method")))
		return
	}

	r.ParseForm()

	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	err = storage.DeleteEvent(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"result\": \"%s\"}", "success")))
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", "wrong method")))
		return
	}

	r.ParseForm()

	uid, err := strconv.Atoi(r.Form["uid"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	var result result
	result.Events, err = storage.GetEventsForDate(uid, 0, 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", "wrong method")))
		return
	}

	r.ParseForm()

	uid, err := strconv.Atoi(r.Form["uid"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	var result result
	result.Events, err = storage.GetEventsForDate(uid, 0, 7)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", "wrong method")))
		return
	}

	r.ParseForm()

	uid, err := strconv.Atoi(r.Form["uid"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	var result result
	result.Events, err = storage.GetEventsForDate(uid, 1, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
