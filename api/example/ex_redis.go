package example

import (
	Type "jungkook/commonType"
)

func exRedis(module *Type.ModuleType, name string) (result Type.ExRedis, err error) {
	dtime := module.Time.Now().Format("2006-01-02 15:04:05")
	err = module.Redis.SetExRedisData(name, dtime)
	if err != nil {
		return
	}
	result, err = module.Redis.GetExRedisData()
	return
}
