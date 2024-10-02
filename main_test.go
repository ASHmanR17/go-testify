package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestMainHandlerWhenCountMoreThanTotal Тест, если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	//проверяем, что сервис возвращает код ответа 200. Если нет, то останавливаем выполнение этого теста
	require.Equal(t, http.StatusOK, status)

	cafe := responseRecorder.Body.String()
	cafeSlice := strings.Split(cafe, ",")
	//проверяем, что в теле ответа все доступные кафе.
	assert.Equal(t, totalCount, len(cafeSlice))
}

// TestMainHandlerWhenOk Тест, что если запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	//проверяем, что сервис возвращает код ответа 200, если нет, то останавливаем выполнение этого теста
	require.Equal(t, http.StatusOK, status)
	//Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
}

// TestMainHandlerWrongCity Тест, что город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=pskov", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	//проверяем, что сервис возвращает код ответа 400
	assert.Equal(t, http.StatusBadRequest, status)
	//Проверяем, что в теле ответа "wrong city value"
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}
