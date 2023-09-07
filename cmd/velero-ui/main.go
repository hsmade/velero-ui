package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/api"
	"log/slog"
	"os"
)

func main() {
	setupLogger()

	r := gin.Default()
	err := api.New(r.Group("/api/v1"))
	if err != nil {
		slog.Error("setting up router group", "err", err.Error())
		os.Exit(-1)
	}

	r.Static("/assets", "./web/dist/assets")
	r.StaticFile("/", "./web/dist/index.html")
	err = r.Run(":8080")
	if err != nil {
		slog.Error("web server exited", "err", err.Error())
		os.Exit(-1)
	}
}

// setupLogger allows to enable debug logging
func setupLogger() {
	gin.SetMode(gin.ReleaseMode)

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	if os.Getenv("DEBUG") != "" {
		gin.SetMode(gin.DebugMode)
		opts.Level = slog.LevelDebug
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
