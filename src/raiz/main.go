package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Raiz: %.1f. %s", loopRaiz(), "Code.education Rocks!")
}

func loopRaiz() float64 {
	val := 0.0001
	for i := 0; i < 1000000; i++ {
		val = math.Sqrt(val)
	}
	return val
}
