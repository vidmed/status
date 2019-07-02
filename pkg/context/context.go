package context

import (
	"context"
	"fmt"
	"net/http"
)

type contextUtils int

const (
	contentTypeKey contextUtils = iota
)

const (
	ContentTypeJSON = iota + 1
	ContentTypeXML
)

func GetContentType(ctx context.Context) int {
	return ctx.Value(contentTypeKey).(int)
}

func SetContentType(ctx context.Context) int {
	return ctx.Value(contentTypeKey).(int)
}

func GetContentTypeFromRequest(r *http.Request) int {
	var (
		xml  = "application/xml"
		json = "application/json"
	)

	ct := NegotiateContentType(r, []string{json, xml}, json)
	fmt.Println(ct)
	switch ct {
	case xml:
		return ContentTypeXML
	default:
		return ContentTypeJSON
	}
}
