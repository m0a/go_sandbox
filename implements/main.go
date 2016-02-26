package main 
import (
    "fmt"
    "net/http"
    "time"
)

// Myhandle test handler
type Myhandle struct{}
// MyHandleWrapper handle Wrapper
type MyHandleWrapper struct{
    Myhandle
    name string
}

var _ http.Handler = &Myhandle{}
var _ http.Handler = &MyHandleWrapper{}

func (h *Myhandle)ServeHTTP(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"Hello")
}

func main() {

    fmt.Println("hello")
    time.Sleep(time.Second * 100)
    
    var h MyHandleWrapper    
     http.ListenAndServe("localhost:4999",&h)
    
}