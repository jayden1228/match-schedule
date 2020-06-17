package configs

import (
	"fmt"

	"github.com/timest/env"
)

//EnvConfig EnvConfig
var EnvConfig *config

type config struct {
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
	Mysql      struct {
		Host    string `default:"120.25.31.149"`
		Port    string `default:"3306"`
		DBName string `default:"match-schedule"`
		User    string `default:"root"`
		Pwd     string `default:"tEKIZtC0CXjg"`
		Charset string `default:"utf8mb4"`
	}
	Redis struct {
		Host   string `default:"120.78.51.14"`
		Port   string `default:"7000"`
		Pwd    string `default:""`
		Prefix string `default:"match-schedule|"`
	}
}

func init() {
	EnvConfig = new(config)
	env.IgnorePrefix()
	err := env.Fill(EnvConfig)
	fmt.Println(EnvConfig)
	if err != nil {
		panic(err)
	}
}
