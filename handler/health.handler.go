package handler

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "API OK", nil)
}
