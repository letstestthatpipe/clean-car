package handlers

import (
	"net/http"
	"net/http/httptest"
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

func TestGetCarStatus(t *testing.T) {

}
