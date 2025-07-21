package handlers

import (
	"net/http"

	"github.com/ggualbertosouza/game/pkg/logger"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	logger.Info(ctx, "Deu bom os logs")

	// return errors.ErrUnauthorized
	// return errors.NewUnauthorizedError("Invalid credentials")
	// return fmt.Errorf("database connection failed")
	// return errors.NewNotFoundError("User not found")

	w.Write([]byte("OK"))
	return nil
}