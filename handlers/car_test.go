package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDummyApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/dummy", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DummyApi)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `get car status.`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}

func TestGetCarStatus_badCarId(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/car/wd123123213", nil)
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("carId", "wd123123213")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCarStatus)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected :=  "\"exveErrorMsg\": \"Not Found\""

	if !strings.Contains(recorder.Body.String(), expected)  {
		t.Errorf("response body contained wrong response, got: %v , but expected was: %v", recorder.Body.String() , expected)
	}
}

func TestGetCarStatus_validCarId(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/car/WDB111111ZZZ22222", nil)
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("carId", "WDB111111ZZZ22222")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCarStatus)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "doorstatusfrontleft"

	if !strings.Contains(recorder.Body.String(), expected) {
		t.Errorf("response body contained wrong response, got: %v , but expected was: %v", recorder.Body.String() , expected)
	}
}

