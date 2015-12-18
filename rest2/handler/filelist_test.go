package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/bitly/go-simplejson"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

func Test_FileListHandlerDontHaveDirQuery(t *testing.T) {

	// router := NewRouter()
	rootpath, _ := filepath.Abs("../")
	ts := httptest.NewServer(http.HandlerFunc(FileList(rootpath)))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
	}

	c, s := ParseResponse(res)
	if s != http.StatusBadRequest {
		t.Errorf("invalid status code:%d", s)
	}

	if c != `{"message":"Not found dir query"}` {
		t.Errorf("invalid response [%s]", c)
	}
}

func Test_FileListHandler(t *testing.T) {

	// router := NewRouter()
	rootpath, _ := filepath.Abs("../")
	filelist := FileList(rootpath)
	ts := httptest.NewServer(http.HandlerFunc(filelist))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?dir=/test/")
	if err != nil {
		t.Error("unexpected")
	}

	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Errorf("invalid status code:%d", s)
	}

	json, err := simplejson.NewJson([]byte(c))
	if err != nil {
		t.Errorf("invalid c [%s]", c)
	}
	oneFile := json.GetIndex(0).Get("files").GetIndex(1)
	name, _ := oneFile.Get("name").String()
	// name := oneFile.Get("name").String()
	if name != "a.txt" {
		t.Errorf("invalid name [%s]", name)
	}

	extension, _ := oneFile.Get("extension").String()
	if extension != ".txt" {
		t.Errorf("invalid extension [%s]", extension)
	}

}
