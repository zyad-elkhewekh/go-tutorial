package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/handlers"

	"github.com/go-chi/chi"
)

// handlers/handlers_test.go
func TestGetChoices(t *testing.T) {
	req := httptest.NewRequest("GET", "/choice?username=zyad", nil)
	req.Header.Set("Authorization", "123abc")
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	handlers.Handler(r)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", w.Code, w.Body.String())
	}
}
