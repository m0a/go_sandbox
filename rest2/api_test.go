package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

func Test_FileRouteing(t *testing.T) {

	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/file?file=./test/aaabbccc.txt")
	if err != nil {
		t.Error("unexpected")
	}
	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("invalid status code")
	}

	if c != `{"Id":1,"Name":"hoge"}` {
		t.Errorf("invalid response [%s]", c)
	}
}

func Test_FileListRouteing(t *testing.T) {

	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/files?dir=./test/")
	if err != nil {
		t.Error("unexpected")
	}
	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("invalid status code")
	}

	if c != `{"Id":1,"Name":"hoge"}` {
		t.Errorf("invalid response [%s]", c)
	}
}
