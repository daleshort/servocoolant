package logger

import(
	log "github.com/sirupsen/logrus"
	"os"
)

func GetLog() *log.Logger{
	logger :=log.New()
	logger.SetLevel(log.TraceLevel)
	logger.SetOutput(os.Stdout)
	logger.Info("Logger Started")
	return logger
}