package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/shizhongwang/myswagger/chatroom-api/data"
	"github.com/shizhongwang/myswagger/chatroom-api/handlers"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := hclog.Default()

	//conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer conn.Close()

	// create database instance
	db := data.NewChatroomsDB(l)

	// create the handlers
	ph := handlers.NewChatrooms(l, db)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/chatrooms", ph.ListAll).Queries("currency", "{[A-Z]{3}}")
	getR.HandleFunc("/chatrooms", ph.ListAll)
	//
	//getR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.ListSingle).Queries("currency", "{[A-Z]{3}}")
	//getR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.ListSingle)
	//
	//putR := sm.Methods(http.MethodPut).Subrouter()
	//putR.HandleFunc("/chatrooms", ph.Update)
	//putR.Use(ph.MiddlewareValidatechatroom)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/chatrooms", ph.Create)
	//postR.Use(ph.MiddlewareValidatechatroom)

	//deleteR := sm.Methods(http.MethodDelete).Subrouter()
	//deleteR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,                                     // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
