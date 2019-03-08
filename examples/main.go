package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/karrick/goterm"
	"github.com/natefinch/lumberjack"
)

const (
	logsDirname    = "logs"
	logsFilename   = "service.log"
	logsMaxDays    = 30
	logsMaxBackups = 50
)

func main() {
	if !goterm.IsTerminal(os.Stderr.Fd()) {
		log.SetOutput(&lumberjack.Logger{
			Filename:   filepath.Join(logsDirname, logsFilename),
			MaxAge:     logsMaxDays,
			MaxBackups: logsMaxBackups,
		})
	}

	log.Printf("log output")
}
