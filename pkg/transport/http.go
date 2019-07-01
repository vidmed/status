package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vidmed/status/pkg/endpoint"

	"github.com/gorilla/mux"
	"github.com/vidmed/status/pkg/service"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s service.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := endpoint.New(s, logger)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path(" /orders/{orderId}/status").Handler(httptransport.NewServer(
		e.GetStatusEndpoint,
		decodeGetProfileRequest,
		encodeResponse,
		options...,
	))

	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
