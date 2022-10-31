package example

import (
	"jungkook/api"
	"jungkook/kernel"
	"net/http"

	"github.com/gorilla/mux"
)

type ExampleSt struct{}

func (ex *ExampleSt) ExRedisHandler(w http.ResponseWriter, r *http.Request) {
	module := api.GetModule()
	name := kernel.ParamString(r.URL.Query(), "name")
	result, err := exRedis(module, name)
	kernel.FormatResult(w, result, err)
}

func (ex *ExampleSt) ExAcc1Handler(w http.ResponseWriter, r *http.Request) {
	module := api.GetModule()
	module.Vendor = api.GetVendor(r)
	vars := mux.Vars(r)
	hid := kernel.PathInt(vars, "hid")
	result, err := exAcc1(module, hid)
	kernel.FormatResult(w, result, err)
}

func (ex *ExampleSt) ExDBHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	acc := kernel.ParamString(r.Form, "account")
	pwd := kernel.ParamString(r.Form, "password")
	result, err := exDB(acc, pwd)
	kernel.FormatResult(w, result, err)
}
