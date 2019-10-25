package TipuJson

import (
	"errors"
	"fmt"
	"github.com/magezeng/TipuJson/Modles"
	"reflect"
)

func StringToObj(src string, direction interface{}) (err error) {

	//对变量进行处理，得到Type和Value
	directionType := reflect.TypeOf(direction).Elem()
	directionValue := reflect.ValueOf(direction).Elem()
	return StringToObjByReflect(src, directionType, directionValue)
}

func StringToObjByReflect(src string, directionType reflect.Type, directionValue reflect.Value) (err error) {

	var srcJsonField *Modles.JsonField
	srcJsonField, err = GetJsonFieldFromString(src)
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
		err = ReflectJsonFieldToAnyType(srcJsonField, directionType, directionValue)
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
	switch switchType.Kind() {
	case reflect.Map, reflect.Slice, reflect.Struct:
		result, err = AnyToJsonString(srcType, srcValue)
		return
	default:
		err = errors.New(fmt.Sprintf("ObjectToJsonString暂不支持将%v转化为String", srcType))
		return
	}
}
