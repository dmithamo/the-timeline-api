package middleware

import (
	"fmt"
	"net/http"

	"github.com/dmithamo/timelineapi/pkg/utils"
)

// CheckAuth checks that authorization header is present and valid
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				utils.SendJSONResponse(w, http.StatusUnauthorized,
					&utils.GenericJSONRes{
						Message: "no valid authorization token. Login to continue",
						Data:    nil,
					})
				return
			}

			// other errs
			utils.SendJSONResponse(w, http.StatusInternalServerError,
				&utils.GenericJSONRes{
					Message: fmt.Sprintf("err reading authorization: %v", err.Error()),
					Data:    nil,
				})

			return
		}

		// if is authorized
		next.ServeHTTP(w, r)
	})
}
