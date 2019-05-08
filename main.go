package main

import (
	"errors"
	"strconv"
	"time"

	log "github.com/BenjaminVanIseghem/be-mobile-logging/log"
)

func main() {
	counter := 1
	for {
		logFile, logger := log.CreateLogBuffer("promtail/logs/", "XML-converter", strconv.Itoa(counter))
		for i := 1; i <= 30; i++ {
			logger.Info("Info " + strconv.Itoa(i))
		}
		log.Error(logger, "Error in loop", errors.New("error"), &logFile)
		log.Flush(logFile)
		time.Sleep(5000 * time.Millisecond)
		counter++
	}
}
