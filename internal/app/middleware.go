package app

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(res, req)
		logger := log.Info()
		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Msg("received a HTTP request")
	})
}
