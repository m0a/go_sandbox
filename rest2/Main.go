package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	// log.Fatal(http.ListenAndServe(":3001", router))
	rootpath, err := filepath.Abs("./")
	if err != nil {
		fmt.Printf("rootpath err:%v", err)
		return
	}

	fmt.Printf("rootpath:[%s]\n", rootpath)

	router := NewRouter(rootpath)

	http.Handle("/api/", router)
	http.Handle("/access/", http.StripPrefix("/access/", http.FileServer(http.Dir(rootpath))))
	http.ListenAndServe(":3001", nil)
}
