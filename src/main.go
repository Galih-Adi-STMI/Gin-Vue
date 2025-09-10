
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"tracker/controller"
	"tracker/repository"
	"tracker/service"
)

func main() {
	log.Println("Menjalankan Server...")

	db, err := initDB()
	if err != nil {
		log.Fatalf("Gagal Inisialisasi Database: %v\n", err)
	}

	router := gin.Default()

	transactionRepository := repository.NewPGTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepository)
	controller.NewController(&controller.Config{
		R:                 router,
		TransactionService: transactionService,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit
	log.Println("Shutdown signal received, shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Close(); err != nil {
		log.Fatalf("Ada Masalah saat menutup server: %v\n", err)
	}

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	<-ctx.Done()
	log.Println("Hitung Mundur 5 Detik.")
	log.Println("Server Ditutup.")
}
