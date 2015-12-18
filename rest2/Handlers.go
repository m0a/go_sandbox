package main

import (
	"fmt"
	"net/http"
)

func openFile(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// filePath := vars["filePath"]
	parameters := r.URL.Query()
	filePath := parameters["file"]
	fmt.Fprintln(w, "openFile:", filePath)
}
