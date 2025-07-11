package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/johnnyFR26/GoMicroservice/internal/config"
	"github.com/johnnyFR26/GoMicroservice/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro conectando ao banco:", err)
	}

	if err := db.AutoMigrate(&model.APIKey{}); err != nil {
		log.Fatal("Erro migrando banco:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}).Methods("GET")

	// 🔹 Aqui depois você pode adicionar seu middleware e rotas protegidas
	// Exemplo:
	// apiKeyRepo := repository.NewAPIKeyRepository(db)
	// api := r.PathPrefix("/api").Subrouter()
	// api.Use(middleware.APIKeyAuth(apiKeyRepo))
	// api.HandleFunc("/convert", convertHandler).Methods("POST")

	port := cfg.Port
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", port),
	}

	go func() {
		log.Printf("Servidor iniciado na porta %s 🚀", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Desligando servidor...")
	if err := srv.Close(); err != nil {
		log.Fatal("Erro ao encerrar servidor:", err)
	}
	log.Println("Servidor desligado com sucesso 👋")
}
