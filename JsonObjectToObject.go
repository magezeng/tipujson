package tipujson

import (
	"fmt"
	"github.com/magezeng/tipujson/ErrorMaker"
	"reflect"
)

func jsonObjectToObject(jsonObject interface{}, direction interface{}) (err error) {

	directionType := reflect.TypeOf(direction)
	directionValue := reflect.ValueOf(direction)
	return jsonObjectToObjectByReflect(jsonObject, directionType, directionValue)
}
func jsonObjectToObjectByReflect(jsonObject interface{}, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if jsonObject == nil {
		//遇到空值的情况直接不进行映射
		return
	}
	switch waitSetType.Kind() {
	case reflect.Ptr:
		waitSetType = waitSetType.Elem()
		if waitSetValue.IsNil() {
			waitSetValue.Set(reflect.New(waitSetType))
		}
		waitSetValue = waitSetValue.Elem()
		err = jsonObjectToObjectByReflect(jsonObject, waitSetType, waitSetValue)
	case reflect.Struct:
		err = jsonObjectToStruct(jsonObject, waitSetType, waitSetValue)
	case reflect.Slice:
		err = jsonObjectToSlice(jsonObject, waitSetType, waitSetValue)
	case reflect.Interface:
		err = jsonObjectToInterface(jsonObject, waitSetType, waitSetValue)
	case reflect.Map:
		err = jsonObjectToMap(jsonObject, waitSetType, waitSetValue)
	default:
		err = jsonObjectToBaseType(jsonObject, waitSetType, waitSetValue)
	}
	return
}

func jsonObjectToInterface(jsonObject interface{}, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	// 假如说双方类型未匹配上，则直接出错退出
	if waitSetType.Kind() != reflect.Interface {
		err = ErrorMaker.GetError(fmt.Sprintf("不能使用ReflectJsonFieldToInterface方法将JsonObject映射到%v中", waitSetType))
		return
	}
	waitSetValue.Set(reflect.ValueOf(jsonObject))
	return
}

func jsonObjectToMap(jsonObject interface{}, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	jsonObjectType := reflect.ValueOf(jsonObject).Kind()
	// 假如说双方类型未匹配上，则直接出错退出
	if jsonObjectType != reflect.Map || waitSetType.Kind() != reflect.Map {
		err = ErrorMaker.GetError(fmt.Sprintf("不能使用ReflectJsonFieldToMap方法将%s映射到%v中", jsonObjectType, waitSetType))
		return
	}
	if waitSetValue.IsNil() {
		waitSetValue.Set(reflect.MakeMap(waitSetType))
	}
	for key, value := range jsonObject.(map[string]interface{}) {
		tempValue := reflect.New(waitSetType.Elem())
		err = jsonObjectToObjectByReflect(value, waitSetType.Elem(), tempValue)
		if err != nil {
			return
		}
		waitSetValue.SetMapIndex(reflect.ValueOf(key), tempValue)
	}
	return
}

func jsonObjectToStruct(jsonObject interface{}, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Struct || reflect.ValueOf(jsonObject).Kind() != reflect.Map {
		err = ErrorMaker.GetError(fmt.Sprintf("不能使用ReflectJsonFieldToStruct方法将%s映射到%v中", reflect.ValueOf(jsonObject).Kind(), waitSetType))
		return
	}
	//content := jsonObject.Content.(map[string]*JsonField)
	for i := 0; i < waitSetType.NumField(); i++ {
		valueField := waitSetValue.Field(i)
		typeField := waitSetType.Field(i)
		if typeField.Anonymous {
			err = jsonObjectToObjectByReflect(jsonObject, typeField.Type, valueField)
			if err != nil {
				return
			}
			continue
		}
		name := typeField.Name
		jsonName := typeField.Tag.Get("json")
		if jsonName == "" {
			jsonName = name
		}
		subField, have := jsonObject.(map[string]interface{})[jsonName]
		if !have {
			continue
		}
		err = jsonObjectToObjectByReflect(subField, typeField.Type, valueField)
		if err != nil {
			return
		}
	}
	return
}

func jsonObjectToSlice(jsonObject interface{}, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Slice || reflect.ValueOf(jsonObject).Kind() != reflect.Slice {
		err = ErrorMaker.GetError(fmt.Sprintf("不能使用ReflectJsonFieldToSlice方法将%s映射到%v中", reflect.ValueOf(jsonObject).Kind(), waitSetType))
		return
	}
	subFieldType := waitSetType.Elem()
	tempArr := reflect.MakeSlice(reflect.SliceOf(subFieldType), len(jsonObject.([]interface{})), len(jsonObject.([]interface{})))
	for index, field := range jsonObject.([]interface{}) {
		err = jsonObjectToObjectByReflect(field, subFieldType, tempArr.Index(index))
	}
	waitSetValue.Set(tempArr)
	return
}

func jsonObjectToBaseType(jsonObject interface{}, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	switch waitSetType.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		switch reflect.TypeOf(jsonObject).Kind() {
		case reflect.Uint:
			waitSetValue.SetUint(uint64(jsonObject.(uint)))
		case reflect.Uint8:
			waitSetValue.SetUint(uint64(jsonObject.(uint8)))
		case reflect.Uint16:
			waitSetValue.SetUint(uint64(jsonObject.(uint16)))
		case reflect.Uint32:
			waitSetValue.SetUint(uint64(jsonObject.(uint32)))
		case reflect.Uint64:
			waitSetValue.SetUint(jsonObject.(uint64))
		case reflect.Uintptr:
			waitSetValue.SetUint(uint64(jsonObject.(uintptr)))
		default:
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", reflect.TypeOf(jsonObject), waitSetType))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch reflect.TypeOf(jsonObject).Kind() {
		case reflect.Int:
			waitSetValue.SetInt(int64(jsonObject.(int)))
		case reflect.Int8:
			waitSetValue.SetInt(int64(jsonObject.(int8)))
		case reflect.Int16:
			waitSetValue.SetInt(int64(jsonObject.(int16)))
		case reflect.Int32:
			waitSetValue.SetInt(int64(jsonObject.(int32)))
		case reflect.Int64:
			waitSetValue.SetInt(jsonObject.(int64))
		case reflect.Float32:
			waitSetValue.SetInt(int64(jsonObject.(float32)))
		case reflect.Float64:
			waitSetValue.SetInt(int64(jsonObject.(float64)))
		default:
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", reflect.TypeOf(jsonObject), waitSetType))
		}
	case reflect.Float32, reflect.Float64:
		switch reflect.TypeOf(jsonObject).Kind() {
		case reflect.Float32:
			waitSetValue.SetFloat(float64(jsonObject.(float32)))
		case reflect.Float64:
			waitSetValue.SetFloat(jsonObject.(float64))
		default:
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", reflect.TypeOf(jsonObject), waitSetType))
		}
	case reflect.String:
		if reflect.TypeOf(jsonObject).Kind() != reflect.String {
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", reflect.TypeOf(jsonObject), waitSetType))
			return
		}
		waitSetValue.SetString(jsonObject.(string))
	case reflect.Bool:
		if reflect.TypeOf(jsonObject).Kind() != reflect.Bool {
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", reflect.TypeOf(jsonObject), waitSetType))
			return
		}
		waitSetValue.SetBool(jsonObject.(bool))
	default:
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", reflect.TypeOf(jsonObject), waitSetType))
	}
	return
}
