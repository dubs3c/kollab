package main

import (
	"net/http"
)

func (a *Api) MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(*a.DefaultPaths) != 0 {
			//ctx := context.WithValue(r.Context(), "user", "123")
			// TODO - create better lookup table. hash table should work
			for _, response := range *a.DefaultPaths {
				if response.Path == r.RequestURI {
					if response.Verb == r.Method {
						if len(response.Headers) != 0 {
							for key, value := range response.Headers {
								w.Header().Add(key, value)
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
