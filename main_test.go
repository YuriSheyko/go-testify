package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequireIsTrue(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := responseRecorder.Body.String()

	// здесь нужно добавить необходимые проверки

	assert.Greater(t, len(body), 0)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
}
func TestMainHandlerWhenNoSuchCity(t *testing.T) {

	expected := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=4&city=taganrog", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//if status := responseRecorder.Code; status != http.StatusOK {
	//t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	//}

	body := responseRecorder.Body.String()

	assert.Equal(t, expected, body)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	expectation := 4
	req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := strings.Split(responseRecorder.Body.String(), ",")

	// здесь нужно добавить необходимые проверки

	require.Len(t, body, expectation)
}
