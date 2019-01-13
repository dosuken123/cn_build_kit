package main

import (
	"io/ioutil"
	"sync"

	"github.com/BurntSushi/toml"
)

type Src struct {
	repo_url    string
	clone_depth int
}

type Service struct {
	Name       string
	Host       string
	Port       int
	Executable string
	Src        Src
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
