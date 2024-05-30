package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size

	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func RequestMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseData := &responseData{}
			var requestCopy bytes.Buffer
			loggingWriter := loggingResponseWriter{
				ResponseWriter: w,
				responseData:   responseData,
			}
			defer r.Body.Close()
			if _, err := requestCopy.ReadFrom(r.Body); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			logger.Info(fmt.Sprintf("request body: %v", string(requestCopy.Bytes())))

			r.Body = io.NopCloser(&requestCopy)
			next.ServeHTTP(&loggingWriter, r)

			logger.Info(
				"got incoming HTTP request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", responseData.status),
			)
		})
	}
}
