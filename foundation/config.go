package foundation

import (
	"fmt"
	Error "jungkook/error"
	customLog "jungkook/log"
	"strings"

	"github.com/spf13/viper"
)

func GetViperConfig(path string, name string, ext string) (config *viper.Viper) {
	config = viper.New()
	config.SetConfigName(name)
	config.SetConfigType(ext)
	config.AddConfigPath(path)
	err := config.ReadInConfig()
	if err != nil {
		msg := fmt.Sprintf("GET_%s_CONFIG_FAILED", strings.ToUpper(name))
		customLog.WriteLog("sys", "ConfigError", err, Error.CustomError{ErrMsg: msg}, "")
	}
	return
}