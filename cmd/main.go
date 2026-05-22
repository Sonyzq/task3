// main.go
package main

import (
	"log"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		log.Fatal("index.html нет в этой директории ")
	}
	logger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)
	err := srv.HTTP.ListenAndServe()
	time.Sleep(500 * time.Millisecond)

	if err != nil {
		logger.Fatal("Ошибка запуска сервера:", err)
	}
}
