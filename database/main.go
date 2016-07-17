package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/m0a/go_sandbox/database/models"
	"time"
	"net/http"
	"fmt"
	"encoding/json"

	"encoding/base64"
	"strings"
)

func feedbackHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "POST" {
		fmt.Fprintf(w, "0")
		return	
	}
	
	m := make(map[string]interface{})
	
	if err := json.Unmarshal([]byte(r.PostFormValue("feedback")), &m); err != nil {
		fmt.Fprintf(w, "0")
		panic(err.Error())
	}
	//  fmt.Println(r.PostFormValue("feedback"))
	fmt.Println(m)
	db ,err := sql.Open("mysql", "root:mysql@tcp(docker:3306)/feedbacks")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	feedback := models.Feedback{}
	feedback.FeedbacksStatusID = 1
	img, ok := m["img"].(string)
	if !ok {
		panic("img err")
	}
	b64data := img[strings.Index(img,",") + 1:]
	decoded, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		panic(err.Error())
	}
	
	feedback.Img = decoded 
	feedback.HTML = m["html"].(string)
	feedback.Note = m["note"].(string)
	feedback.Created = new(time.Time)
	t := time.Now()
	feedback.Created = &t 
	feedback.Modified = &t
	if err := feedback.Insert(db); err != nil {
		panic(err.Error())
	}

	fmt.Println("OK. recive")
	fmt.Fprintf(w, "1")


}
func main()  {

	http.HandleFunc("/feedback/send", feedbackHandler)
	http.ListenAndServe(":9999", nil)
}