package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"github.com/shizhongwang/myswagger/nitro2/handlers"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)

	router := gin.Default()
	router.POST("/chatrooms", ph.Create)
	router.PUT("/chatrooms", ph.Update)
	router.DELETE("/chatrooms/:id", ph.Delete)
	router.GET("/chatrooms", ph.ListAll)
	router.GET("/chatrooms/:id", ph.ListSingle)

//http://your-mattermost-url.com/api/v4/channels/{channel_id}/members

	chatroommemberDB := data.NewChatroomMembersDB(l)
	chatroommemberHandler := handlers.NewChatroomMembers(l,chatroommemberDB)
	router.POST("/chatroom/:chatroomid/members", chatroommemberHandler.Create)
	router.GET("/chatroom/members", chatroommemberHandler.ListAll)
	router.GET("/chatroom/:chatroomid/members", chatroommemberHandler.ListMembersByChatroomID)

	router.Run(":8080")

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	//// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	//getR.HandleFunc("/chatrooms", ph.ListAll).Queries("currency", "{[A-Z]{3}}")
	//getR.HandleFunc("/chatrooms", ph.ListAll)
	//
	//getR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.ListMembersByChatroomID).Queries("currency", "{[A-Z]{3}}")
	//getR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.ListMembersByChatroomID)
	//
	//putR := sm.Methods(http.MethodPut).Subrouter()
	//putR.HandleFunc("/chatrooms", ph.Update)
	//putR.Use(ph.MiddlewareValidatechatroom)

	//postR := sm.Methods(http.MethodPost).Subrouter()
	//postR.HandleFunc("/chatrooms", ph.Create)
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

//func main() {
//
//	env.Parse()
//
//	l := hclog.Default()
//
//	//conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//defer conn.Close()
//
//	// create database instance
//	db := data.NewChatroomsDB(l)
//
//	// create the handlers
//	ph := handlers.NewChatrooms(l, db)
//
//	// create a new serve mux and register the handlers
//	sm := mux.NewRouter()
//
//	// handlers for API
//	getR := sm.Methods(http.MethodGet).Subrouter()
//	getR.HandleFunc("/chatrooms", ph.ListAll).Queries("currency", "{[A-Z]{3}}")
//	getR.HandleFunc("/chatrooms", ph.ListAll)
//	//
//	//getR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.ListMembersByChatroomID).Queries("currency", "{[A-Z]{3}}")
//	//getR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.ListMembersByChatroomID)
//	//
//	//putR := sm.Methods(http.MethodPut).Subrouter()
//	//putR.HandleFunc("/chatrooms", ph.Update)
//	//putR.Use(ph.MiddlewareValidatechatroom)
//
//	postR := sm.Methods(http.MethodPost).Subrouter()
//	postR.HandleFunc("/chatrooms", ph.Create)
//	//postR.Use(ph.MiddlewareValidatechatroom)
//
//	//deleteR := sm.Methods(http.MethodDelete).Subrouter()
//	//deleteR.HandleFunc("/chatrooms/{id:[0-9]+}", ph.Delete)
//
//	// handler for documentation
//	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
//	sh := middleware.Redoc(opts, nil)
//
//	getR.Handle("/docs", sh)
//	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
//
//	// CORS
//	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
//
//	// create a new server
//	s := http.Server{
//		Addr:         *bindAddress,                                     // configure the bind address
//		Handler:      ch(sm),                                           // set the default handler
//		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
//		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
//		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
//		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
//	}
//
//	// start the server
//	go func() {
//		l.Info("Starting server on port 9090")
//
//		err := s.ListenAndServe()
//		if err != nil {
//			l.Error("Error starting server", "error", err)
//			os.Exit(1)
//		}
//	}()
//
//	// trap sigterm or interupt and gracefully shutdown the server
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt)
//	signal.Notify(c, os.Kill)
//
//	// Block until a signal is received.
//	sig := <-c
//	log.Println("Got signal:", sig)
//
//	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
//	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//	s.Shutdown(ctx)
//}

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//// Binding from JSON
//type Login struct {
//	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
//	Password string `form:"password" json:"password" xml:"password" binding:"required"`
//}
//
//func main() {
//	router := gin.Default()
//
//	// Example for binding JSON ({"user": "manu", "password": "123"})
//	router.POST("/loginJSON", JsonHandler)
//
//	// Example for binding XML (
//	//	<?xml version="1.0" encoding="UTF-8"?>
//	//	<root>
//	//		<user>user</user>
//	//		<password>123</password>
//	//	</root>)
//	router.POST("/loginXML", func(c *gin.Context) {
//		var xml Login
//		if err := c.ShouldBindXML(&xml); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if xml.User != "manu" || xml.Password != "123" {
//			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//	})
//
//	// Example for binding a HTML form (user=manu&password=123)
//	router.POST("/loginForm", func(c *gin.Context) {
//		var form Login
//		// This will infer what binder to use depending on the content-type header.
//		if err := c.ShouldBind(&form); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if form.User != "manu" || form.Password != "123" {
//			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//	})
//
//	// Listen and serve on 0.0.0.0:8080
//	router.Run(":8080")
//}
//
//func JsonHandler(c *gin.Context) {
//	var json Login
//	if err := c.ShouldBindJSON(&json); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	if json.User != "manu" || json.Password != "123" {
//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//}
