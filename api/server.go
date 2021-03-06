package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	flagPort  = flag.Int("port", 5000, "Port")
	flagDebug = flag.Bool("debug", false, "Run in debug mode")
)

func init() {
	flag.Parse()
	initRouter()
}

func Serve() {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", *flagPort),
		Handler: ROUTER,
	}

	// start server
	log.Println("Serving at ", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	// shut down gracefully, 5 second expiry
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	server.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}
