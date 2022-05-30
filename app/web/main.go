package main

import (
	"context"
	"flag"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/app/web/routes"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/buildinfo"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/config"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/db"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/httprouter"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	logger.Init()
	buildinfo.Init()

	config.PrintFlags()

	db.Connect()
	db.Migrate()

	defer func() {
		logger.Flush()
		db.Close()
	}()

	router := httprouter.NewRouter()
	routes.Register(router)

	srv := &http.Server{
		Addr:    httprouter.GetBindAddr(),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.DefaultLogger.Fatalw("failed to listen", "error", err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server gracefully stopped")
}
