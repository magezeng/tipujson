package tipujson

import (
	"errors"
	"fmt"
	. "github.com/magezeng/tipujson/Modles"
	"reflect"
)

func JsonStringToObject(src string, direction interface{}) (err error) {

	//对变量进行处理，得到Type和Value
	directionType := reflect.TypeOf(direction)
	directionValue := reflect.ValueOf(direction)
	return JsonStringToObjectByReflect(src, directionType, directionValue)
}

func JsonStringToObjectByReflect(src string, directionType reflect.Type, directionValue reflect.Value) (err error) {

	var srcJsonField *JsonField
	srcJsonField, err = getJsonFieldFromString(src)
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
		err = jsonFieldToObject(srcJsonField, directionType, directionValue)
		return
	default:
		err = errors.New(fmt.Sprintf("暂不支持%v的映射", directionType))
	}
	return
}

func ObjectToJsonString(src interface{}) (result string, err error) {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	return ReflectObjectToJsonString(srcType, srcValue)
}

func ReflectObjectToJsonString(srcType reflect.Type, srcValue reflect.Value) (result string, err error) {
	switchType := srcType
	for switchType.Kind() == reflect.Ptr {
		switchType = switchType.Elem()
	}
	result, err = objectToJsonString(srcType, srcValue)
	return
}

func ObjectToObject(src interface{}, direction interface{}) (err error) {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	err = ReflectObjectToObject(srcType, srcValue, direction)
	return
}

func ReflectObjectToObject(srcType reflect.Type, srcValue reflect.Value, direction interface{}) (err error) {
	var field *JsonField
	field, err = objectToJsonField(srcType, srcValue)
	directionType := reflect.TypeOf(direction).Elem()
	directionValue := reflect.ValueOf(direction).Elem()
	err = jsonFieldToObject(field, directionType, directionValue)
	return
}
