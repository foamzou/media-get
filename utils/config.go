package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/foamzou/audio-get/logger"
)

const ConfigFile = ".media-get.json"
const ConfigENV = "ENV_MEDIA_GET"

type MediaGetConfig struct {
	Proxy map[string]string `json:"proxy"`
}

var config *MediaGetConfig

func InitConfig() {
	var conf MediaGetConfig
	// get config from env first
	configEnv := os.Getenv(ConfigENV)
	if configEnv != "" {
		err := json.Unmarshal([]byte(configEnv), &conf)
		if err != nil {
			logger.Error("json decode has error.", err)
			return
		}
		config = &conf
		return
	}
	// get config from file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Debug("got user home dir failed.", err)
		return
	}
	f, err := os.Open(filepath.Join(homeDir, ConfigFile))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Debug("open file err.", err)
		}
		return
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		logger.Error("json decode has error.", err)
		return
	}
	config = &conf
}

func GetConfig() *MediaGetConfig {
	return config
}
