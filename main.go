package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	s := &Api{
		Servers:      &[]*Server{},
		DefaultPaths: &[]*HttpResponse{},
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(s.MyMiddleware)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Post("/http", s.AddHttpServerHandler)
		r.Get("/http", s.GetAllHttpServer)
		r.Get("/http/{serverId}", s.GetHttpServer)
		r.Delete("/http/{serverId}", s.DeleteHttpServer)
		r.Put("/http/{serverId}", s.PutHttpServer)

		r.Post("/defaulthttp", s.AddDefaultHttpHandler)

		r.Post("/tcp", s.AddTcpServer)

	})

	h := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		// TODO: in addition, loop over all servers and issue a shutdown signal
		log.Println("Received interrupt signal, shutting down this operation...")
		if err := h.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	fmt.Println(`
_  _____  _    _      _   ___  ___  ___    _ _____ _  _ ___ 
| |/ / _ \| |  | |    /_\ | _ )/ _ \| _ \  /_\_   _(_)(_) _ \
| ' < (_) | |__| |__ / _ \| _ \ (_) |   / / _ \| | / __ \   /
|_|\_\___/|____|____/_/ \_\___/\___/|_|_\/_/ \_\_| \____/_|_\ 
                                         ɃɎ ᴋᴀᴩᴛᴇɴ ꜱᴠᴀʀᴛꜱᴋÄɢɢ
    `)

	log.Printf("Setting sail for %s...", h.Addr)
	if err := h.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
