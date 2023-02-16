package main

import (
	"net/http"
	"strings"
)

func (a *Api) MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(*a.DefaultPaths) != 0 {
			var h []string = []string{}
			//ctx := context.WithValue(r.Context(), "user", "123")
			// TODO - create better lookup table. hash table should work
			for _, response := range *a.DefaultPaths {
				if response.Path == r.RequestURI {
					if response.Verb == r.Method {
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
