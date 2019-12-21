package tipujson

import (
	"errors"
	"fmt"
	"reflect"
)

func JsonStringToObject(src string, direction interface{}) (err error) {
	directionType := reflect.TypeOf(direction)
	directionValue := reflect.ValueOf(direction)
	return JsonStringToObjectByReflect(src, directionType, directionValue)
}

func JsonStringToObjectByReflect(src string, directionType reflect.Type, directionValue reflect.Value) (err error) {
	var jsonObject interface{}
	jsonObject, err = getJsonObjectFromString([]byte(src))
	if err != nil {
		return
	}
	switchDirectionType := directionType
	for switchDirectionType.Kind() == reflect.Ptr {
		switchDirectionType = switchDirectionType.Elem()
	}
	// 分类型进行扫描反射字段,并映射值到目标
	switch switchDirectionType.Kind() {
	case reflect.Struct, reflect.Map, reflect.Interface, reflect.Slice:
		err = jsonObjectToObjectByReflect(jsonObject, directionType, directionValue)
		return
	default:
		err = errors.New(fmt.Sprintf("暂不支持%v的映射", directionType))
	}
	return
}

func ObjectToJsonString(src interface{}) (result string, err error) {
	return objectToJsonString(src)
}

func ObjectToJsonStringByReflect(srcType reflect.Type, srcValue reflect.Value) (result string, err error) {
	result, err = objectToJsonStringByReflect(srcType, srcValue)
	return
}

func ObjectToJsonObject(src interface{}) (result interface{}, err error) {
	return objectToJsonObject(src)
}

func ObjectToJsonObjectByReflect(srcType reflect.Type, srcValue reflect.Value) (result interface{}, err error) {
	return objectToJsonObjectByReflect(srcType, srcValue)
}
