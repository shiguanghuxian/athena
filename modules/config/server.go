package config

import (
	"errors"
	"os"

	"github.com/naoina/toml"
)

/* server配置实例话和结构体定义 */

// ServerConfig server配置
type ServerConfig struct {
	Debug  bool                `toml:"debug"`
	Logger *LoggerConfig       `toml:"logger"`
	HTTP   *HTTPListenerConfig `toml:"http"`
	NSQ    *NsqConfig          `toml:"nsq"`
	Redis  *RedisConfig        `toml:"redis"`
}

// NewServerConfig 实例化server配置
func NewServerConfig(path string) (*ServerConfig, error) {
	if path == "" {
		return nil, errors.New("配置文件路径不能为空")
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config := new(ServerConfig)
	if err := toml.NewDecoder(f).Decode(config); err != nil {
		return nil, err
	}
	// 默认值处理
	if config.Logger == nil {
		config.Logger = &LoggerConfig{
			Name:      "athena",
			Verbose:   true,
			SystemLog: false,
		}
	}
	return config, nil
}
