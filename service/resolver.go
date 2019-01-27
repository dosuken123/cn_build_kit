package service

import (
	"errors"
	"strings"
)

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

func GetConcatenatedServices(concatenatedName string) ([]Service, error) {
	var services []Service
	serviceNames := strings.Split(concatenatedName, ",")
	for _, serviceName := range serviceNames {
		for _, service := range GetConfig().Services {
			if service.Name == serviceName {
				services = append(services, service)
				break
			}
		}
	}

	if len(services) == 0 {
		return services, errors.New("Concatenated Services Not Found")
	}

	return services, nil
}

func ResolveTargetName(ambiguousName string) ([]Service, error) {
	var results []Service
	if service, error := GetService(ambiguousName); error == nil {
		results = append(results, service)
		return results, nil
	} else if services, error := GetAllServicesOfGroup(ambiguousName); error == nil {
		results = append(results, services...)
		return results, nil
	} else if services, error := GetConcatenatedServices(ambiguousName); error == nil {
		results = append(results, services...)
		return results, nil
	} else {
		return results, errors.New("Failed to resolve target name. Given name: " + ambiguousName)
	}
}
