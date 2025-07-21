package middleware

import (
	"net/http"

	"github.com/ggualbertosouza/game/pkg/logger"
	"go.uber.org/zap"
)

func LogMidd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		ctx := logger.NewContextWithTraceID(r.Context())
		
		logger.Info(ctx, "Request started",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remoteAddr", r.RemoteAddr),
		)
		
		ww := &statusWriter{ResponseWriter: w}
		
		next.ServeHTTP(ww, r.WithContext(ctx))
		
		logger.Info(ctx, "Request completed",
			zap.Int("status", ww.status),
		)
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	return n, err
}