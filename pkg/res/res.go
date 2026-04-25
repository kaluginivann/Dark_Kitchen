package res

import (
	"encoding/json"
	"net/http"

	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
)

func JSON(w http.ResponseWriter, data any, statusCode int, logger logger.Interface) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Error from encode data", "error", err)
	}
}
