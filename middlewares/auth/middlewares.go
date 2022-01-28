package auth

import (
	"net/http"
	"strings"
)

func SetHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		username, _, ok := req.BasicAuth()

		if !ok {
			_, _ = res.Write([]byte(`{"error" : "Basic Auth is not defined for this request"}`))
			return
		}

		isValid := strings.Contains(username, "-zs")
		if !isValid {
			_, _ = res.Write([]byte(`{"error": "This api is only accessible by zopsmart employees"}`))
			return
		}

		h.ServeHTTP(res, req)
	})
}

func CheckAuthToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

	})
}
