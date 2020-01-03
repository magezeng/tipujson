package tipujson

import (
	"fmt"
	"github.com/magezeng/tipujson/ErrorMaker"
	"reflect"
)

//fromObject 源数据,directionObject目标数据,sliceHandler Slice处理函数(position为Slice对应所在源数据内的位置,fromSlice该数组的源数据，directionSlice该数组的目标数据)
//递归源数据内所有字段,一般情况下(字段类型不是Slice) 非Zero的字段覆盖掉目标数据内对应位置的字段
//该值类型为Slice时分几种情况特殊处理,情况分别如下:
//1.Slice处理函数为nil时:按照一般情况进行处理
//2.Slice处理函数不为nil时: 调用处理函数 得到两个变量分别为  覆盖结果和生效与否，如果该次处理结果为不生效，则该字段按照一般情况处理，否则将返回结果直接覆盖到目标字段
func ObjectFillToObject(fromObject interface{}, directionObject interface{}, sliceHandler func(position []string, fromSlice interface{}, directionSlice interface{}) (interface{}, bool)) (err error) {

	directionType := reflect.TypeOf(directionObject)
	directionValue := reflect.ValueOf(directionObject)
	fromType := reflect.TypeOf(fromObject)
	fromValue := reflect.ValueOf(fromObject)
	return ObjectFillToObjectByReflect(fromType, fromValue, directionType, directionValue, sliceHandler)
}

func ObjectFillToObjectByReflect(fromType reflect.Type, fromValue reflect.Value, waitSetType reflect.Type, waitSetValue reflect.Value, sliceHandler func(position []string, fromSlice interface{}, directionSlice interface{}) (interface{}, bool)) (err error) {
	return objectFillToObjectByReflectWithPositions(fromType, fromValue, waitSetType, waitSetValue, []string{}, sliceHandler)
}

func objectFillToObjectByReflectWithPositions(fromType reflect.Type, fromValue reflect.Value, waitSetType reflect.Type, waitSetValue reflect.Value, positions []string, sliceHandler func(position []string, fromSlice interface{}, directionSlice interface{}) (interface{}, bool)) (err error) {
	if fromValue.IsZero() {
		//遇到空值的情况直接不进行映射
		return
	}
	//from必须是具体的值才进行下一轮   Interface和Ptr都代表from还未指向具体的值  所以不断遍历   直到from为值为止
	for fromType.Kind() == reflect.Interface || fromType.Kind() == reflect.Ptr {
		if fromType.Kind() == reflect.Interface {
			fromType = fromValue.Elem().Type()
			fromValue = fromValue.Elem()
		} else {
			fromType = fromType.Elem()
			fromValue = fromValue.Elem()
		}
		if fromValue.IsZero() {
			//遇到空值的情况直接不进行映射
			return
		}
	}
	switch waitSetType.Kind() {
	case reflect.Ptr:
		waitSetType = waitSetType.Elem()
		if waitSetValue.IsNil() {
			waitSetValue.Set(reflect.New(waitSetType))
		}
		waitSetValue = waitSetValue.Elem()
		err = objectFillToObjectByReflectWithPositions(fromType, fromValue, waitSetType, waitSetValue, append(positions, waitSetType.Name()), sliceHandler)
	case reflect.Struct:
		err = ObjectFillToStruct(fromType, fromValue, waitSetType, waitSetValue)
	case reflect.Slice:
		//数组直接进行覆盖
		waitSetValue.Set(fromValue)
	case reflect.Interface:
		//interface直接进行覆盖
		waitSetValue.Set(fromValue)
	case reflect.Map:
		err = ObjectFillToMap(fromType, fromValue, waitSetType, waitSetValue)
	default:
		err = ObjectFillToBaseType(fromType, fromValue, waitSetType, waitSetValue)
	}
	return
}

func ObjectFillToStruct(fromType reflect.Type, fromValue reflect.Value, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	if waitSetType.Kind() != reflect.Struct || (fromType.Kind() != reflect.Map && fromType.Kind() != reflect.Struct) {
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, waitSetType))
		return
	}
	waitSetMap := map[string]struct {
		waitSetType  reflect.Type
		waitSetValue reflect.Value
	}{}
	for i := 0; i < waitSetType.NumField(); i++ {

	}
	for i := 0; i < waitSetType.NumField(); i++ {
		valueField := waitSetValue.Field(i)
		typeField := waitSetType.Field(i)
		if typeField.Anonymous {
			err = objectFillToObjectByReflectWithPositions(jsonObject, typeField.Type, valueField)
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
		err = objectFillToObjectByReflectWithPositions(subField, typeField.Type, valueField)
		if err != nil {
			return
		}
	}
	return
}

func ObjectFillToMap(fromType reflect.Type, fromValue reflect.Value, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
	jsonObjectType := reflect.ValueOf(jsonObject).Kind()
	// 假如说双方类型未匹配上，则直接出错退出
	if jsonObjectType != reflect.Map || waitSetType.Kind() != reflect.Map {
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", jsonObjectType, waitSetType))
		return
	}
	if waitSetValue.IsNil() {
		waitSetValue.Set(reflect.MakeMap(waitSetType))
	}
	for key, value := range jsonObject.(map[string]interface{}) {
		tempValue := reflect.New(waitSetType.Elem())
		err = objectFillToObjectByReflectWithPositions(value, waitSetType.Elem(), tempValue)
		if err != nil {
			return
		}
		waitSetValue.SetMapIndex(reflect.ValueOf(key), tempValue)
	}
	return
}

func ObjectFillToBaseType(fromType reflect.Type, fromValue reflect.Value, waitSetType reflect.Type, waitSetValue reflect.Value) (err error) {
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
