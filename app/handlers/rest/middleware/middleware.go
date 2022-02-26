package middleware

import (
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"net/http"
)

func SetDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx = ctxlib.SetDB(ctx)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
