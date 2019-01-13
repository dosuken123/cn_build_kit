package main

import (
	"reflect"

	"github.com/dosuken123/cn_build_kit/command"
)

var typeRegistry = make(map[string]reflect.Type)

func init() {
	typeRegistry["Clone"] = reflect.TypeOf(command.Clone{})
}

func MakeInstance(name string) interface{} {
	v := reflect.New(typeRegistry[name]).Elem()
	// Maybe fill in fields here if necessary
	return v.Interface()
}
