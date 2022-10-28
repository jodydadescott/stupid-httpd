package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/jodydadescott/stupid-httpd/httpserver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {

	err := initLogger()
	if err != nil {
		return err
	}

	ipPort := os.Getenv("LISTEN")

	if ipPort == "" {
		ipPort = ":8080"
		zap.L().Debug("Listening on :8080 (default). Set LISTEN to override. For example LISTEN=:8081")
	} else {
		zap.L().Debug(fmt.Sprintf("Listening on %s", ipPort))
	}

	httpserver := httpserver.NewServer()

	zap.L().Debug("Starting")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done
		httpserver.Shutdown()
	}()

	httpserver.Listen(ipPort)

	zap.L().Debug("Shutting Down")
	return nil
}

func initLogger() error {

	zapConfig := &zap.Config{
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	zapConfig.Encoding = "json"
	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logger, err := zapConfig.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)

	return nil
}
