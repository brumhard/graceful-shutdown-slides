package main

import (
	"log"
	"time"
)

func main() {
	somec := make(chan struct{})

	select { // HL
	case <-somec: // HL
	case <-time.After(time.Second): // HL
		log.Print("timeout") // HL
	} // HL

	log.Print("done")
}
