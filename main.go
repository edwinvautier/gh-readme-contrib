package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"time"

	"github.com/edwinvautier/gh-readme-contrib/api/routes"
	"github.com/edwinvautier/gh-readme-contrib/shared/database"
	"github.com/edwinvautier/gh-readme-contrib/shared/env"
	"github.com/edwinvautier/gh-readme-contrib/shared/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Connect to database and execute migrations
	cfg := database.Config{}
	cfg.User = env.GoDotEnvVariable("DB_USER")
	cfg.Password = env.GoDotEnvVariable("DB_PASSWORD")
	cfg.Port, _ = strconv.ParseInt(env.GoDotEnvVariable("DB_PORT"), 10, 0)
	cfg.Name = env.GoDotEnvVariable("DB_NAME")
	cfg.Host = env.GoDotEnvVariable("DB_HOST")
	err := database.Init(cfg)
	helpers.DieOnError("database connection failed", err)
	database.Migrate()

	// Setup router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		MaxAge:           50 * time.Second,
		AllowCredentials: true,
	}))
	port := env.GoDotEnvVariable("PORT")
	if port == "" {
		port = ":8000"
	} else {
		port = ":" + port
	}
	log.Info("try to run app on port ", port)
	routes.Init(router)
	go func() {
		if err := router.Run(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// ----------------- CLOSE APP -----------------

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
}
