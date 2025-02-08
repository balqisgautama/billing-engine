package config

import (
	configmodels "billing-engine/internal/models/config"
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string, fileName string) (appConfig configmodels.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&appConfig)
	PrintConfig(appConfig)

	return
}

func PrintConfig(c configmodels.Config) {
	data, _ := json.MarshalIndent(c, "", "\t")
	log.Println(string(data))
}
