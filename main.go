package main

import ( 
	"errors"
	"os" 
	"time"

	"github.com/sirupsen/logrus" 
	log "github.com/BenjaminVanIseghem/be-mobile-logging/log"
)
func main() {
	w, err := os.Create("/promtail/logs/logfile1.log ")
	if err != nil {
	panic ( err ) }
	logger := logrus.New()
	logger.SetOutput(w)
	for {
		logger.Info(" Info message ")
		logger.Warn("Warning message")
		logger.Error("Error message", errors.New("Error")) time.Sleep(1200 ∗ time.Millisecond)
	}
}