package tipujson

import (
	"errors"
	"fmt"
	"reflect"
	. "tipujson/Modles"
	"tipujson/Utils"
)

func jsonFieldToObject(field *JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if field.Type == JsonFieldTypeNull {
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
		err = jsonFieldToObject(field, waitSetType, waitSetValue)
	case reflect.Struct:
		err = jsonFieldToStruct(field, waitSetType, waitSetValue)
	case reflect.Slice:
		err = jsonFieldToSlice(field, waitSetType, waitSetValue)
	case reflect.Interface:
		err = jsonFieldToInterface(field, waitSetType, waitSetValue)
	case reflect.Map:
		err = jsonFieldToMap(field, waitSetType, waitSetValue)
	default:
		err = jsonFieldToBaseType(field, waitSetType, waitSetValue)
	}
	return
}

func jsonFieldToInterface(field *JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	// 假如说双方类型未匹配上，则直接出错退出
	if waitSetType.Kind() != reflect.Interface {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToInterface方法将%s映射到%v中", field.Type, waitSetType))
		return
	}
	var tempInterface interface{}
	tempInterface, err = jsonFieldToJsonObject(field)
	waitSetValue.Set(reflect.ValueOf(tempInterface))
	return
}

func jsonFieldToMap(field *JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	// 假如说双方类型未匹配上，则直接出错退出
	if field.Type != JsonFieldTypeMap || waitSetType.Kind() != reflect.Map {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToMap方法将%s映射到%v中", field.Type, waitSetType))
		return
	}
	var tempInterface interface{}
	tempInterface, err = jsonFieldToJsonObject(field)
	if waitSetValue.IsNil() {
		waitSetValue.Set(reflect.MakeMap(waitSetType))
	}
	for key, value := range tempInterface.(map[string]interface{}) {
		waitSetValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	}
	return
}

func jsonFieldToStruct(field *JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Struct || field.Type != JsonFieldTypeMap {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToStruct方法将%s映射到%v中", field.Type, waitSetType))
		return
	}
	content := field.Content.(map[string]*JsonField)
	for i := 0; i < waitSetType.NumField(); i++ {
		valueField := waitSetValue.Field(i)
		typeField := waitSetType.Field(i)
		if typeField.Anonymous {
			err = jsonFieldToObject(field, typeField.Type, valueField)
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
		subField, have := content[jsonName]
		if !have {
			continue
		}
		err = jsonFieldToObject(subField, typeField.Type, valueField)
		if err != nil {
			return
		}
	}
	return
}

func jsonFieldToSlice(field *JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Slice || field.Type != JsonFieldTypeList {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToSlice方法将%s映射到%v中", field.Type, waitSetType))
		return
	}
	subFieldType := waitSetType.Elem()
	fields := field.Content.([]*JsonField)
	tempArr := reflect.MakeSlice(reflect.SliceOf(subFieldType), len(fields), len(fields))
	for index, field := range fields {
		err = jsonFieldToObject(field, subFieldType, tempArr.Index(index))
	}
	waitSetValue.Set(tempArr)
	return
}

func jsonFieldToJsonObject(field *JsonField) (result interface{}, err error) {
	switch field.Type {
	case JsonFieldTypeString:
		result = field.Content
		return
	case JsonFieldTypeNumber:
		result, err = Utils.StringToFloat64(field.Content.(string))
		return
	case JsonFieldTypeBool:
		result, err = Utils.StringToBool(field.Content.(string))
		return
	case JsonFieldTypeList:
		content := field.Content.([]*JsonField)
		length := len(content)
		tempResult := make([]interface{}, length)
		for index, value := range content {
			tempResult[index], err = jsonFieldToJsonObject(value)
			if err != nil {
				return
			}
		}
		result = tempResult
		return
	case JsonFieldTypeMap:
		tempResult := map[string]interface{}{}
		content := field.Content.(map[string]*JsonField)
		for key, value := range content {
			tempResult[key], err = jsonFieldToJsonObject(value)
			if err != nil {
				return
			}
		}
		result = tempResult
		return
	case JsonFieldTypeNull:
		result = nil
		return
	default:
		//不可能有其他类型出现
		panic(errors.New(fmt.Sprintf("JsonField.ToMapElement出现了不支持的类型%s", field.Type)))
		return
	}
}

func jsonFieldToBaseType(field *JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	switch waitSetType.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if field.Type != JsonFieldTypeNumber {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%v中", field.Type, waitSetType))
			return
		}
		var tempValue uint64
		tempValue, err = Utils.StringToUInt64(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetUint(tempValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Type != JsonFieldTypeNumber {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%v中", field.Type, waitSetType))
			return
		}
		var tempValue int64
		tempValue, err = Utils.StringToInt64(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetInt(tempValue)
	case reflect.Float32, reflect.Float64:
		if field.Type != JsonFieldTypeNumber {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%v中", field.Type, waitSetType))
			return
		}
		var tempValue float64
		tempValue, err = Utils.StringToFloat64(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetFloat(tempValue)
	case reflect.String:
		if field.Type != JsonFieldTypeString {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%v中", field.Type, waitSetType))
			return
		}
		waitSetValue.SetString(field.Content.(string))
	case reflect.Bool:
		if field.Type != JsonFieldTypeBool {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%v中", field.Type, waitSetType))
			return
		}
		var tempValue bool
		tempValue, err = Utils.StringToBool(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetBool(tempValue)
	default:
		err = errors.New(fmt.Sprintf("暂不支持%v类型字段填充", waitSetType))
	}
	return
}
