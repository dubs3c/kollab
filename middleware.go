package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (a *Api) MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		paths, err := GetAllPaths(a.DB)
		if err != nil {
			log.Println("Middleware:", err)
			RespondWithError(w, 500, "Something went wrong processing your request")
			return
		}

		if len(*paths) != 0 {
			var h []string = []string{}
			var event EventPath = EventPath{}
			//ctx := context.WithValue(r.Context(), "user", "123")
			for _, response := range *paths {
				if response.Path == r.RequestURI {
					if response.Verb == r.Method {
						event = EventPath{
							Path:           response.Path,
							IP:             r.RemoteAddr,
							RequestHeaders: r.Header,
						}

						eventJson, _ := json.Marshal(event)

						if err := CreateEventLogPath(a.DB, response.Id, eventJson); err != nil {
							log.Println("Middleware:", err)
							//RespondWithError(w, 500, "Something went wrong processing your request")
							//return
						}

						if len(response.Headers) != 0 {
							for _, header := range response.Headers {
								h = strings.Split(header, ":")
								w.Header().Add(h[0], h[1])
							}
						}

						w.Write(response.Body)
						return
					}
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
