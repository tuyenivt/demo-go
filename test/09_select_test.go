package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Microsecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	expected := fastUrl
	actual := Racer(slowUrl, fastUrl)

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}

	slowServer.Close()
	fastServer.Close()
}
