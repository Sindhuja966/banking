package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()
	if router == nil {
		t.Error("No router is connected")
	}

	t.Run("CustomerRoutertesting", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/customers", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

	})
}
