package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	// "jungkook/modules/acc"
	"jungkook/modules/mysql"
	// "jungkook/modules/portal"
	"jungkook/modules/redis"
	"jungkook/router"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handleHealthz)
	os.Setenv("CONFIGPATH", "./")

	// 服務初始化(只執行一次)
	var onceInit sync.Once
	onceInit.Do(serviceInit)
	router.SetMap(r)

	log.Println("server start at :80")
	_ = http.ListenAndServe(":80", r)
}

func serviceInit() {
	// DB初始化
	mysql.Init()
	// RB初始化
	redis.Init()
	// // ACC初始化
	// acc.Init()
	// // Portal初始化
	// portal.Init()
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "ok")
}
