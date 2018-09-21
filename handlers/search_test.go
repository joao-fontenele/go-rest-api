package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(nil)

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{})

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	handler := Search{}
	rw := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest("POST", "/search", nil), rw, handler
	}

	data, _ := json.Marshal(d)
	return httptest.NewRequest("POST", "/search", bytes.NewReader(data)), rw, handler
}
