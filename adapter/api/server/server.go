package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cart2 "github.com/nade-harlow/E-commerce/adapter/api/controllers/cart"
	product2 "github.com/nade-harlow/E-commerce/adapter/api/controllers/product"
	user2 "github.com/nade-harlow/E-commerce/adapter/api/controllers/user"
	"github.com/nade-harlow/E-commerce/adapter/api/routes"
	"github.com/nade-harlow/E-commerce/adapter/repository/database/client"
	"github.com/nade-harlow/E-commerce/adapter/repository/database/redisql"
	"github.com/nade-harlow/E-commerce/adapter/repository/notification"
	"github.com/nade-harlow/E-commerce/core/repositories/cart"
	"github.com/nade-harlow/E-commerce/core/repositories/product"
	"github.com/nade-harlow/E-commerce/core/repositories/user"
	"github.com/nade-harlow/E-commerce/ports/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	router := setupRoute()
	db := client.InitializeConnection()
	mg := notification.MailgunRepository{}
	mg.NewMailgunRepository()
	tr := notification.TwilloRepository{}
	tr.NewTwillo()
	redis := redisql.Redis{}
	redis.NewRedisClient()
	redis.PingRedis()

	products := product2.NewProductController(services.NewProductService(product.New(db)))
	users := user2.NewUserController(services.NewUserService(user.New(db), &mg, &tr, &redis))
	carts := cart2.NewCartController(services.NewCartService(cart.New(db), &redis))

	routes.CartRoutes(router, carts)
	routes.ProductRoutes(router, products)
	routes.UserRoutes(router, users)
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
			log.Printf("An error occurred while starting the server: %s", err)
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

func setupRoute() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("core/template/change_password.html")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}
