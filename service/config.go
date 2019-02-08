package service

import (
	"io/ioutil"
	"sync"

	"github.com/BurntSushi/toml"
)

type User struct {
	Name    string `toml:"name"`
	Group   string `toml:"group"`
	GID     string `toml:"gid"`
	UID     string `toml:"uid"`
	HomeDir string `toml:"home_dir"`
}

type Src struct {
	RepoURL    string `toml:"repo_url"`
	CloneDepth int    `toml:"clone_depth"`
}

type Variable struct {
	Key   string `toml:"repo_url"`
	Value string `toml:"clone_depth"`
}

type Service struct {
	Name       string
	Host       string
	Port       int
	Executable string
	Src        Src
	Group      string
	User       User
	Variables  map[string]string
}

type Config struct {
	Services []Service `toml:"service"`
}

var configMap Config
var confMux = &sync.Mutex{}

func init() {
	data, err := ioutil.ReadFile("./config.toml")

	if err != nil {
		panic(err)
	}

	if _, err := toml.Decode(string(data), &configMap); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	confMux.Lock()
	defer confMux.Unlock()
	return configMap
}
