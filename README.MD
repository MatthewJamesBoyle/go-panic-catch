# go-panic-catch
[![GoDoc](https://godoc.org/github.com/MatthewJamesBoyle/go-panic-catch?status.svg)](https://godoc.org/github.com/MatthewJamesBoyle/go-panic-catch)

go-panic-catch is a simple middleware for golang web servers that recovers from panics and enables you to handle panics gracefully.

To get it:

```go get github.com/matthewjamesboyle/go-panic-catch```

And then to use it:
```
package main

import (
	"github.com/matthewjamesboyle/go-panic-catch"
	"github.com/matthewjamesboyle/go-panic-catch/catchers"
	"net/http"
)

func main() {
	fn := func(writer http.ResponseWriter, req *http.Request) {
		panic("ut oh")
	}

	//something that satisfies panicHandler interface
	slack := catchers.NewSlack("some-webhook-url")

	server := &http.Server{
		Handler: goCatch.PanicMiddleware(*slack, "recovering from panic will write this to slack", http.HandlerFunc(fn)),
	}

	server.ListenAndServe()
}

```

It should also work with all the popular web servers such as gorillia mux and gin.

go-panic-catch currently comes with two "catchers" that handle panics.
* slack catcher - will write a message to slack everytime your web server panics
* log - Logs "message" every time there is a panic.

Want to implement your own handler? Simply implement the following interface:
```
type PanicHandler interface {
	HandlePanic(message string) error
}
```

Some ideas for other handlers:
* email everytime there is a panic.
* increment a prometheus counter.

(contributions appreciated if you want to work on these)
