package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	helpers "github.com/hellogo/internal/common"
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
	Env      string         `toml:"env" json:"env"` // dev | test | stage | prod
	Host     string         `toml:"host" json:"host"`
	Log      LoggerConfig   `toml:"log" json:"log"`
	Database DatabaseConfig `toml:"database" json:"database"`
	Redis    RedisConfig    `toml:"redis" json:"redis"`
	RabbitMQ RabbitMQConfig `toml:"rabbitmq" json:"rabbitmq"`
}

type LoggerConfig struct {
	Level           string `toml:"level" json:"level"`
	Path            string `toml:"path" json:"path"`
	FileName        string `toml:"file_name" json:"fileName"`
	MaxSize         int    `toml:"max_size" json:"maxSize"`
	MaxAgeDay       int    `toml:"max_age_day" json:"maxAgeDay"`
	CompressEnabled bool   `toml:"compress_enabled" json:"compressEnabled"`
	StdoutEnabled   bool   `toml:"stdout_enabled" json:"stdoutEnabled"`
}

type DatabaseConfig struct {
	Driver   string `toml:"driver" json:"driver"`
	Host     string `toml:"host" json:"host"`
	Port     int32  `toml:"port" json:"port"`
	Database string `toml:"database" json:"database"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
}

type RedisConfig struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	Password string `toml:"password" json:"password"`
	Database int8   `toml:"database" json:"database"`
}

type RabbitMQConfig struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	User     string `toml:"user"  json:"user"`
	Password string `toml:"password" json:"password"`
	Vhosts   string `toml:"vhosts" json:"vhosts"`
}

var conf Config

func Load(path string) (err error) {
	_, err = toml.DecodeFile(path, &conf)
	fmt.Println("--- >>> init the app config successfully <<< ---")

	if index := helpers.StringsContains(profiles, conf.Env); index == -1 {
		return fmt.Errorf("the env config candidate value is:%v", profiles)
	}

	// printConfig()

	return
}

func printConfig() {
	pretty, _ := helpers.ToJSONPretty(conf)
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
