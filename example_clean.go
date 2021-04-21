package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalc := make(chan os.Signal, 1)
	defer signal.Stop(signalc)

	signal.Notify(signalc, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signalc: // first signal, cancel context // HL
			cancel()
		case <-ctx.Done():
		}
		<-signalc // second signal, hard exit // HL
		log.Print("shutdown forcefully")
		os.Exit(exitCodeInterrupt)
	}()

	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context) error {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Print("doing sth while waiting for ctx cancel")
		<-ctx.Done() // blocking until ctx cancel // HL

		log.Print("shutting down")
		time.Sleep(5 * time.Second) // work to do before shutdown // HL

		log.Print("graceful shutdown done")
	}()

	<-ctx.Done() // waiting for context cancel from main // HL
	wg.Wait()

	return nil
}
