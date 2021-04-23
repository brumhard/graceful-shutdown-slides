package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for {
			log.Print("waiting")

			select { // HL
			case <-time.After(time.Second): // HL
			} // HL
		}
	})

	log.Printf("accepting requests at :8080")
	http.ListenAndServe(":8080", nil)
}
