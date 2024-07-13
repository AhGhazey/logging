package middleware

import "net/http"

func ContextHolderMiddlewareChiV5(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {

		// TODO: extract information from request auth header and store it in context

		next.ServeHTTP(rw, request)
	})
}
