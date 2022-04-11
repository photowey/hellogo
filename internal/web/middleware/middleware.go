package middleware

import (
	"mime"
	"net/http"

	"github.com/hellogo/pkg/logger"
)

var (
	ContentTypeHeader        = "Content-Type"
	MediaTypeApplicationJson = "application/json"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.Info("recv a %s request from %s", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get(ContentTypeHeader)
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if mediaType != MediaTypeApplicationJson {
			http.Error(w, "invalid Content-Type", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, req)
	})
}
