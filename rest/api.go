package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

// parse url string
func pathTokens(r *http.Request) []string {
	var tokens = make([]string, 0)

	path := html.EscapeString(r.URL.Path)
	tempTokens := strings.Split(path, "/")

	for _, v := range tempTokens {
		// pathTokens
		if v != "" {
			tokens = append(tokens, v)
		}
	}

	return tokens
}

type errorMessage struct {
	Message string `json:"message"`
}

func pop(slice []string) (string, []string) {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

func shift(slice []string) (string, []string) {
	ans := slice[0]
	slice = slice[1:]
	return ans, slice
}

func apiError(w http.ResponseWriter, mes string) {
	message := errorMessage{Message: mes}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Cache-Control", "max-age=0, private, must-revalidate")
	w.WriteHeader(http.StatusBadRequest)
	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

//FileInfo is file info
type FileInfo struct {
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	Size      int       `json:"size"`
	Modified  time.Time `json:"modfied"`
	IsDir     bool      `json:"isDir"`
}

//FileListInfo is list of fileinfo
type FileListInfo struct {
	Path        string     `json:"path"`
	PublishedAt time.Time  `json:"published_at"`
	Files       []FileInfo `json:"files"`
}

// func NewFileInfo(files []os.FileInfo) []FileInfo {
// 	fileinfoList := make([]FileInfo, 0)
//
// 	return fileinfoList
// }

func fileList(dir string) ([]FileListInfo, error) {
	if dir == "" {
		dir = "./"
	}
	pp.Printf("run filelist dir = [%s]\n", dir)

	list, err := ioutil.ReadDir(dir)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "error: %v", err)
		// os.Exit(1)
		return nil, fmt.Errorf("not found")
	}
	// pp.Print(list)
	// fmt.Fprintf(w, "%#v\n", list)

	var files = make([]FileInfo, 0)
	for _, finfo := range list {

		fileinfo := FileInfo{Name: finfo.Name()}
		fileinfo.Name = finfo.Name()
		fileinfo.Size = int(finfo.Size())
		fileinfo.IsDir = finfo.IsDir()
		fileinfo.Modified = finfo.ModTime()
		fileinfo.Extension = filepath.Ext(fileinfo.Name)

		// realPath := "./" + finfo.Name()
		// pp.Println(realPath)

		if fileinfo.Extension == ".mp4" || fileinfo.IsDir {
			files = append(files, fileinfo)
		} else {
			pp.Println(fileinfo)
		}

	}
	fileListInfo := []FileListInfo{FileListInfo{}}

	if path, err := filepath.Abs(dir); err == nil {
		fileListInfo[0].Path = path
	} else {
		fileListInfo[0].Path = dir
	}

	fileListInfo[0].Files = files

	pp.Println(fileListInfo)
	return fileListInfo, nil

}

const defualtAPIVersion = 1

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// req := Request(*r)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	// w.Header().Set("Access-Control-Max-Age", "3600")
	// w.Header().Set("Access-Control-Allow-Headers", "x-requested-with")

	pathTokens := pathTokens(r)
	var s string
	if s, pathTokens = shift(pathTokens); s != "api" {
		apiError(w, "not found api")
		return
	}

	apiVersion := defualtAPIVersion
	_ = apiVersion
	apiName := ""

	if s, pathTokens = shift(pathTokens); strings.HasPrefix(strings.ToLower(s), "v") {
		apiVersion, _ = strconv.Atoi(s[1:])
	} else {
		apiName = s
	}

	if apiName == "" {
		apiName = pathTokens[0]
	}

	// fmt.Printf("apiVersion -> %d\n", apiVersion)
	// fmt.Printf("apiName -> %s\n", apiName)
	// http.ParseQuery
	// fmt.Printf("parameters -> %+v", r.URL.Query())
	parameters := r.URL.Query()
	fmt.Printf("parameters -> %+v", parameters)

	switch apiName {
	case "files":
		var targetPath string
		if val, ok := parameters["dir"]; ok {
			targetPath = val[0]
		} else {
			targetPath = "./"
		}

		fmt.Println(targetPath)
		dirs, err := fileList(targetPath)

		if err != nil {
			// fmt.Println(dirs)
			return
		}

		js, err := json.Marshal(dirs)
		if err != nil {
			apiError(w, "not found dir")
			return
		}
		w.Write(js)
	}

}
