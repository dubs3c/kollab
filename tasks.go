package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateHttpServer(server *Server) {

	go func() {
		mux := http.NewServeMux()

		defaultHandler := func(w http.ResponseWriter, r *http.Request) {
			var h []string = []string{}
			for _, response := range server.HttpServer.Responses {
				if response.Path == r.RequestURI {
					if r.Method == response.Verb {
						if len(response.Headers) != 0 {
							for _, header := range response.Headers {
								h = strings.Split(header, ":")
								w.Header().Add(h[0], h[1])
							}
						}
						w.Write(response.Body)
					}
				}
			}
		}

		s := &http.Server{
			Addr:           ":" + strconv.Itoa(server.HttpServer.Port),
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		for _, resp := range server.HttpServer.Responses {
			mux.HandleFunc(resp.Path, defaultHandler)
		}

		go func(s *http.Server) {
			<-server.Shutdown
			log.Printf("[%s] Received shutdown signal....", server.Id.String())
			if err := s.Shutdown(context.Background()); err != nil {
				log.Printf("[%s] HTTP server Shutdown: %v", server.Id.String(), err)
			}
			close(server.Shutdown)
		}(s)

		s.ListenAndServe()

	}()

}
