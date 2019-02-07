package service

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func (s Service) GetInjectableVariables() string {
	variables := s.GetVariables()
	var str strings.Builder
	for key, value := range variables {
		if value != "" {
			str.WriteString(fmt.Sprintf("%s=%s ", key, value))
		}
	}

	return str.String()
}

func (s Service) GetVariables() map[string]string {
	var variables map[string]string
	variables = make(map[string]string)

	for _, service := range GetAllServices() {
		variables[fmt.Sprintf("CN_%s_DIR", strings.ToUpper(service.Name))] = service.GetServiceDir()
		variables[fmt.Sprintf("CN_%s_SCRIPT_DIR", strings.ToUpper(service.Name))] = service.GetScriptDir()
		variables[fmt.Sprintf("CN_%s_SRC_DIR", strings.ToUpper(service.Name))] = service.GetSrcDir()
		variables[fmt.Sprintf("CN_%s_CACHE_DIR", strings.ToUpper(service.Name))] = service.GetCacheDir()
		variables[fmt.Sprintf("CN_%s_DATA_DIR", strings.ToUpper(service.Name))] = service.GetDataDir()
		variables[fmt.Sprintf("CN_%s_LOG_DIR", strings.ToUpper(service.Name))] = service.GetLogDir()
		variables[fmt.Sprintf("CN_%s_EXAMPLE_DIR", strings.ToUpper(service.Name))] = service.GetExampleDir()
		variables[fmt.Sprintf("CN_%s_DOCKERFILE_DIR", strings.ToUpper(service.Name))] = service.GetDockerFileDir()
		variables[fmt.Sprintf("CN_%s_HOST", strings.ToUpper(service.Name))] = service.Host
		variables[fmt.Sprintf("CN_%s_PORT", strings.ToUpper(service.Name))] = strconv.Itoa(service.Port)
		variables[fmt.Sprintf("CN_%s_EXECUTABLE", strings.ToUpper(service.Name))] = service.Executable
	}

	variables["CN_SELF_DIR"] = s.GetServiceDir()
	variables["CN_SELF_SCRIPT_DIR"] = s.GetScriptDir()
	variables["CN_SELF_SRC_DIR"] = s.GetSrcDir()
	variables["CN_SELF_CACHE_DIR"] = s.GetCacheDir()
	variables["CN_SELF_DATA_DIR"] = s.GetDataDir()
	variables["CN_SELF_LOG_DIR"] = s.GetLogDir()
	variables["CN_SELF_EXAMPLE_DIR"] = s.GetExampleDir()
	variables["CN_SELF_DOCKERFILE_DIR"] = s.GetDockerFileDir()
	variables["CN_SELF_HOST"] = s.Host
	variables["CN_SELF_PORT"] = strconv.Itoa(s.Port)
	variables["CN_SELF_EXECUTABLE"] = s.Executable

	variables["GOPATH"] = filepath.Join(s.GetServiceDir(), "go")

	variables["USER"] = s.GetUserName()
	variables["GROUP"] = s.GetUserGroup()
	variables["UID"] = s.GetUID()
	variables["GID"] = s.GetGID()

	return variables
}
