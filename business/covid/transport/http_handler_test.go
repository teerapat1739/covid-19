package transport_test

import (
	"covid-19/pkg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCovidSummary(t *testing.T) {
	r := pkg.InintService()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}
