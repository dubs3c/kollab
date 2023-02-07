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
					// TODO - include HTTP Header data as well
					w.Write(response.Body)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
