package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
	// CTX START OMIT
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// CTX END OMIT

	// SIG START OMIT
	signalc := make(chan os.Signal, 1)
	defer signal.Stop(signalc)

	signal.Notify(signalc, syscall.SIGINT, syscall.SIGTERM)
	// SIG END OMIT

	// HANDLE START OMIT
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
	// HANDLE END OMIT

	// RUN START OMIT
	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
	// RUN END OMIT
}

func run(ctx context.Context) error {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Print("graceful shutdown done")

		log.Print("doing sth while waiting for ctx cancel")
		<-ctx.Done() // blocking until ctx cancel // HL

		log.Print("shutting down")
		time.Sleep(5 * time.Second) // work to do before shutdown // HL
	}()

	<-ctx.Done() // waiting for context cancel from main // HL
	wg.Wait()

	return nil
}

func http_run(ctx context.Context) error {
	defer log.Print("api shutdown done")

	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("hello there\n"))
		}),
	}

	go func() {
		defer log.Print("stopped listening to requests")
		log.Printf("accepting requests at %s ", server.Addr)
		server.ListenAndServe() // HL
	}()

	<-ctx.Done() // HL

	log.Printf("shutting down server gracefully")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	return server.Shutdown(shutdownCtx) // HL
}
