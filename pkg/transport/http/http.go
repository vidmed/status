package http

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strconv"

	"github.com/vidmed/status/pkg/endpoint/dto"

	"github.com/vidmed/status/pkg/endpoint"

	"github.com/gorilla/mux"
	"github.com/vidmed/status/pkg/service"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s service.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := endpoint.New(s, logger)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerBefore(setupContext),
	}

	// todo (d.medvedev) perhaps make server factory depend on request headers (xml, json)
	r.Methods("GET").Path("/orders/{id}/status").Handler(httptransport.NewServer(
		e.GetStatusEndpoint,
		decodeGetStatusRequest,
		encodeResponse,
		options...,
	))

	return r
}

func setupContext(ctx context.Context, r *http.Request) context.Context {
	ctx = setupResponseContentType(ctx, r)

	//fmt.Println(httputils.GetContentTypeFromRequest(r))
	return ctx
}

func decodeGetStatusRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return &dto.OrderStatusRequest{OrderId: intID}, nil
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if xmlResponseContentType(ctx) {
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		return xml.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(response)
	}
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.WriteHeader(codeFrom(err))
	// todo (d.medvedev) handle error
	encodeResponse(ctx, w, map[string]interface{}{
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
