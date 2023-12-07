package WebTest

import (
	"fmt"
	"net/http"
	"testing"
)

type HelloWorldHandler struct {
	num int
}

func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.num += 1
	w.Write([]byte(fmt.Sprintf("Hello World: %d", h.num)))
}

func TestHttpServer(t *testing.T) {
	http.Handle("/", &HelloWorldHandler{})
	http.ListenAndServe(":8080", nil)
}
