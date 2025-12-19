package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/database"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers/middleware"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
)

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env variables")
	}

	server_ip := os.Getenv("SERVER_IP")

	// Set up the API
	igdb.Setup()

	// Set up the db
	database.Setup()

	defer func() {
		if err := database.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Make router
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthOptional())
	router.SetTrustedProxies(nil)

	// Set up the routes
	handlers.Setup(router)

	// Look for certs in common locations (repo path for local runs, /certs for container mounts)
	localCert := "../certs/localhost.pem"
	localKey := "../certs/localhost-key.pem"
	containerCert := "/certs/localhost.pem"
	containerKey := "/certs/localhost-key.pem"

	certFile := ""
	keyFile := ""
	if _, err := os.Stat(localCert); err == nil {
		certFile = localCert
		keyFile = localKey
	} else if _, err := os.Stat(containerCert); err == nil {
		certFile = containerCert
		keyFile = containerKey
	}

	if certFile != "" && keyFile != "" {
		addr := server_ip + ":8443"
		log.Printf("Starting server with TLS on %s", addr)
		if err := router.RunTLS(addr, certFile, keyFile); err != nil {
			log.Fatal(err)
		}
	} else {
		addr := server_ip + ":8080"
		log.Printf("Starting server without TLS on %s", addr)
		if err := router.Run(addr); err != nil {
			log.Fatal(err)
		}
	}
}
