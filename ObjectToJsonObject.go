package tipujson

import (
	"errors"
	"fmt"
	"reflect"
)

func objectToJsonObject(src interface{}) (result interface{}, err error) {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	return objectToJsonObjectByReflect(srcType, srcValue)
}
func objectToJsonObjectByReflect(srcType reflect.Type, srcValue reflect.Value) (result interface{}, err error) {
	switch srcType.Kind() {
	case reflect.Ptr:
		return objectToJsonObjectByReflect(srcType.Elem(), srcValue.Elem())
	case reflect.Interface:
		return objectToJsonObjectByReflect(srcValue.Elem().Type(), srcValue.Elem())
	case reflect.Slice:
		subResults := make([]interface{}, srcValue.Len())
		for index := 0; index < srcValue.Len(); index++ {
			subResults[index], err = objectToJsonObjectByReflect(srcType.Elem(), srcValue.Index(index))
			if err != nil {
				return
			}
		}
		result = subResults
		return
	case reflect.Struct:
		subResults := make(map[string]interface{})
		for index := 0; index < srcType.NumField(); index++ {
			valueField := srcValue.Field(index)
			typeField := srcType.Field(index)
			name := typeField.Name
			jsonName := typeField.Tag.Get("json")
			if jsonName == "" {
				jsonName = name
			}
			subResults[jsonName], err = objectToJsonObjectByReflect(typeField.Type, valueField)

		}
		result = subResults
		return
	case reflect.Map:
		subResults := make(map[string]interface{})
		for _, key := range srcValue.MapKeys() {
			tempKey := key
			for tempKey.Kind() == reflect.Interface || tempKey.Kind() == reflect.Ptr {
				tempKey = tempKey.Elem()
			}
			if tempKey.Kind() == reflect.String {
				subResults[key.String()], err = objectToJsonObjectByReflect(srcValue.MapIndex(tempKey).Type(), srcValue.MapIndex(tempKey))
				if err != nil {
					return
				}
			} else {
				panic(errors.New(fmt.Sprintf("objectToJsonField不能将key为%v的map转为JsonField", tempKey)))
			}
		}
		result = subResults
		return
	case reflect.Bool, reflect.String,
		reflect.Float32, reflect.Float64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = srcValue.Interface()
		return
	default:
		err = errors.New(fmt.Sprintf("TipuJson暂不支持将%v转换为JsonField", srcType))
		return
	}
}
