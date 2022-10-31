package api

import (
	Type "jungkook/commonType"
	"jungkook/foundation"
	"jungkook/modules/acc"
	"jungkook/modules/redis"
	"net/http"
)

func GetModule() *Type.ModuleType {
	module := &Type.ModuleType{
		Redis: redis.GetRedis(),
		Acc:   acc.GetAcc(),
		Time:  foundation.GetRealTime(),
	}
	return module
}

func GetVendor(r *http.Request) (vendor int) {
	platform := r.Header.Get("Vendor")
	if platform == "BBIN" {
		vendor = 1
	} else if platform == "BBGP" {
		vendor = 2
	}
	return
}
