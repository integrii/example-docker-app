package main

import (
  "fmt"
  "net/http"
  "os"
  "log"
)


func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  	fmt.Fprintf(w, "Version: "+os.Getenv("VERSION"))
    log.Println(r.RemoteAddr,"-",r.URL.String(),"-",r.Referer(),"-",r.UserAgent())
  })

  log.Fatal(http.ListenAndServe(":80", nil))
}
