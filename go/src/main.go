package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func MainRequestHandler(w http.ResponseWriter, r *http.Request) {
	x := 0.0001

	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	fmt.Fprintf(w, "%s", "Code.education Rocks!")

}

func main() {
	http.HandleFunc("/", MainRequestHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
