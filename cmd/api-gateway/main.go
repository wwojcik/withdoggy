package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/withdoggy/withdoggy/gateway"
	"go.uber.org/zap"
)

var Version string

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigsCh := make(chan os.Signal, 1)
	doneCh := make(chan bool, 1)
	errorCh := make(chan error, 1)

	signal.Notify(sigsCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	go func() {
		sugar.Infof("running api-gateway version: %s \n", Version)
		go gateway.RunServer(ctx, logger)
	}()

	select {
	case sig := <-sigsCh:
		fmt.Println(sig)
		cancel()
	case err := <-errorCh:
		sugar.Error(err)
	case <-doneCh:
		sugar.Info("Exiting")
	}
}
