package app

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Servidor struct {
	App     *fiber.App
	Puerto  string
	Debug   bool
	Timeout time.Duration
}

func (s *Servidor) IniciarServidor() {
	err := s.App.Listen(s.Puerto)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Servidor) IniciarRutas() {
	configurarRutas(s.App)
}
