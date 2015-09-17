package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/k0kubun/pp"
)

const assetDir = "./assets/"
const thumbnailDir = assetDir + "thumbnails/"

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if buffer, err := ioutil.ReadFile(assetDir + "index.html"); err == nil {
		fmt.Fprint(w, string(buffer))
	} else {
		pp.Print(err)
	}
	// pp.Print("inter root Handler")
}

func main() {
	// http.HandleFunc("/hello", helloHandler)     // ハンドラを登録してウェブページを表示させる
	// http.HandleFunc("/goodbye", goodbyeHandler) // ハンドラを登録してウェブページを表示させる
	http.Handle("/access/", http.StripPrefix("/access/", http.FileServer(http.Dir("/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetDir))))
	http.Handle("/thumbnails/", http.StripPrefix("/thumbnails/", http.FileServer(http.Dir(thumbnailDir))))
	http.HandleFunc("/api/", apiHandler) // ハンドラを登録してウェブページを表示させる

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":3001", nil)
}
