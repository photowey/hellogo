package app

import (
	"fmt"

	"github.com/hellogo/internal/config"
	"github.com/hellogo/pkg/logger"
)

/**
 * app
 */

func Run(conf string) error {
	if conf == "" {
		conf = "config.toml"
	}

	if err := config.Load(conf); err != nil {
		return fmt.Errorf("load the app config error:%w", err)
	}

	if err := logger.New(config.Log()); err != nil {
		return fmt.Errorf("init the logger error:%w", err)
	}

	// init db
	// init cache
	// init mq

	return nil
}
