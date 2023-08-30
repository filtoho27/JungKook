package router

import (
	"errors"
	customLog "jungkook/log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetMap(r *mux.Router) {
	accountGroup(r)
	shoppingGroup(r)
}

func recoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rc := recover()
			if rc != nil {
				var err error
				switch t := rc.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				errMsg := err.Error()
				customLog.WritePanicLog(r, errMsg)
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
