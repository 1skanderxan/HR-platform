package main

import (
	"log"

	"github.com/yourname/hr-portal-backend/config"
	myhttp "github.com/yourname/hr-portal-backend/internal/delivery/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config yuklanmadi:", err)
	}

	server := myhttp.NewServer(cfg)
	log.Printf("Server %s portda ishlamoqda...", cfg.ServerPort)
	if err := server.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Server xatosi:", err)
	}
}
