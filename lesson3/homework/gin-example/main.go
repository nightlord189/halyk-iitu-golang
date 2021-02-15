package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"iitu/gin-example/handlers"
	"iitu/gin-example/middlewares"
	"iitu/gin-example/repo"

	"github.com/gin-gonic/gin"
)

func init() {
	err := repo.Init()
	if err != nil {
		panic(err)
	}
}

func startServer() (*http.Server, <-chan error) {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middlewares.AccessLog(false))

	router.Use(gin.Recovery())

	returnChannel := make(chan error, 1)

	err := handlers.MapRoutes(router)
	if err != nil {
		returnChannel <- err
		return nil, returnChannel
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			returnChannel <- err
		}
	}()

	return srv, returnChannel
}

func main() {
	server, startServerError := startServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Server started")
	select {
	case sig := <-quit:
		fmt.Println(fmt.Sprintf("Received signal: %s\nStopping server...", sig.String()))
	case err := <-startServerError:
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	fmt.Println("Server stopped")
}
