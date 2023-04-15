package http

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/configs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitServer(config configs.Config, r *mux.Router) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", config.GracefulTimeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.WebServerPort),
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
