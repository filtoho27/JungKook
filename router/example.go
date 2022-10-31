package router

import (
	"jungkook/api/example"
	"jungkook/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func exampleGroup(r *mux.Router) {
	ex := &example.ExampleSt{}
	// 不經過中介層
	sr := r.PathPrefix("/api/example").Subrouter()
	sr.Methods("GET").Path("/test_redis").Handler(recoverWrap(http.HandlerFunc(ex.ExRedisHandler)))
	sr.Methods("POST").Path("/test_db").Handler(recoverWrap(http.HandlerFunc(ex.ExDBHandler)))
	// 使用中介層
	sr2 := r.PathPrefix("/api/example").Subrouter()
	sr2.Use(middleware.CheckVendor)
	sr2.Methods("GET").Path("/test_acc1/{hid:[0-9]+}").Handler(recoverWrap(http.HandlerFunc(ex.ExAcc1Handler)))
}
