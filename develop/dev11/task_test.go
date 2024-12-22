package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestServerWrongMethod(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/create_event", nil)
	rec := httptest.NewRecorder()

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Fatalf("expected status: %d, got %d", http.StatusBadRequest, status)
	}
}

func TestServerCreate(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(-1 * time.Hour).Format(time.DateTime)

	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("expected status: %d, got %d", http.StatusOK, status)
	}

	if rec.Body.String() != "{\"result\": \"success\"}" {
		t.Errorf("expected result: %s, got %s", "{\"result\": \"success\"}", rec.Body.String())
	}
}

func TestServerCreateInvalid(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(-1 * time.Hour).Format(time.DateTime)

	data := url.Values{}
	data.Set("uid", "1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Fatalf("expected status: %d, got %d", http.StatusBadRequest, status)
	}

	if rec.Body.String() != "{\"error\": \"desc param not found\"}" {
		t.Errorf("expected result: %s, got %s", "{\"error\": \"desc param not found\"}", rec.Body.String())
	}
}

func TestServerUpdate(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(-1 * time.Hour).Format(time.DateTime)

	// Post
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	// Update
	data2 := url.Values{}
	data2.Set("id", "0")
	data2.Set("uid", "1")
	data2.Set("desc", "task1")
	data2.Set("date", n)
	enc2 := data2.Encode()

	req2 := httptest.NewRequest("POST", "/update_event", strings.NewReader(enc2))
	rec2 := httptest.NewRecorder()

	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(data2.Encode())))

	h2 := http.HandlerFunc(UpdateEvent)
	h2.ServeHTTP(rec2, req2)

	if status := rec2.Code; status != http.StatusOK {
		t.Fatalf("expected status: %d, got %d", http.StatusOK, status)
	}

	if rec2.Body.String() != "{\"result\": \"success\"}" {
		t.Errorf("expected result: %s, got %s", "{\"result\": \"success\"}", rec2.Body.String())
	}
}

func TestServerUpdateInvalid(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(-1 * time.Hour).Format(time.DateTime)

	// Post
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	// Update
	data2 := url.Values{}
	data2.Set("id", "999999999999")
	data2.Set("uid", "1")
	data2.Set("desc", "task1")
	data2.Set("date", n)
	enc2 := data2.Encode()

	req2 := httptest.NewRequest("POST", "/update_event", strings.NewReader(enc2))
	rec2 := httptest.NewRecorder()

	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(data2.Encode())))

	h2 := http.HandlerFunc(UpdateEvent)
	h2.ServeHTTP(rec2, req2)

	if status := rec2.Code; status != http.StatusBadRequest {
		t.Fatalf("expected status: %d, got %d", http.StatusBadRequest, status)
	}
}

func TestServerDelete(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(-1 * time.Hour).Format(time.DateTime)

	// Post
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	// Delete
	data2 := url.Values{}
	data2.Set("id", "0")
	enc2 := data2.Encode()

	req2 := httptest.NewRequest("POST", "/delete_event", strings.NewReader(enc2))
	rec2 := httptest.NewRecorder()

	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(data2.Encode())))

	h2 := http.HandlerFunc(DeleteEvent)
	h2.ServeHTTP(rec2, req2)

	if status := rec2.Code; status != http.StatusOK {
		t.Fatalf("expected status: %d, got %d", http.StatusOK, status)
	}

	if rec2.Body.String() != "{\"result\": \"success\"}" {
		t.Errorf("expected result: %s, got %s", "{\"result\": \"success\"}", rec2.Body.String())
	}
}

func TestServerDeleteInvalid(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(-1 * time.Hour).Format(time.DateTime)

	// Post
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	// Delete
	data2 := url.Values{}
	data2.Set("id", "99999999999999")
	enc2 := data2.Encode()

	req2 := httptest.NewRequest("POST", "/delete_event", strings.NewReader(enc2))
	rec2 := httptest.NewRecorder()

	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(data2.Encode())))

	h2 := http.HandlerFunc(DeleteEvent)
	h2.ServeHTTP(rec2, req2)

	if status := rec2.Code; status != http.StatusBadRequest {
		t.Fatalf("expected status: %d, got %d", http.StatusBadRequest, status)
	}
}

func TestServerGet(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(1 * time.Hour).Format(time.DateTime)

	// Post
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	// Get
	req2 := httptest.NewRequest("GET", "/events_for_day?uid=1", nil)
	rec2 := httptest.NewRecorder()

	h2 := http.HandlerFunc(EventsForDay)
	h2.ServeHTTP(rec2, req2)

	if status := rec2.Code; status != http.StatusOK {
		t.Fatalf("expected status: %d, got %d", http.StatusOK, status)
	}

	if !strings.Contains(rec2.Body.String(), "{\"result\":{") {
		t.Errorf("expected result, got %s", rec2.Body.String())
	}
}

func TestServerGetInvalid(t *testing.T) {
	storage = ReturnStorage()

	n := time.Now().Add(1 * time.Hour).Format(time.DateTime)

	// Post
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("desc", "task1")
	data.Set("date", n)
	enc := data.Encode()

	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(enc))
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	h := http.HandlerFunc(CreateEvent)
	h.ServeHTTP(rec, req)

	// Get
	req2 := httptest.NewRequest("GET", "/events_for_day?uid=99999", nil)
	rec2 := httptest.NewRecorder()

	h2 := http.HandlerFunc(EventsForDay)
	h2.ServeHTTP(rec2, req2)

	if status := rec2.Code; status != http.StatusBadRequest {
		t.Fatalf("expected status: %d, got %d", http.StatusBadRequest, status)
	}
}
