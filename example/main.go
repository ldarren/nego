package main

import (
    "fmt"
    "net/http"
    "log"
	"github.com/ldarren/nego"
)

func Index(w http.ResponseWriter, r *http.Request, _ nego.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps nego.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.Get("name"))
}

func main() {
    router := nego.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
