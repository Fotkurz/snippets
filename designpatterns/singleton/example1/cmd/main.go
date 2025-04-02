package main

import (
	"gosingleton/internal/logging"
)

func main() {
	log := logging.NewLogger()

	log.Info("Message")
}
