package log

import (
	// "log"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	// logPath := viper.GetString("LOG_PATH")
	// logFileName := viper.GetString("LOG_FILENAME")
	// logFile := path.Join(logPath, logFileName)
	// fd, err := os.Create(logFile)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error init log file: %s \n", err))
	// }
	// Log = Logger{log.New(fd, "", log.LstdFlags|log.Lshortfile)}
}
