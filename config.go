package velocity

import (
	"fmt"
	"os"
)

type UserConfig struct {
	AppDir *string
}

type Config struct {
	AppDir string
}

func Validate(config UserConfig) (Config, error) {
	localConfig := Config{}

	if config.AppDir == nil || *config.AppDir == "" {
		fmt.Printf("'AppDir' not set, default to '<root>/app'\n")

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current working directory: %s\n", err)
			return Config{}, err
		}

		appDirPath := fmt.Sprint(cwd, "/app")
		stat, err := os.Stat(appDirPath)
		if err != nil {
			fmt.Printf("'app' directory or not exists.\n")
			return Config{}, err
		}
		if !stat.IsDir() {
			fmt.Printf("'app' directory or not exists.\n")
			return Config{}, err
		}

		localConfig.AppDir = appDirPath
	}

	return localConfig, nil
}
