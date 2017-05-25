package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(`
		{
			"name": "foo",
			"token": "bar"
		}
	`))
	w := httptest.NewRecorder()

	RegisterUser(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status %v", resp.Status)
	}
}

func TestRegisterDevice(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(`
		{
			"name": "foo",
			"token": "t"
		}
	`))
	w := httptest.NewRecorder()

	RegisterUser(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status %v", resp.Status)
	}

	var userResponse newUserResponse
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &userResponse)

	req := httptest.NewRequest("POST", "/", strings.NewReader`
	
	`)
}
