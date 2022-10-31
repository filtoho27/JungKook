package shopping

import (
	"jungkook/api"
	"jungkook/kernel"
	"net/http"
)

type ShoppingSt struct{}

func (ex *ShoppingSt) ShoppingListHandler(w http.ResponseWriter, r *http.Request) {
	module := api.GetModule()
	err := getShoppingList(module)
	kernel.FormatResult(w, nil, err)
}
