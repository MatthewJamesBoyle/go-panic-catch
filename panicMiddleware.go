package goCatch

import (
	"fmt"
	"net/http"
)

type PanicHandler interface {
	HandlePanic(message string) error
}

func PanicMiddleware(ph PanicHandler, message string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				err := ph.HandlePanic(message)
				if err != nil {
					fmt.Printf("panic handler fails to recover from the panic.")
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
