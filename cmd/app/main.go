package main

import (
	"github.com/Hatsunmikk/linkedin-automation/internal/logger"
)

//main is the entry point of the LinkedIn automation proof-of-concept.
//This application demonstrates browser automation, stealth techniques,
//and clean Go architecture for educational purposes only.

func main() {
	log := logger.New(true)

	log.Info("LinkedIn Automation PoC started")
	log.Debug("Debug logging enabled")
}
