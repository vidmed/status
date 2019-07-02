package http

import (
	"context"
	"fmt"
	"net/http"

	httputils "github.com/vidmed/status/pkg/utils/http"
)

const (
	TypeJson = "application/json"
	TypeXML  = "application/xml"
)

const (
	responseContentTypeKey uint8 = iota + 1
)

type contentType uint8

const (
	contentTypeJSON contentType = iota + 1
	contentTypeXML
)

func getResponseContentType(ctx context.Context) contentType {
	return ctx.Value(responseContentTypeKey).(contentType)
}

func contextWithResponseContentType(ctx context.Context, t contentType) context.Context {
	return context.WithValue(ctx, responseContentTypeKey, t)
}

func setupResponseContentType(ctx context.Context, r *http.Request) context.Context {
	ct := httputils.NegotiateContentType(r, []string{TypeJson, TypeXML}, TypeJson)
	fmt.Println(ct)
	switch ct {
	case TypeXML:
		return contextWithResponseContentType(ctx, contentTypeXML)
	default:
		return contextWithResponseContentType(ctx, contentTypeJSON)
	}
}

func xmlResponseContentType(ctx context.Context) bool {
	ct := getResponseContentType(ctx)
	return ct == contentTypeXML
}

func jsonResponseContentType(ctx context.Context) bool {
	ct := getResponseContentType(ctx)
	return ct == contentTypeJSON
}
