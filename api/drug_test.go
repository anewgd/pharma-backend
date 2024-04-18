package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDrug(t *testing.T) {

	drugHandler := &DrugHandler{}
	testServer := NewServer(*drugHandler)
	req, _ := http.NewRequest(http.MethodPost, "/drugs", nil)
	res := httptest.NewRecorder()
	testServer.router.ServeHTTP(res, req)
	got := res.Body.String()
	want := "test drug"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}