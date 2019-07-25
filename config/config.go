package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// C 全局配置
var C *Config

// Config 配置
type Config struct {
	Server *Server       `toml:"server"`
	Redis  []RedisDetail `toml:"redis"`
	Log    *Log
}

// Log 日志配置
type Log struct {
	LogPath  string   `toml:"log_path"`
	LogLevel []string `toml:"log_level"`
}

// Server 服务端配置
type Server struct {
	ListenAddress string `toml:"listen_address"`
	PidFile       string `toml:"pid_file"`
}

// RedisDetail ...
type RedisDetail struct {
	Name    string   `toml:"name"`
	Address []string `toml:"address"`
	DB      int      `toml:"db"`
}

func init() {
	configPath := flag.String("c", "./conf/server.conf.toml", "config path")
	flag.Parse()
	if C == nil {
		fdata, openError := ioutil.ReadFile(*configPath)
		if openError != nil {
			fmt.Fprintf(os.Stderr, "%s", openError.Error())
		}
		C = &Config{}
		toml.Decode(string(fdata), C)
	}
}
