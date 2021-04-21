package main

import (
	"log"
	"time"
)

func main() {
	somec := make(chan struct{})

	go func() {
		somec <- struct{}{}
	}()

	select {
	case <-somec:
	case <-time.After(2 * time.Second):
		log.Print("timeout")
	}

	log.Print("done")
}
