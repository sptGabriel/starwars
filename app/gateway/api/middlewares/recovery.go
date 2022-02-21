package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

func Recovery(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				logger := hlog.FromRequest(req)
				logger.Print("panic occurred:", r)
				msg, _ := json.Marshal(map[string]string{"Message": "Internal Error"})
				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write(msg)
				if err != nil {
					logger.Error().Err(err)
				}
			}
		}()
		next.ServeHTTP(w, req)
	})
}
