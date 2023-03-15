package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Println("/hello endpoint called")
	hostname, err := os.Hostname()
	if err != nil {
		req.Response.StatusCode = 500
		req.Response.Status = err.Error()
		return
	}
    fmt.Fprintf(w, "hello my computer name is %s\n", hostname)
}

func main() {
    http.HandleFunc("/hello", hello)
    fmt.Println("Server up and listening...")
    err := http.ListenAndServe(":1080", nil)
	fmt.Println(err)
}