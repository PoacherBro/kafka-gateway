package global

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	FileName string // config file name

	ServerName string `toml:"server_name"`
	HTTPAddr string `toml:"http_addr"`
	GRPCAddr string `toml:"grpc_addr"`
	LogLevel string `toml:"log_level"`

	KafkaConfig *kafkaConfig `toml:"kafka"`
}

type kafkaConfig struct {
	Addr []string `toml:"addr"`
	BackupAddr []string `toml:"bak_addr"`
}

var gConf *Config
func GetConfig() Config {
	if gConf == nil {
		panic("global config is nil, please init first")
	}
	return *gConf
}

func ParseConfig(fileName string) (*Config, error) {
	cfg := new(Config)
	if _, err := toml.DecodeFile(fileName, cfg); err != nil {
		log.Fatalf("parse config(file=%s) err=%v", fileName, err)
		return nil, err
	}
	cfg.FileName = fileName
	gConf = cfg

	return cfg, nil
}

