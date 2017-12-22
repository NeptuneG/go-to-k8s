package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := Router()
	ts := httptest.NewServer(r)

	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}
	// 200
	if res.StatusCode != http.StatusOK {
		t.Errorf("Got unexpected Http Status Code:\n\tExpected: %d, Actual: %d\n",
			http.StatusOK, res.StatusCode)
	}

	res, err = http.Post(ts.URL+"/home", "test/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	// 405
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Got unexpected Http Status Code:\n\tExpected: %d, Actual: %d\n",
			http.StatusMethodNotAllowed, res.StatusCode)
	}

	res, err = http.Get(ts.URL + "/dumb-jerk")
	if err != nil {
		t.Fatal(err)
	}
	// 404
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Got unexpected Http Status Code:\n\tExpected: %d, Actual: %d\n",
			http.StatusNotFound, res.StatusCode)
	}
}

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	home(w, nil)
	res := w.Result()
	if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
		t.Errorf("Got unexpected Http Status Code:\n\tExpected: %d, Actual: %d\n",
			expected, actual)
	}

	msgBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if expected, actual := "#good luck\n", string(msgBody); expected != actual {
		t.Errorf("Got unexpected message:\n\tExpected: %s, Actual: %s\n",
			expected, actual)
	}
}
