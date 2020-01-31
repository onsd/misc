package main

import(
    "fmt"
    "net/http"
)

func main() {
    // register function
    // Hello
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello NITTC!")
}

func dangerousHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "THIS IS DANGEROUSE")
}
