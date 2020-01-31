package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var COUNT = 10

func main() {
	// register function
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dangerous", fixedDangerousHandler)
	http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello SecHack365!</h1>")
}

func dangerousHandler(w http.ResponseWriter, r *http.Request) {
	COUNT = COUNT - 1
	s := strconv.Itoa(COUNT)
	if s == "0" {
		os.Exit(1)
	}

	fmt.Fprintf(w, "<h1>"+s+" TIMES UNTIL DIE</h1>")
}

func fixedDangerousHandler(w http.ResponseWriter, r *http.Request) {
	if COUNT == 0 {
		fmt.Fprintf(w, "<h1>HEALING DANGEROUS POINT</h1>")
		return
	}
	COUNT = COUNT - 1
	s := strconv.Itoa(COUNT)
	if s == "0" {
		// fmt.Fprintf(w, "DIED")
		go func() {
			time.Sleep(10 * time.Second)
			COUNT = 10
		}()
	}
	fmt.Fprintf(w, "<h1>"+s+" TIMES UNTIL DIE</h1>")
}
