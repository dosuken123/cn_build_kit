package service

import (
	"log"
	"os"
)

func (s Service) GetServiceDir() string {
	return getCurrentDir() + "/services/" + s.Name
}

func (s Service) GetSrcDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/src"
}

func (s Service) GetDockerFileDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/dockerfile"
}

func (s Service) GetCacheDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/cache"
}

func (s Service) GetDataDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/data"
}

func (s Service) GetScriptDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/script"
}

func (s Service) GetLogDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/log"
}

func (s Service) GetExampleDir() string {
	return getCurrentDir() + "/services/" + s.Name + "/example"
}

func getCurrentDir() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return dir
}
