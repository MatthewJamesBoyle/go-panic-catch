package goCatch

import (
	"fmt"
	"net/http"
)

//PanicHander is an interface. It is taken by the PanicMiddleware. All PanicHandlers are responsible for dealing with unexpected
//panics in your server.
type PanicHandler interface {
	HandlePanic(message string) error
}

//PanicMiddleware should be wrapped around all other handlers in your web server. It returns a http.handler and should be flexible enough
// to work with all popular go Web servers. If your PanicHandler fails, it will log to the console.
func PanicMiddleware(ph PanicHandler, message string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				err := ph.HandlePanic(message)
				if err != nil {
					fmt.Printf("panic handler failed to recover from the panic.")
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
