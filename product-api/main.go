package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"sbrockmann.com/product-api/handlers"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	ph := handlers.NewProducts(l)

	// sm := http.NewServeMux()
	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// sm.Handle("/products", ph)


	s := &http.Server{
		Addr: ":9006",
		Handler: sm,
		IdleTimeout: 120 *time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	
	tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)
	s.Shutdown(tc)
}