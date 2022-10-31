package kernel

import (
	"net/url"
	"strconv"
)

func PathInt(vars map[string]string, key string) (result int) {
	if v, ok := vars[key]; ok {
		result, _ = strconv.Atoi(v)
	}
	return
}

func ParamInt(params url.Values, key string) (result int) {
	if v, ok := params[key]; ok {
		result, _ = strconv.Atoi(v[0])
	}
	return
}

func ParamString(params url.Values, key string) (result string) {
	if v, ok := params[key]; ok {
		result, _ = url.QueryUnescape(v[0])
	}
	return
}

func ParamFloat(params url.Values, key string) (result float64) {
	if v, ok := params[key]; ok {
		result, _ = strconv.ParseFloat(v[0], 64)
	}
	return
}

func ParamIntA(params url.Values, key string) (result []int) {
	if params, ok := params[key]; ok {
		for _, v := range params {
			param, _ := strconv.Atoi(v)
			result = append(result, param)
		}
	}
	return
}

func ParamStringA(params url.Values, key string) (result []string) {
	if params, ok := params[key]; ok {
		for _, v := range params {
			param, _ := url.QueryUnescape(v)
			result = append(result, param)
		}
	}
	return
}

func ParamBool(params url.Values, key string) (result bool) {
	if v, ok := params[key]; ok {
		result, _ = strconv.ParseBool(v[0])
	}
	return
}