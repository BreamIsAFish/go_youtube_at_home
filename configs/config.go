package configs

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var config *Config

type MongoDB struct {
	URI 					string 	`mapstructure:"uri" validate:"required"`
	DatabaseName 	string 	`mapstructure:"database_name" validate:"required"`
	Timeout 			int 		`mapstructure:"timeout" validate:"required"`
}

type JWT struct {
	AccessSecret string `mapstructure:"access_secret" validate:"required"`
}

type App struct {
	Port int `mapstructure:"port" validate:"required"`
}

type Config struct {
	MongoDB			*MongoDB 	`mapstructure:"mongodb" validate:"required"`
	JWT 				*JWT 			`mapstructure:"jwt" validate:"required"`
	App 				*App 			`mapstructure:"app" validate:"required"`
}

func readFile(cfg interface{}, conPath, conFileName, envPrefix string) error {
	ymlConfig := viper.New()
	ymlConfig.AddConfigPath(conPath)
	ymlConfig.SetConfigName(conFileName)
	viper.SetConfigType("yml")
	ymlConfig.SetEnvPrefix(envPrefix)
	ymlConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	ymlConfig.AutomaticEnv()

	err := ymlConfig.ReadInConfig()
	if err != nil {
		return err
	}
	err = ymlConfig.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func SetupViperConfig() error {
	// viper.SetConfigName("config")
	// viper.AddConfigPath("/configs/config.yml")
	// // viper.AutomaticEnv()
	// viper.SetConfigType("yml")
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get current frame")
	}

	dir := filepath.Dir(filename)
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	
	config = &Config{}
	err := readFile(config, dir, "config", "")
	if err != nil {
		fmt.Errorf("Error reading config file, %s", err)
		return err
	}
	return nil
}

func GetConfig() *Config {
	return config
}