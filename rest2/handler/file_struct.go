package handler

//FileInfo is file info
import "time"

type FileInfo struct {
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	FullPath  string    `json:"fullPath"`
	Size      int       `json:"size"`
	Modified  time.Time `json:"modfied"`
	IsDir     bool      `json:"isDir"`
}

//FileListInfo is list of fileinfo
type FileListInfo struct {
	DirectoryPath string     `json:"directroyPath"`
	Files         []FileInfo `json:"files"`
}
