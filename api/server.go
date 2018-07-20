package api

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/Sirupsen/logrus"
)

var (
	flagPort     = flag.Int("port", 5000, "Port")
	flagDBStr    = flag.String("dbstr", "dbname=elacarte sslmode=disable", "DB connection string")
	flagPoolsize = flag.Int("poolsize", 50, "DB connection pool size")
	flagDebug    = flag.Bool("debug", false, "Run in debug mode")
)

func init() {
	flag.Parse()
	initDB()
}

func Serve() {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", *flagPort),
		Handler: Router,
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
