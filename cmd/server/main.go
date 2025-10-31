package main

import (
	"task_manager/internal/config"
	"task_manager/internal/handlers"
	"task_manager/internal/models"
	"os/exec"
	"runtime"
	"time"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	cfg := config.Load()
	setupLogging(cfg)

	log.WithFields(logrus.Fields{
		"app":     cfg.App.Name,
		"version": cfg.App.Version,
		"port":    cfg.Server.Port,
	}).Info("Starting application")

	lib := models.CreateContainer()
	router := handlers.CreateRouter(lib, cfg)
	
	adr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Infof("Server starting on %s", adr)

	go func() {
		time.Sleep(300 * time.Millisecond)
		openBrowser("http://" + adr)
	}()

	if err := router.Run(adr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupLogging(cfg *config.Config) {
	level, err := logrus.ParseLevel(cfg.Logging.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	if cfg.Logging.Format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}


func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	_ = cmd.Start()
}