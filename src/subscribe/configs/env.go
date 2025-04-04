package configs

import (
	"os"

	"subscribe/logger"

	"github.com/spf13/viper"
)

func SetDevEnv() {
	// root
	os.Setenv("ROOT_PATH", "/home/kyle/code/kyle-solana/src/subscribe")
	viper.Set("ROOT_PATH", "/home/kyle/code/kyle-solana/src/subscribe")

	// logger
	os.Setenv("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
	viper.Set("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
	logger.InitLogger()
	logger.Log.Info("hi! i'm solana subscription.")

	// data
	os.Setenv("DATA_PATH", viper.GetString("ROOT_PATH")+"/../data")
	viper.Set("DATA_PATH", viper.GetString("ROOT_PATH")+"/../data")

	// viper
	os.Setenv("CONFIG_PATH", viper.GetString("ROOT_PATH")+"/configs/config.yaml")
	if err := InitConfig(); err != nil {
		logger.Log.Fatalf("check errors, %v", err)
	}
}

func SetProdEnv() {
	// root
	os.Setenv("ROOT_PATH", "/app")
	viper.Set("ROOT_PATH", "/app")

	// logger
	os.Setenv("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
	viper.Set("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
	logger.InitLogger()
	logger.Log.Info("hi! i'm solana subscription.")

	// viper
	os.Setenv("CONFIG_PATH", viper.GetString("ROOT_PATH")+"/configs/config.yaml")
	if err := InitConfig(); err != nil {
		logger.Log.Fatalf("check errors, %v", err)
	}
}
