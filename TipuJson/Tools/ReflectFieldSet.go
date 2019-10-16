package Tools

import (
	"TipuJson/TipuJson/Modles"
	"TipuJson/TipuJson/Tools/StringToAny"
	"errors"
	"fmt"
	"reflect"
)

func ReflectJsonFieldToAnyType(field *Modles.JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	switch waitSetType.Kind() {
	case reflect.Struct:
		err = ReflectJsonFieldToStruct(field, waitSetType, waitSetValue)
	case reflect.Slice:
		err = ReflectJsonFieldToSlice(field, waitSetType, waitSetValue)
	case reflect.Interface:
		err = ReflectJsonFieldToInterface(field, waitSetType, waitSetValue)
	case reflect.Map:
		err = ReflectJsonFieldToMap(field, waitSetType, waitSetValue)
	default:
		err = ReflectJsonFieldToBaseType(field, waitSetType, waitSetValue)
	}
	return
}
func ReflectJsonFieldToInterface(field *Modles.JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	// 假如说双方类型未匹配上，则直接出错退出
	if waitSetType.Kind() != reflect.Interface {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToInterface方法将%s映射到%V中", field.Type, waitSetType))
		return
	}
	var tempInterface interface{}
	tempInterface, err = field.ToMapElement()
	waitSetValue.Set(reflect.ValueOf(tempInterface))
	return
}

func ReflectJsonFieldToMap(field *Modles.JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	// 假如说双方类型未匹配上，则直接出错退出
	if field.Type != Modles.JsonFieldTypeMap || waitSetType.Kind() != reflect.Map {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToMap方法将%s映射到%V中", field.Type, waitSetType))
		return
	}
	var tempInterface interface{}
	tempInterface, err = field.ToMapElement()
	if waitSetValue.IsNil() {
		waitSetValue.Set(reflect.MakeMap(waitSetType))
	}
	for key, value := range tempInterface.(map[string]interface{}) {
		waitSetValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	}
	return
}

func ReflectJsonFieldToStruct(field *Modles.JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Struct || field.Type != Modles.JsonFieldTypeMap {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToStruct方法将%s映射到%V中", field.Type, waitSetType))
		return
	}
	content := field.Content.(map[string]*Modles.JsonField)
	for i := 0; i < waitSetType.NumField(); i++ {
		valueField := waitSetValue.Field(i)
		typeField := waitSetType.Field(i)
		name := typeField.Name
		jsonName := typeField.Tag.Get("json")
		fmt.Printf("%V", reflect.TypeOf(reflect.ValueOf(valueField)))
		if jsonName == "" {
			jsonName = name
		}
		subField, have := content[jsonName]
		if !have {
			continue
		}
		err = ReflectJsonFieldToAnyType(subField, typeField.Type, valueField)
		if err != nil {
			return
		}
	}
	return
}

func ReflectJsonFieldToSlice(field *Modles.JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Slice || field.Type != Modles.JsonFieldTypeList {
		err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToSlice方法将%s映射到%V中", field.Type, waitSetType))
		return
	}
	subFieldType := waitSetType.Elem()
	fields := field.Content.([]*Modles.JsonField)
	tempArr := reflect.MakeSlice(reflect.SliceOf(subFieldType), len(fields), len(fields))
	for index, field := range fields {
		err = ReflectJsonFieldToAnyType(field, subFieldType, tempArr.Index(index))
	}
	waitSetValue.Set(tempArr)
	return
}

func ReflectJsonFieldToBaseType(field *Modles.JsonField, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	switch waitSetType.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if field.Type != Modles.JsonFieldTypeNumber {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%V中", field.Type, waitSetType))
			return
		}
		var tempValue uint64
		tempValue, err = StringToAny.StringToUInt64(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetUint(tempValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Type != Modles.JsonFieldTypeNumber {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%V中", field.Type, waitSetType))
			return
		}
		var tempValue int64
		tempValue, err = StringToAny.StringToInt64(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetInt(tempValue)
	case reflect.String:
		if field.Type != Modles.JsonFieldTypeString {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%V中", field.Type, waitSetType))
			return
		}
		waitSetValue.SetString(field.Content.(string))
	case reflect.Bool:
		if field.Type != Modles.JsonFieldTypeBool {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%V中", field.Type, waitSetType))
			return
		}
		var tempValue bool
		tempValue, err = StringToAny.StringToBool(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetBool(tempValue)
	case reflect.Float32, reflect.Float64:
		if field.Type != Modles.JsonFieldTypeNumber {
			err = errors.New(fmt.Sprintf("不能使用ReflectJsonFieldToAny方法将%s映射到%V中", field.Type, waitSetType))
			return
		}
		var tempValue float64
		tempValue, err = StringToAny.StringToFloat64(field.Content.(string))
		if err != nil {
			return
		}
		waitSetValue.SetFloat(tempValue)
	default:
		err = errors.New(fmt.Sprintf("暂不支持%V类型字段填充", waitSetType))
	}
	return
}
