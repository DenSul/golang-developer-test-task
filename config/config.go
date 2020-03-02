package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Redis RedisConfiguration
}

type RedisConfiguration struct {
	Hostname string
	Port     int8
	Username string
	Password string
}

func Load() Configurations {
	envPrefix := flag.String("env", "dev", "env for app")
	flag.Parse()

	// Set the file name of the configurations file
	viper.SetConfigName("config." + *envPrefix)

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading config file, %s", err))
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode into struct, %v", err))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())

	return configuration
}
