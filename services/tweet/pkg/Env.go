package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println(err)
	}
	return value

}
