package router

import (
	"jungkook/api/shopping"
	"net/http"

	"github.com/gorilla/mux"
)

func shoppingGroup(r *mux.Router) {
	sp := &shopping.ShoppingSt{}
	// 不經過中介層
	sr := r.PathPrefix("/api/shopping").Subrouter()
	sr.Methods("GET").Path("/getshoppinglist").Handler(recoverWrap(http.HandlerFunc(sp.ShoppingListHandler)))
}
