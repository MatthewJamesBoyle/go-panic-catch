package goCatch

import (
	"fmt"
	"github.com/matthewJamesBoyle/go-panic-catch/catchers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatchPanicMiddleware(t *testing.T) {
	t.Run("Calls next successfully with no panic", func(t *testing.T) {

		fn := func(writer http.ResponseWriter, req *http.Request) {
			writer.WriteHeader(http.StatusTeapot)
			writer.Write([]byte("some-string"))
		}

		req := httptest.NewRequest("GET", "/aPath", nil)
		w := httptest.NewRecorder()
		log := catchers.Log{}
		PanicMiddleware(log, "", http.HandlerFunc(fn)).ServeHTTP(w, req)
		if w.Code != http.StatusTeapot {
			t.Fatal(fmt.Sprintf("Expected %d, but got %d", http.StatusTeapot, w.Code))
		}
		if w.Body.String() != "some-string" {
			t.Fatal(fmt.Sprintf("Expected %s, but got %s", "some-string", w.Body.String()))
		}
	})

	t.Run("Catches panic if next panics", func(t *testing.T) {

		fn := func(writer http.ResponseWriter, req *http.Request) {
			panic("ut oh")
		}

		req := httptest.NewRequest("GET", "/aPath", nil)
		w := httptest.NewRecorder()
		log := catchers.Log{}
		PanicMiddleware(log, "", http.HandlerFunc(fn)).ServeHTTP(w, req)

	})
	// Todo: add a webhook url below and run this test to watch the panic appear in slack :)
	//t.Run("test slack handler", func(t *testing.T) {
	//	fn := func(writer http.ResponseWriter, req *http.Request) {
	//		panic("ut oh")
	//	}
	//
	//	req := httptest.NewRequest("GET", "/aPath", nil)
	//	w := httptest.NewRecorder()
	//	slack := catchers.NewSlack("")
	//	PanicMiddleware(*slack, "you just panicked!", http.HandlerFunc(fn)).ServeHTTP(w, req)
	//
	//})
}
