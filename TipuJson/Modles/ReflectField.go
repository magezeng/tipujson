package Modles

import "reflect"

type ReflectField struct {
	Type     reflect.Type
	Value    reflect.Value
	JsonName string
	Name     string
}
