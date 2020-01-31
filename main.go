package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var COUNT = 10

func main() {
	// register function
	http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/dangerous", dangerousHandler)
	http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello SecHack365!</h1>")
}

func dangerousHandler(w http.ResponseWriter, r *http.Request) {
	COUNT = COUNT - 1
	s := strconv.Itoa(COUNT)
	if s == "0" {
		// fmt.Fprintf(w, "DIED")
		os.Exit(1)
	}

	fmt.Fprintf(w, "<h1>"+s+" TIMES UNTIL DIE</h1>")
}
