package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog" //Eklenecek.
	"github.com/rs/zerolog/log"

	"github.com/samcho07/getir-test-case-MS/data"
	"github.com/samcho07/getir-test-case-MS/data/store"
)

/*
func Init() {
	log.Logger = logger
}

func Error(msg error) {
	log.Logger.Error().Err(msg)
}

func Info(msg string) {
	log.Logger.Info().Msg(msg)
}
*/

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func StatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Allrgiht..."))
}

func main() {
	// Connect to MongoDB
	MongoDB_Server = new(search.mongodb)
	log.Logger = Logger

	Hold := store.NewCacheProvider()
	dataHand := data.New(Hold)

	http.HandleFunc("/in-memory", dataHandler.GetInMemory)
	http.HandleFunc("/in-memory/", dataHandler.SetInMemory)
	http.HandleFunc("/records", mongoServer.SearchMongo)
	http.HandleFunc("/", StatusCheck)

	// connectting the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpServer := &http.Server{
		Addr: ":" + port,
	}

	log.Logger.Info().Msg("App has been started at 8080...")

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// making an channel and go routine..
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Print("Interrupt closed..\n")

	go func() {
		<-signalChan
		log.Fatal("Kill terminate..\n")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("error handled: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("stopped\n")
	}

	defer os.Exit(0)

}
