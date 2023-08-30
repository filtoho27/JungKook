package api

import (
	Type "jungkook/commonType"
	"jungkook/foundation"
	"jungkook/modules/redis"
)

func GetModule() *Type.ModuleType {
	module := &Type.ModuleType{
		Redis: redis.GetRedis(),
		Time:  foundation.GetRealTime(),
	}
	return module
}
