package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/BurntSushi/toml"
)

type Src struct {
	RepoURL    string `toml:"repo_url"`
	CloneDepth int
}

type Meta struct {
	SrcPath        string
	DockerFilePath string
	CachePath      string
	DataPath       string
}

type Service struct {
	Name       string
	Host       string
	Port       int
	Executable string
	Src        Src
	Group      string
	Meta       Meta
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

func (s Service) GenerateMeta() {
	s.Meta.SrcPath = "services/" + s.Name + "/src"
	s.Meta.DockerFilePath = "services/" + s.Name + "/dockerfile"
	s.Meta.CachePath = "services/" + s.Name + "/cache"
	s.Meta.DataPath = "services/" + s.Name + "/data"
}

func ValidateIfServiceNameUnique() bool {
	// TODO
	return true
}

func GetAllServiceNames() []string {
	unique := make(map[string]bool)
	var names []string
	for _, service := range GetConfig().Services {
		if _, value := unique[service.Group]; !value {
			names = append(names, service.Name)
			unique[service.Name] = true
		}
	}
	return names
}

func GetService(serviceName string) (Service, error) {
	var s Service
	for _, service := range GetConfig().Services {
		if service.Name == serviceName {
			s = service
			return s, nil
		}
	}
	return s, errors.New("Service Not Found")
}

func GetAllServicesOfGroup(groupName string) ([]Service, error) {
	var services []Service
	for _, service := range GetConfig().Services {
		if service.Group == groupName {
			services = append(services, service)
		}
	}

	if len(services) == 0 {
		return services, errors.New("Group Not Found")
	}

	return services, nil
}

func ResolveTargetName(ambiguousName string) ([]Service, error) {
	var results []Service
	if service, error := GetService(ambiguousName); error == nil {
		fmt.Println("1")
		results = append(results, service)
		fmt.Printf("service: %+v\n", service)
		return results, nil
	} else if services, error := GetAllServicesOfGroup(ambiguousName); error == nil {
		results = append(results, services...)
		return results, nil
	} else {
		return results, errors.New("Failed to resolve target name. Given name: " + ambiguousName)
	}
}

func GetAllGroupNames() []string {
	unique := make(map[string]bool)
	var names []string
	for _, service := range GetConfig().Services {
		if _, value := unique[service.Group]; !value {
			names = append(names, service.Group)
			unique[service.Group] = true
		}
	}
	return names
}

func IsIncludedInServiceNames(groupName string) bool {
	for _, name := range GetAllServiceNames() {
		if groupName == name {
			return true
		}
	}

	return false
}

func IsIncludedInGroupNames(groupName string) bool {
	for _, name := range GetAllGroupNames() {
		if groupName == name {
			return true
		}
	}

	return false
}
