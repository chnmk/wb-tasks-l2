package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type response struct {
	Result result `json:"result"`
}

type result struct {
	Events []Event `json:"events"`
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("wrong method")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	r.ParseForm()

	uid, err := validateUid(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	desc, err := validateDesc(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	date, err := validateDate(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	storage.CreateEvent(uid, desc, date)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"result\": \"%s\"}", "success")))
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("wrong method")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	r.ParseForm()

	id, err := validateId(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	uid, err := validateUid(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	desc, err := validateDesc(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	date, err := validateDate(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	err = storage.UpdateEvent(id, uid, desc, date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"result\": \"%s\"}", "success")))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("wrong method")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	r.ParseForm()

	id, err := validateId(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	err = storage.DeleteEvent(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"result\": \"%s\"}", "success")))
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New("wrong method")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	r.ParseForm()

	uid, err := validateUid(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	var res response
	res.Result.Events, err = storage.GetEventsForDate(uid, 0, 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	resp, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New("wrong method")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	r.ParseForm()

	uid, err := validateUid(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	var res response
	res.Result.Events, err = storage.GetEventsForDate(uid, 0, 7)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	resp, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New("wrong method")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	r.ParseForm()

	uid, err := validateUid(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	var res response
	res.Result.Events, err = storage.GetEventsForDate(uid, 1, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	resp, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
