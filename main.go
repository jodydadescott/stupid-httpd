package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/jodydadescott/stupid-httpd/httpserver"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	listen := os.Getenv("LISTEN")

	if listen == "" {
		listen = ":8080"
		zap.L().Debug("Listening on :8080 (default). Set LISTEN to override. For example LISTEN=:8081")
	} else {
		zap.L().Debug(fmt.Sprintf("Listening on %s", listen))
	}

	zap.L().Debug("Starting")
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	httpserver := httpserver.NewServer(listen)

	<-sig

	httpserver.Shutdown()

	zap.L().Debug("Shutting Down")

}
