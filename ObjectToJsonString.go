package tipujson

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func objectToJsonString(src interface{}) (result string, err error) {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	return objectToJsonStringByReflect(srcType, srcValue)
}
func objectToJsonStringByReflect(srcType reflect.Type, srcValue reflect.Value) (result string, err error) {
	switch srcType.Kind() {
	case reflect.Ptr:
		return objectToJsonStringByReflect(srcType.Elem(), srcValue.Elem())
	case reflect.Interface:
		return objectToJsonStringByReflect(srcValue.Elem().Type(), srcValue.Elem())
	case reflect.Slice:
		subResults := make([]string, srcValue.Len())
		for index := 0; index < srcValue.Len(); index++ {
			subResults[index], err = objectToJsonStringByReflect(srcType.Elem(), srcValue.Index(index))
			if err != nil {
				return
			}
		}
		result = "[" + strings.Join(subResults, ",") + "]"
		return
	case reflect.Struct:
		subResults := make([]string, srcType.NumField())
		for index := 0; index < srcType.NumField(); index++ {
			valueField := srcValue.Field(index)
			typeField := srcType.Field(index)
			name := typeField.Name
			jsonName := typeField.Tag.Get("json")
			if jsonName == "" {
				jsonName = name
			}
			var subValueString string
			subValueString, err = objectToJsonStringByReflect(typeField.Type, valueField)
			if err != nil {
				return
			}
			subResults[index] = "\"" + jsonName + "\":" + subValueString
		}
		result = "{" + strings.Join(subResults, ",") + "}"
		return
	case reflect.Map:
		subResults := make([]string, srcValue.Len())
		for index, key := range srcValue.MapKeys() {
			var subValueString string
			subValueString, err = objectToJsonStringByReflect(srcValue.MapIndex(key).Type(), srcValue.MapIndex(key))
			if err != nil {
				return
			}
			subResults[index] = "\"" + fmt.Sprint(key) + "\":" + subValueString
		}
		result = "{" + strings.Join(subResults, ",") + "}"
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		result = fmt.Sprint(srcValue)
		return
	case reflect.String:
		result = "\"" + fmt.Sprint(srcValue) + "\""
		return
	case reflect.Bool:
		if srcValue.Bool() {
			result = "true"
		} else {
			result = "false"
		}
		return
	default:
		err = errors.New(fmt.Sprintf("TipuJson暂不支持将%v转换为JsonString", srcType))
		return
	}
}
