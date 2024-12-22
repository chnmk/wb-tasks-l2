package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func validateId(w http.ResponseWriter, r *http.Request) (int, error) {
	if len(r.Form["id"]) == 0 {
		err := errors.New("id param not found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return 0, err
	}

	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return id, err
	}

	return id, nil
}

func validateUid(w http.ResponseWriter, r *http.Request) (int, error) {
	if len(r.Form["uid"]) == 0 {
		err := errors.New("uid param not found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err)))
		return 0, err
	}

	uid, err := strconv.Atoi(r.Form["uid"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return uid, err
	}

	return uid, nil
}

func validateDesc(w http.ResponseWriter, r *http.Request) (string, error) {
	if len(r.Form["desc"]) == 0 {
		err := errors.New("desc param not found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return "", err
	}

	desc := r.Form["desc"][0]

	return desc, nil
}

func validateDate(w http.ResponseWriter, r *http.Request) (time.Time, error) {
	if len(r.Form["date"]) == 0 {
		err := errors.New("date param not found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err)))
		return time.Now(), err
	}

	date, err := time.Parse(time.DateTime, r.Form["date"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())))
		return date, err
	}

	return date, nil
}
