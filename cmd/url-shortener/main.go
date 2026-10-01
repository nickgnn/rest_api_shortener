package main

import (
	"fmt"
	"log/slog"
	"os"
	"rest_api_shortener/internal/config"
	"rest_api_shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	// TODO: delete
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	//log = log.With(slog.String("env", cfg.Env)) // добавляет строчку с уточнением окружения во все строки лога
	log.Info("START URL-SHORTENER", slog.String("env", cfg.Env))
	log.Debug("DEBUG messages are enabled")

	storage, err := sqlite.NewStorage(cfg.Storage)
	if err != nil {
		log.Error("Failed to initialize storage", err)
		os.Exit(1)
	}
	_ = storage

	// TODO: init router: chi, "chi render"
	// TODO: run server
}

// параметр env, потому что локально будут текстовые логи
// на сервере в окружении dev/prod будет json
// на dev - debug, а на prod - не ниже info
func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
