package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOnlineEndpoint(t *testing.T) {
	router := setUpServer()
	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, status)
	}
}

func TestPlantingEndpoint(t *testing.T) {
	router := setUpServer()
	req, _ := http.NewRequest("GET", "/planting-periods", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, status)
	}
}

func TestSeasonEndpoint(t *testing.T) {
	router := setUpServer()
	req, _ := http.NewRequest("GET", "/in-season", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, status)
	}
}

// func TestBenchmark(b *testing.B) {
// 	for i := 0; i < b.N; i++ {

// 	}
// }
