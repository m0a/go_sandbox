package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/k0kubun/pp"
)

const (
	errorMessageNotFoundDirQuery  = "Not found dir query"
	errorMessageInvalidDirQuery   = "Invalid dir query"
	errorMessageNotFoundFileQuery = "Not found file query"
)

// FileList file list from dir query
//
func FileList(rootPath string) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// dirPath := vars["dirPath"]
		// pp.Printf("rootPath:[%s]", rootPath)

		parameters := r.URL.Query()

		var absDir string

		if val, ok := parameters["dir"]; ok {
			absDir = path.Join(rootPath, val[0])
		} else {
			apiError(w, errorMessageNotFoundDirQuery)
			return
		}
		pp.Println("absDir", absDir)
		if !strings.HasPrefix(absDir, rootPath) {
			apiError(w, errorMessageInvalidDirQuery)
			return
		}
		// fmt.Fprintln(w, "fileList:", dirPath)
		// pp.Println(dirPath)
		// path.

		dirPath := strings.TrimPrefix(absDir, rootPath)
		dirs, err := fileListOnSys(dirPath, rootPath)

		if err != nil {
			// fmt.Println(dirs)
			return
		}

		js, err := json.Marshal(dirs)
		if err != nil {
			apiError(w, "not found dir")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Cache-Control", "max-age=0, private, must-revalidate")

		w.Write(js)
	})

}

type errorMessage struct {
	Message string `json:"message"`
}

func apiError(w http.ResponseWriter, mes string) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Cache-Control", "max-age=0, private, must-revalidate")
	w.WriteHeader(http.StatusBadRequest)
	js, err := json.Marshal(errorMessage{Message: mes})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func fileListOnSys(dir, rootPath string) ([]FileListInfo, error) {
	if dir == "" {
		dir = "/"
	}
	absDir := ""
	var err error
	absDir, err = filepath.Abs(dir)
	if err != nil {
		dir = "/"
		absDir = rootPath
	}

	if !strings.HasPrefix(absDir, rootPath) {
		// apiError(w, errorMessageInvalidDirQuery)
		dir = "/"
		absDir = rootPath
	}

	// pp.Printf("run filelist dir = [%s]\n", dir)
	list, err := ioutil.ReadDir(absDir)
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

		// if fileinfo.Extension == ".mp4" || fileinfo.IsDir {
		// 	files = append(files, fileinfo)
		// } else {
		// 	pp.Println(fileinfo)
		// }

		fileinfo.FullPath = path.Clean(path.Join(dir, finfo.Name()))
		// if err != nil {
		// 	os.Exit(1)
		// }

		files = append(files, fileinfo)
	}
	fileListInfo := []FileListInfo{FileListInfo{}}
	fileListInfo[0].DirectoryPath = dir
	fileListInfo[0].Files = files
	// pp.Println(fileListInfo)
	return fileListInfo, nil

}
