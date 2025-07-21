package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/ggualbertosouza/game/api/errors"
	"github.com/ggualbertosouza/game/pkg/logger"
	"go.uber.org/zap"
)


func ErrorMidd(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				ctx := r.Context()
				
				var apiErr *errors.APIError
				
				switch e := err.(type) {
				case *errors.APIError:
					apiErr = e
				case error:
					apiErr = errors.ErrInternalServer
					apiErr.Message = e.Error()
				default:
					apiErr = errors.ErrInternalServer
					apiErr.Message = "An unexpected error occurred"
				}

				logger.Error(ctx, "Request error",
					zap.Int("status", apiErr.StatusCode),
					zap.String("code", apiErr.Code),
					zap.String("error", apiErr.Message),
				)

				w.WriteHeader(apiErr.StatusCode)
				json.NewEncoder(w).Encode(apiErr)
			}
		}()
		
		next(w, r)
	}
}