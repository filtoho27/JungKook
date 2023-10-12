package api

import (
	Type "jungkook/commonType"
	"jungkook/modules/redis"
)

func GetModule() *Type.ModuleType {
	module := &Type.ModuleType{
		Redis: redis.GetRedis(),
	}
	return module
}
