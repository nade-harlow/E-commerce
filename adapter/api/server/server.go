package server

import (
	"context"
	"github.com/gin-gonic/gin"
	product2 "github.com/nade-harlow/E-commerce/adapter/api/controllers/product"
	"github.com/nade-harlow/E-commerce/adapter/api/routes"
	"github.com/nade-harlow/E-commerce/adapter/repository/database/client"
	"github.com/nade-harlow/E-commerce/core/repositories"
	"github.com/nade-harlow/E-commerce/ports"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	router := gin.Default()
	db := client.InitializeConnection()
	prod := product2.NewProductController(ports.NewService(repositories.New(db)))
	routes.DefineRoutes(router, prod)
	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	wait := make(chan os.Signal) // creates a channel that will be used to wait for a signal

	log.Printf("Server Started at Port%s", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("An error occurred with the server: %s", err)
			return
		}
	}() // go routine to start the server
	// sends a signal to the wait channel if there is an interrupt signal
	signal.Notify(wait, os.Interrupt)

	<-wait // waits here until a signal is received
	log.Printf("Shutting down the server...")

	time.Sleep(time.Second * 2) // sleep for 1 second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shuts down the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
	log.Printf("Server exits successfully")
}
