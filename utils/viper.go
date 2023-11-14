package utils

import "github.com/spf13/viper"

var myViper *viper.Viper

func InitViper() {
	myViper = viper.New()
	myViper.SetConfigFile(".env")
	myViper.AutomaticEnv()
	myViper.AddConfigPath(GetRootPath())

	if err := myViper.ReadInConfig(); err != nil {
		panic(err)
		return
	}
}

func GetEnvValue(key string, defaultValue string) string {
	val := myViper.GetString(key)
	if val == "" {
		return defaultValue
	}
	return val
}
