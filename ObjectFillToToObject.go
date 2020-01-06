package tipujson

import (
	"fmt"
	"github.com/magezeng/tipujson/ErrorMaker"
	"reflect"
)

//在遍历
type subType struct {
	waitSetType  reflect.StructField
	waitSetValue reflect.Value
	index        int
}
type SliceHandler func(positions []string, fromSlice interface{}, directionSlice interface{}) (interface{}, bool)

//fromObject 源数据,directionObject目标数据,sliceHandler Slice处理函数(position为Slice对应所在源数据内的位置,fromSlice该数组的源数据，directionSlice该数组的目标数据)
//递归源数据内所有字段,一般情况下(字段类型不是Slice) 非Zero的字段覆盖掉目标数据内对应位置的字段
//该值类型为Slice时分几种情况特殊处理,情况分别如下:
//1.Slice处理函数为nil时:按照一般情况进行处理
//2.Slice处理函数不为nil时: 调用处理函数 得到两个变量分别为  覆盖结果和生效与否，如果该次处理结果为不生效，则该字段按照一般情况处理，否则将返回结果直接覆盖到目标字段
func objectFillToObject(fromObject interface{}, directionObject interface{}, sliceHandler SliceHandler) (err error) {
	directionType := reflect.TypeOf(directionObject)
	directionValue := reflect.ValueOf(directionObject)
	fromType := reflect.TypeOf(fromObject)
	fromValue := reflect.ValueOf(fromObject)
	return objectFillToObjectByReflect(fromType, fromValue, directionType, directionValue, sliceHandler)
}

func objectFillToObjectByReflect(fromType reflect.Type, fromValue reflect.Value, waitSetType reflect.Type, waitSetValue reflect.Value, sliceHandler SliceHandler) (err error) {
	return objectFillToObjectByReflectWithPositions(fromType, fromValue, waitSetType, waitSetValue, []string{}, sliceHandler)
}

func objectFillToObjectByReflectWithPositions(fromType reflect.Type, fromValue reflect.Value, directionType reflect.Type, directionValue reflect.Value, prePositions []string, sliceHandler SliceHandler) (err error) {
	if fromValue.IsZero() {
		//遇到空值的情况直接不进行映射
		return
	}
	//from必须是具体的值才进行下一轮   Interface和Ptr都代表from还未指向具体的值  所以不断遍历   直到from为值为止
	for fromType.Kind() == reflect.Interface || fromType.Kind() == reflect.Ptr {
		if fromType.Kind() == reflect.Interface {
			fromType = fromValue.Type()
		} else {
			fromType = fromType.Elem()
			fromValue = fromValue.Elem()
		}
		if fromValue.IsZero() {
			//遇到空值的情况直接不进行映射
			return
		}
	}
	switch directionType.Kind() {
	case reflect.Ptr:
		directionType = directionType.Elem()
		if directionValue.IsNil() {
			directionValue.Set(reflect.New(directionType))
		}
		directionValue = directionValue.Elem()
		err = objectFillToObjectByReflectWithPositions(fromType, fromValue, directionType, directionValue, prePositions, sliceHandler)
	case reflect.Struct:
		err = ObjectFillToStruct(fromType, fromValue, directionType, directionValue, prePositions, sliceHandler)
	case reflect.Slice:
		//数组直接进行覆盖
		err = ObjectFillToSlice(fromType, fromValue, directionType, directionValue, prePositions, sliceHandler)
	case reflect.Interface:
		//interface直接进行覆盖
		directionValue.Set(fromValue)
	case reflect.Map:
		err = ObjectFillToMap(fromType, fromValue, directionType, directionValue, prePositions, sliceHandler)
	default:
		err = ObjectFillToBaseType(fromType, fromValue, directionType, directionValue)
	}
	return
}
func ObjectFillToSlice(fromType reflect.Type, fromValue reflect.Value, directionType reflect.Type, directionValue reflect.Value, prePositions []string, sliceHandler SliceHandler) (err error) {
	if fromType.Kind() != reflect.Slice || directionType.Kind() != reflect.Slice {
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
		return
	}
	resultSlice, effective := sliceHandler(prePositions, fromValue.Interface(), directionValue.Interface())
	if !effective {
		tempDirectionValue := reflect.MakeSlice(directionType, fromValue.Len(), fromValue.Len())
		for i := 0; i < fromValue.Len(); i++ {
			err = objectFillToObjectByReflectWithPositions(fromType.Elem(), fromValue.Index(i), directionType.Elem(), tempDirectionValue.Index(i), prePositions, sliceHandler)
			if err != nil {
				return
			}
		}
		directionValue.Set(tempDirectionValue)
		return
	}
	switch reflect.TypeOf(resultSlice) {
	case nil:
		directionValue.Set(reflect.New(directionType))
	case directionType:
		directionValue.Set(reflect.ValueOf(resultSlice))
	default:
		err = ErrorMaker.GetError("Slice处理函数返回的类型不对")
	}
	return
}

func ObjectFillToStruct(fromType reflect.Type, fromValue reflect.Value, directionType reflect.Type, directionValue reflect.Value, prePositions []string, sliceHandler SliceHandler) (err error) {

	//类型不匹配不进行映射  直接返回错误
	if fromType.Kind() != reflect.Map && fromType.Kind() != reflect.Struct {
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
		return
	}

	//将目标转换为Map  方便判断是否存在Key
	directionSubTypeMap := map[string]subType{}
	for i := 0; i < directionType.NumField(); i++ {
		directionSubTypeMap[directionType.Field(i).Name] = subType{directionType.Field(i), directionValue.Field(i), i}
		jsonName := directionType.Field(i).Tag.Get("json")
		if jsonName != "" {
			directionSubTypeMap[jsonName] = subType{directionType.Field(i), directionValue.Field(i), i}
		}
	}
	//遍历数据源对象,一一将数据映射存入到目标对象内
	switch fromType.Kind() {
	//fromType只可能是Map或者Struct   其它类型在本函数开始的位置已经被排除掉了
	case reflect.Map:
		//遍历字典   对对应的字段进行赋值
		for _, key := range fromValue.MapKeys() {
			keyName, isString := key.Interface().(string)
			if !isString {
				continue
			}
			subDirection, directionHaveKey := directionSubTypeMap[keyName]
			if !directionHaveKey {
				continue
			}
			err = objectFillToObjectByReflectWithPositions(
				fromType.Elem(), fromValue.MapIndex(key).Elem(),
				subDirection.waitSetType.Type, subDirection.waitSetValue,
				append(prePositions, keyName), sliceHandler,
			)
			if err != nil {
				return
			}
		}
	case reflect.Struct:
		//对源结构体进行遍历，对对应字段进行赋值
		for i := 0; i < fromType.NumField(); i++ {
			fromValueField := fromValue.Field(i)
			fromTypeField := fromType.Field(i)
			if fromTypeField.Anonymous {
				err = ObjectFillToStruct(fromTypeField.Type, fromValueField, directionType, directionValue, prePositions, sliceHandler)
				if err != nil {
					return
				}
				continue
			}
			keyName := fromTypeField.Name
			subDirection, directionHaveKey := directionSubTypeMap[keyName]
			if !directionHaveKey {
				continue
			}
			err = objectFillToObjectByReflectWithPositions(
				fromTypeField.Type, fromValueField,
				subDirection.waitSetType.Type, subDirection.waitSetValue,
				append(prePositions, keyName), sliceHandler,
			)
			if err != nil {
				return
			}
		}
	}
	return
}

/**/
func ObjectFillToMap(fromType reflect.Type, fromValue reflect.Value, directionType reflect.Type, directionValue reflect.Value, prePositions []string, sliceHandler SliceHandler) (err error) {
	//类型不匹配不进行映射  直接返回错误
	if fromType.Kind() != reflect.Map && fromType.Kind() != reflect.Struct {
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
		return
	}
	if directionValue.IsNil() {
		directionValue.Set(reflect.MakeMap(directionType))
	}
	switch directionType.Elem().Kind() {
	case reflect.Interface:
		var tempMap interface{}
		tempMap, err = objectToJsonObjectByReflect(fromType, fromValue)
		if err != nil {
			return
		}
		directionValue.Set(reflect.ValueOf(tempMap))
	default:
		subDirectionType := directionType.Elem()
		switch fromType.Kind() {
		case reflect.Map:
			for _, key := range fromValue.MapKeys() {
				keyName, isString := key.Interface().(string)
				if !isString {
					continue
				}
				subDirectionValue := reflect.New(subDirectionType)
				err = objectFillToObjectByReflectWithPositions(
					fromValue.MapIndex(key).Type(), fromValue.MapIndex(key).Elem(),
					subDirectionType, subDirectionValue,
					append(prePositions, keyName), sliceHandler,
				)
				if err != nil {
					return
				}
				directionValue.SetMapIndex(key, subDirectionValue)
			}
		case reflect.Struct:
			for i := 0; i < fromType.NumField(); i++ {
				fromValueField := fromValue.Field(i)
				fromTypeField := fromType.Field(i)
				if fromTypeField.Anonymous {
					err = ObjectFillToStruct(fromTypeField.Type, fromValueField, directionType, directionValue, prePositions, sliceHandler)
					if err != nil {
						return
					}
					continue
				}
				keyName := fromTypeField.Name
				subDirectionValue := reflect.New(subDirectionType)
				err = objectFillToObjectByReflectWithPositions(
					fromTypeField.Type, fromValueField,
					subDirectionType, subDirectionValue,
					append(prePositions, keyName), sliceHandler,
				)
				if err != nil {
					return
				}
				directionValue.SetMapIndex(reflect.ValueOf(keyName), subDirectionValue)
			}
		}
	}
	return
}

func ObjectFillToBaseType(fromType reflect.Type, fromValue reflect.Value, directionType reflect.Type, directionValue reflect.Value) (err error) {
	switch directionType.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		switch fromType.Kind() {
		case reflect.Uint:
			directionValue.SetUint(uint64(fromValue.Interface().(uint)))
		case reflect.Uint8:
			directionValue.SetUint(uint64(fromValue.Interface().(uint8)))
		case reflect.Uint16:
			directionValue.SetUint(uint64(fromValue.Interface().(uint16)))
		case reflect.Uint32:
			directionValue.SetUint(uint64(fromValue.Interface().(uint32)))
		case reflect.Uint64:
			directionValue.SetUint(fromValue.Interface().(uint64))
		case reflect.Uintptr:
			directionValue.SetUint(uint64(fromValue.Interface().(uintptr)))
		default:
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch fromType.Kind() {
		case reflect.Int:
			directionValue.SetInt(int64(fromValue.Interface().(int)))
		case reflect.Int8:
			directionValue.SetInt(int64(fromValue.Interface().(int8)))
		case reflect.Int16:
			directionValue.SetInt(int64(fromValue.Interface().(int16)))
		case reflect.Int32:
			directionValue.SetInt(int64(fromValue.Interface().(int32)))
		case reflect.Int64:
			directionValue.SetInt(fromValue.Interface().(int64))
		case reflect.Float32:
			directionValue.SetInt(int64(fromValue.Interface().(float32)))
		case reflect.Float64:
			directionValue.SetInt(int64(fromValue.Interface().(float64)))
		default:
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
		}
	case reflect.Float32, reflect.Float64:
		switch fromType.Kind() {
		case reflect.Float32:
			directionValue.SetFloat(float64(fromValue.Interface().(float32)))
		case reflect.Float64:
			directionValue.SetFloat(fromValue.Interface().(float64))
		default:
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
		}
	case reflect.String:
		if fromType.Kind() != reflect.String {
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
			return
		}
		directionValue.SetString(fromValue.Interface().(string))
	case reflect.Bool:
		if fromType.Kind() != reflect.Bool {
			err = ErrorMaker.GetError(fmt.Sprintf("不能将%v映射到%v中", fromType, directionType))
			return
		}
		directionValue.SetBool(fromValue.Interface().(bool))
	default:
		err = ErrorMaker.GetError(fmt.Sprintf("不能将%s映射到%v中", fromType, directionType))
	}
	return
}
