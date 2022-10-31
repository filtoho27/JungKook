package middleware

import (
	Error "jungkook/error"
	"jungkook/kernel"
	"net/http"
)

func CheckVendor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vendor := r.Header.Get("Vendor")
		if vendor != "BBIN" && vendor != "BBGP" {
			err := Error.CustomError{ErrMsg: "VENDOR_TYPE_ILLEGLE", ErrCode: 1338000002}
			kernel.FormatResult(w, nil, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
