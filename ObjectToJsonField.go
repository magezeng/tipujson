package TipuJson

import (
	"errors"
	"fmt"
	. "github.com/magezeng/TipuJson/Modles"
	"reflect"
)

func objectToJsonField(srcType reflect.Type, srcValue reflect.Value) (result *JsonField, err error) {
	switch srcType.Kind() {
	case reflect.Ptr:
		return objectToJsonField(srcType.Elem(), srcValue.Elem())
	case reflect.Interface:
		return objectToJsonField(srcValue.Elem().Type(), srcValue.Elem())
	case reflect.Slice:
		result = &JsonField{
			Type: JsonFieldTypeList,
		}
		subResults := make([]*JsonField, srcValue.Len())
		for index := 0; index < srcValue.Len(); index++ {
			subResults[index], err = objectToJsonField(srcType.Elem(), srcValue.Index(index))
			if err != nil {
				return
			}
		}
		result.Content = subResults
		return
	case reflect.Struct:
		result = &JsonField{
			Type: JsonFieldTypeMap,
		}
		subResults := make(map[string]*JsonField)
		for index := 0; index < srcType.NumField(); index++ {
			valueField := srcValue.Field(index)
			typeField := srcType.Field(index)
			name := typeField.Name
			jsonName := typeField.Tag.Get("json")
			if jsonName == "" {
				jsonName = name
			}
			subResults[jsonName], err = objectToJsonField(typeField.Type, valueField)

		}
		result.Content = subResults
		return
	case reflect.Map:
		result = &JsonField{
			Type: JsonFieldTypeMap,
		}
		subResults := make(map[string]*JsonField)
		for _, key := range srcValue.MapKeys() {
			tempKey := key
			for tempKey.Kind() == reflect.Interface || tempKey.Kind() == reflect.Ptr {
				tempKey = tempKey.Elem()
			}
			if tempKey.Kind() == reflect.String {
				subResults[key.String()], err = objectToJsonField(srcValue.MapIndex(tempKey).Type(), srcValue.MapIndex(tempKey))
				if err != nil {
					return
				}
			} else {
				panic(errors.New(fmt.Sprintf("objectToJsonField不能将key为%v的map转为JsonField", tempKey)))
			}
		}
		result.Content = subResults
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		result = &JsonField{
			Type:    JsonFieldTypeNumber,
			Content: fmt.Sprint(srcValue),
		}
		return
	case reflect.String:
		result = &JsonField{
			Type:    JsonFieldTypeString,
			Content: fmt.Sprint(srcValue),
		}
		return
	case reflect.Bool:
		if srcValue.Bool() {
			result = &JsonField{
				Type:    JsonFieldTypeBool,
				Content: "true",
			}
		} else {
			result = &JsonField{
				Type:    JsonFieldTypeBool,
				Content: "false",
			}
		}
		return
	default:
		err = errors.New(fmt.Sprintf("TipuJson暂不支持将%v转换为JsonField", srcType))
		return
	}
}
