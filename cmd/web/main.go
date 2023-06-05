package main

import (
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shortLinks/cmd/web/handler"
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

	storageFlag := flag.String("storage", "postgres", "Select storage: postgres (by default) or cache")
	flag.Parse()

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

	staticFileServer := http.FileServer(http.Dir("./ui/static"))
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	mux := http.NewServeMux()
	mux.HandleFunc("/home", handlers.Home)
	mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))
	if *storageFlag == "postgres" {
		log.Print("Remember to apply migrations with Makefile scripts")
		mux.HandleFunc("/", handlers.HandleShortUrlRedirect)
		mux.HandleFunc("/api/get-short-link", handlers.GetShortLink)
		mux.HandleFunc("/api/get-default-link", handlers.GetDefaultLink)
	} else if *storageFlag == "cache" {
		mux.HandleFunc("/", handlers.HandleShortUrlRedirectWithCache)
		mux.HandleFunc("/api/get-short-link", handlers.GetShortLinkFromCache)
		mux.HandleFunc("/api/get-default-link", handlers.GetDefaultLinkFromCache)
	} else {
		log.Fatal("Failed to initialize handlers: storage is not specified")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	go func() {
		signalType := <-quit
		signal.Stop(quit)
		log.Print("Shutdown server...")
		log.Print("Signal type: ", signalType)

		os.Exit(0)
	}()

	log.Print("Starting server on :8080")
	err = http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
