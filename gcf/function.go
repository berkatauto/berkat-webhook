package gcf

import (
	"net/http"

	"github.com/Befous/webhook"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", enableCORS(webhook.PostBalasan))
}

// enableCORS adalah fungsi middleware untuk menambahkan header CORS ke handler HTTP
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set header untuk izinkan permintaan dari domain tertentu
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight options request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Lanjutkan ke handler berikutnya
		next(w, r)
	}
}
