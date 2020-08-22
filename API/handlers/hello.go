package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Got an error in reading request", http.StatusBadRequest)
		return
	}
	h.l.Printf("Bro %s is here\n", body)
	fmt.Fprintf(rw, "What's up %s", body)
}
