package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexey-dobry/rental-service/internal/pkg/logger"
	"github.com/alexey-dobry/rental-service/internal/pkg/logger/zap"
	"github.com/alexey-dobry/rental-service/internal/pkg/validator"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/domain/jwt"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/repository/pg"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/server/grpc"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logger     zap.Config  `yaml:"logger"`
	gRPC       grpc.Config `yaml:"grpc"`
	Repository pg.Config   `yaml:"repository"`
	JWT        jwt.Config  `yaml:"jwt"`
}

func MustLoad() Config {
	logger := zap.NewLogger(zap.Config{}).WithFields().WithFields("layer", "config")
	var cfg Config
	configPath := ParseFlag(cfg, &logger)

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		logger.Fatalf("Failed to read config on path(%s): %s", configPath, err)
	}

	if err := validator.V.Struct(&cfg); err != nil {
		logger.Fatalf("Failed to validate config: %s", err)
	}

	return cfg
}

func ParseFlag(cfg Config, logger *logger.Logger) string {
	configPath := flag.String("config", "configs/config.yaml", "config file path")
	configHelp := flag.Bool("help", false, "show configuration help")

	if *configHelp {
		headerText := "Configuration options:"
		help, err := cleanenv.GetDescription(&cfg, &headerText)
		if err != nil {
			(*logger).Fatalf("error getting configuration description: %s", err.Error())
		}

		fmt.Println(help)
		os.Exit(0)
	}

	return *configPath
}
