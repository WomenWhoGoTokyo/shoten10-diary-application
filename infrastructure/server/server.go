package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/config"
)

func NewServer(handler http.Handler) *http.Server {
	c := config.Config
	return &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("0.0.0.0:%s", c.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
