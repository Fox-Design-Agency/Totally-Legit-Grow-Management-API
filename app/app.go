// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"totally-legit-grow-management/v1/app/routes"
	"totally-legit-grow-management/v1/internal/server"
	"totally-legit-grow-management/v1/resources/config"

	"github.com/gorilla/mux"
)

func Run(ctx context.Context, cfg *config.Config, svr *server.Server) error {
	/**********************************************************************
	/
	/	Initialize router and middlesware
	/
	/**********************************************************************/

	r := mux.NewRouter()
	// register routes
	if err := routes.RegisterRoutes(r, svr); err != nil {
		return err
	}
	/**********************************************************************
	/
	/	Start Server
	/
	/**********************************************************************/
	httpServer := &http.Server{
		Addr:        ":" + cfg.Port,
		Handler:     r,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	log.Println("Running local on port: ", cfg.Port)
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// run the stuff
	return Shutdown(httpServer)
}

func Shutdown(httpServer *http.Server) error {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return err
	} else {
		log.Printf("gracefully stopped\n")
	}
	return nil
}
