package main

import (
	"log"
	"time"
)

func main() {
	somec := make(chan struct{})
	// OMIT
	go func() { // OMIT
		somec <- struct{}{} // OMIT
	}() // OMIT

	select { // HL
	case <-somec: // HL
	case <-time.After(time.Second): // HL
		log.Print("timeout") // HL
	} // HL

	log.Print("done")
}
