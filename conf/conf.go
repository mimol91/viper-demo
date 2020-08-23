package conf

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var R RootConf

var stringToLocationHookFun = func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String {
		return data, nil
	}
	if t != reflect.TypeOf(time.Location{}) {
		return data, nil
	}
	return time.LoadLocation(data.(string))
}

type RootConf struct {
	App           AppConf
	Curlec        CurlecConf
	Elasticsearch ElasticsearchConf
}
type AppConf struct {
	Name string
}
type CurlecConf struct {
	Host       string
	MerchantID string
	Location   time.Location
	CacheTTL   time.Duration
}
type ElasticsearchConf struct {
	Host       string
	Repository RepositoryConf
}

type RepositoryConf struct {
	Mapping map[string]string
}

func Init() {
	hooks := []mapstructure.DecodeHookFunc{
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		stringToLocationHookFun,
	}

	viperDecodeHook := mapstructure.ComposeDecodeHookFunc(hooks...)
	viper.AddConfigPath("conf")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("default")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading cfg %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&R, viper.DecodeHook(viperDecodeHook)); err != nil {
			panic(fmt.Sprintf("Failed to unmarshall cfg %v", err))
		}

		fmt.Println("Update config")
	})

	if err := viper.Unmarshal(&R, viper.DecodeHook(viperDecodeHook)); err != nil {
		panic(fmt.Sprintf("Failed to unmarshall cfg %v", err))
	}
}
