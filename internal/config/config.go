package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	helpers `github.com/hellogo/internal/common/str`
	"gopkg.in/yaml.v3"

	"github.com/hellogo/internal/jsonz"
)

/**
 * config
 */

type Profile string

var (
	envDev   = "dev"
	envTest  = "test"
	envStage = "final"
	envProd  = "prod"

	profiles = []string{envDev, envTest, envStage, envProd}
)

type Config struct {
	Env      string         `toml:"env" json:"env" yaml:"env"` // dev | test | stage | prod
	Host     string         `toml:"host" json:"host" yaml:"host"`
	Log      LoggerConfig   `toml:"log" json:"log" yaml:"log"`
	Database DatabaseConfig `toml:"database" json:"database" yaml:"database"`
	Redis    RedisConfig    `toml:"redis" json:"redis" yaml:"redis"`
	RabbitMQ RabbitMQConfig `toml:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
}

type LoggerConfig struct {
	Level           string `toml:"level" json:"level" yaml:"level"`
	Path            string `toml:"path" json:"path" yaml:"path"`
	FileName        string `toml:"file_name" json:"fileName" yaml:"fileName"`
	MaxSize         int    `toml:"max_size" json:"maxSize" yaml:"maxSize"`
	MaxAgeDay       int    `toml:"max_age_day" json:"maxAgeDay" yaml:"maxAgeDay"`
	CompressEnabled bool   `toml:"compress_enabled" json:"compressEnabled" yaml:"compressEnabled"`
	StdoutEnabled   bool   `toml:"stdout_enabled" json:"stdoutEnabled" yaml:"stdoutEnabled"`
}

type DatabaseConfig struct {
	Driver   string `toml:"driver" json:"driver" yaml:"driver"`
	Host     string `toml:"host" json:"host" yaml:"host"`
	Port     int32  `toml:"port" json:"port" yaml:"port"`
	Database string `toml:"database" json:"database" yaml:"database"`
	Username string `toml:"username" json:"username" yaml:"username"`
	Password string `toml:"password" json:"password" yaml:"password"`
}

type RedisConfig struct {
	Host     string `toml:"host" json:"host" yaml:"host"`
	Port     int    `toml:"port" json:"port" yaml:"port"`
	Password string `toml:"password" json:"password" yaml:"password"`
	Database int8   `toml:"database" json:"database" yaml:"database"`
}

type RabbitMQConfig struct {
	Host     string `toml:"host" json:"host" yaml:"host"`
	Port     int    `toml:"port" json:"port" yaml:"port"`
	User     string `toml:"user"  json:"user" yaml:"user"`
	Password string `toml:"password" json:"password" yaml:"password"`
	Vhosts   string `toml:"vhosts" json:"vhosts" yaml:"vhosts"`
}

var conf Config

func LoadToml(path string) (err error) {
	_, err = toml.DecodeFile(path, &conf)
	fmt.Println("--- >>> init the app toml config successfully <<< ---")

	if index := helpers.ArrayContains(profiles, conf.Env); index == -1 {
		return fmt.Errorf("the env config candidate value is:%v", profiles)
	}

	// printConfig()

	return
}

func LoadYaml(path string) (err error) {
	if yamlConfig, err := os.Open(path); err != nil {
		return err
	} else {
		if err = yaml.NewDecoder(yamlConfig).Decode(&conf); err != nil {
			return err
		}
	}

	fmt.Println("--- >>> init the app yaml config successfully <<< ---")

	if index := helpers.ArrayContains(profiles, conf.Env); index == -1 {
		return fmt.Errorf("the env config candidate value is:%v", profiles)
	}

	// printConfig()

	return
}

func printConfig() {
	pretty, _ := jsonz.Pretty(conf)
	fmt.Println(pretty)
}

// ---------------------- getter ----------------------

func Env() string {
	return conf.Env
}

func Host() string {
	return conf.Host
}

func Log() LoggerConfig {
	return conf.Log
}

func Database() DatabaseConfig {
	return conf.Database
}

func Redis() RedisConfig {
	return conf.Redis
}

func RabbitMQ() RabbitMQConfig {
	return conf.RabbitMQ
}
