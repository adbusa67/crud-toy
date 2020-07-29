package server

import (
	"crud-toy/internal/app/service/infra/config"
	"github.com/urfave/negroni"
	"crud-toy/internal/app/service/infra/logger"
	"os/signal"
	"os"
	"syscall"
	"net/http"
	"context"

)

func Start() error {
	
	appPort := ":" +config.Config().AppPort
	server := negroni.New(negroni.NewRecovery())
	router, err := NewRouter()
	if err != nil {
		return err
	}
	server.UseHandler(router)
	logger.Info(`Starting server on port 3000`)
	httpServer := &http.Server{
		Addr:    appPort,
		Handler: server,
	}

	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		if shutdownErr := httpServer.Shutdown(context.Background()); shutdownErr != nil {
			logger.Error("Received an Interrupt Signal", shutdownErr)
		}
	}()

	if err = httpServer.ListenAndServe(); err != nil {
		logger.Error("HTTP Server Failed ", err)
	}
	logger.Info("Stopped server gracefully")
	return nil
}