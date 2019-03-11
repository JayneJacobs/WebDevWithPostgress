package controller

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginExecutesCorrectTemmplate(t *testing.T) {
	h := new(home)
	expected := "login template"
	h.loginTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()

	h.handleLogin(w, r)
	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected /* +"foo"*/ {
		t.Error("Failed execute correct template")
	}
}
