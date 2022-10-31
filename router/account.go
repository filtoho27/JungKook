package router

import (
	"jungkook/api/account"
	"net/http"

	"github.com/gorilla/mux"
)

func accountGroup(r *mux.Router) {
	ac := &account.AccountSt{}
	// 不經過中介層
	sr := r.PathPrefix("/api/account").Subrouter()
	sr.Methods("POST").Path("/sendemailcode").Handler(recoverWrap(http.HandlerFunc(ac.SendEmailCodeHandler)))
	sr.Methods("POST").Path("/register").Handler(recoverWrap(http.HandlerFunc(ac.RegisterHandler)))
	sr.Methods("POST").Path("/login").Handler(recoverWrap(http.HandlerFunc(ac.LoginHandler)))
}
