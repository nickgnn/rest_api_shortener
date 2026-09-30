package main

import (
	"fmt"
	"rest_api_shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	// TODO: delete
	fmt.Println(cfg)

	// TODO: init logger: slog
	// TODO: init storage: sqlite
	// TODO: init router: chi, "chi render"
	// TODO: run server
}
