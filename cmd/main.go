package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	transporthttp "github.com/vidmed/status/pkg/transport/http"

	"github.com/vidmed/status/pkg/service"
	"github.com/vidmed/status/pkg/service/storage"
	"github.com/vidmed/status/pkg/service/storage/dummy"

	"github.com/go-kit/kit/log"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "transport", "HTTP")
	}

	var s storage.Storage
	{
		s = dummy.NewDummyStorage()
		defer s.Close()
	}

	var svc service.Service
	{
		svc = service.NewWithMiddleware(s, []service.Middleware{service.LoggingMiddleware(logger)})
	}

	var h http.Handler
	{
		h = transporthttp.MakeHTTPHandler(svc, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	hs := &http.Server{Addr: *httpAddr, Handler: h}

	go func() {
		logger.Log("msg", "Server started", "addr", *httpAddr)
		if err := hs.ListenAndServe(); err != http.ErrServerClosed {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		logger.Log("msg", "Server error occurred", "error", err)
	case <-stop:
		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		logger.Log("msg", "Shutting down", "timeout", timeout)

		if err := hs.Shutdown(ctx); err != nil {
			logger.Log("msg", "Shutdown failed", "error", err)
		} else {
			logger.Log("msg", "Shutdown success")
		}
		cancel()
	}
}
