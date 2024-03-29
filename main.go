package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

/*
//go:embed all:frontend/build
var assets embed.FS

	func Assets() (fs.FS, error) {
		return fs.Sub(assets, "frontend/build")
	}
*/
func main() {

	sqlite, err := sql.Open("sqlite3", "file:the.db?_loc=auto")

	if err != nil {
		log.Fatal(err)
	}

	defer sqlite.Close()

	_, err = sqlite.Exec(SQL_INIT)

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	s := &Api{
		DB:           sqlite,
		Servers:      &[]*Server{},
		DefaultPaths: &[]*PathResponse{},
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(s.MyMiddleware)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//assets, _ := Assets()
	//fs := http.FileServer(http.FS(assets))

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Post("/http", s.AddHttpServerHandler)
		r.Get("/http", s.GetAllHttpServer)
		r.Get("/http/{serverId}", s.GetHttpServer)
		r.Delete("/http/{serverId}", s.DeleteHttpServer)
		r.Put("/http/{serverId}", s.PutHttpServer)

		r.Post("/defaulthttp", s.AddPath)
		r.Get("/defaulthttp", s.GetAllPaths)
		r.Get("/defaulthttp/{pathId}", s.GetPath)
		r.Put("/defaulthttp/{pathId}", s.UpdatePath)
		r.Delete("/defaulthttp/{pathId}", s.DeletePath)

		r.Post("/tcp", s.AddTcpServer)

		r.Get("/events", s.GetEvents)
		r.Get("/events/{pathId}", s.GetEvents)

	})

	defaultPort := 8080
	addr := fmt.Sprintf(":%d", defaultPort)

	h := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv := &Server{}
	srv, err = GetServerById(s.DB, 0)

	if err != nil {
		log.Fatalln("something went wrong fetching default server", err)
	}

	// If there is no default server, assumed to be ID 0 in the database,
	// Create a new default http server
	if *srv == (Server{}) {
		srv.Id = uuid.UUID{}
		srv.Name = "Default HTTP Server"
		srv.Type = "HTTP"
		srv.Port = defaultPort

		if err := InsertServer(s.DB, srv); err != nil {
			log.Println("could not insert default http server, error:", err)
		}
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
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
