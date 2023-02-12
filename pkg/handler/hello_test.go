package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
)

func TestHelloMessage(t *testing.T) {
	t.Run("failed param", func(t *testing.T) {
		e := echo.New()
		q := make(url.Values)
		q.Set("id", "")
		request, _ := http.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		response := httptest.NewRecorder()

		c := e.NewContext(request, response)

		HelloMessage(c)

		assertQueryParam(t, response.Code, http.StatusBadRequest)
	})
	t.Run("empty param", func(t *testing.T) {
		e := echo.New()
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		c := e.NewContext(request, response)

		HelloMessage(c)

		assertQueryParam(t, response.Code, http.StatusBadRequest)
	})
	t.Run("valid param", func(t *testing.T) {
		e := echo.New()
		q := make(url.Values)
		q.Set("id", "test")
		request, _ := http.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		response := httptest.NewRecorder()

		c := e.NewContext(request, response)

		HelloMessage(c)

		assertQueryParam(t, response.Code, http.StatusOK)
	})
}

func assertQueryParam(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("handler returned wrong status code: got %v want %v",
			got, want)
	}
}
