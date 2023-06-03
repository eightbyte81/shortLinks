package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shortLinks/internal/repository"
	"shortLinks/internal/repository/postgres"
	"shortLinks/internal/service"
	"syscall"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal("Failed to load configs: ", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env variables: ", err)
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal("Failed to initialize db: ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := NewHandler(services)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/api/get-short-link", handlers.GetShortLink)
	mux.HandleFunc("/api/get-default-link", handlers.GetDefaultLink)

	log.Print("Starting server on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Shutdown server...")

	err = db.Close()
	if err != nil {
		log.Fatal("Failed to close db connection: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
