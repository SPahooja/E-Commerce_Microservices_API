package main

import (
	"GoMicroservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)
	//gh := handlers.Newbye(l)

	sm := mux.NewRouter()
	getrouter := sm.Methods(http.MethodGet).Subrouter()
	getrouter.HandleFunc("/", ph.GetProducts)

	putrouter := sm.Methods(http.MethodPut).Subrouter()
	putrouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putrouter.Use(ph.MiddlewareProductValidation)
	
	postrouter := sm.Methods(http.MethodPost).Subrouter()
	postrouter.HandleFunc("/", ph.AddProduct)
	postrouter.Use(ph.MiddlewareProductValidation)

	// sm.Handle("/", ph)
	//sm.Handle("/bye", gh)

	ser := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := ser.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	_ = <-sigchan
	l.Println("Received terminate, I'm leaving bro")

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	ser.Shutdown(tc)
}
