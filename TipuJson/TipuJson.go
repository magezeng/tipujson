package TipuJson

import (
	"TipuJson/TipuJson/Modles"
	"TipuJson/TipuJson/Tools"
	"TipuJson/TipuJson/Tools/JsonScan"
	"errors"
	"fmt"
	"reflect"
)

func StringToObj(src string, direction interface{}) (err error) {
	expression, err := JsonScan.ScanJsonExpressions([]byte(src))
	if err != nil {
		return
	}
	var tempField *Modles.JsonField
	var lastExpression *Modles.JsonExpression
	switch expression.Type {
	case Modles.JsonExpressionTypeListStart:
		lastExpression, tempField, err = JsonScan.GetJsonListField(expression)
	case Modles.JsonExpressionTypeMapStart:
		lastExpression, tempField, err = JsonScan.GetJsonMapField(expression)
	default:
		err = errors.New("Json顶层必须是一个字典或者数组")
	}
	if tempField == nil {
		err = errors.New("检索字符串得到的内容为空")
		return
	}
	if lastExpression != nil {
		err = errors.New("形成独立Json时间早于字符串结束(字符串前面一部分已经形成了完整Json)")
		return
	}
	directionType := reflect.TypeOf(direction)
	directionValue := reflect.ValueOf(direction)

	switch directionType.Kind() {
	case reflect.Map, reflect.Struct:
		if tempField.Type != Modles.JsonFieldTypeMap {
			err = errors.New(fmt.Sprintf("不能将字典映射到%V", directionType))
			return
		}
	case reflect.Slice:
		if tempField.Type != Modles.JsonFieldTypeList {
			err = errors.New(fmt.Sprintf("不能将List映射到%V", directionType))
			return
		}
	default:
		err = errors.New(fmt.Sprintf("暂不支持%V的映射", directionType))
	}
	reflectFields, err := Tools.ScanReflectFeild(directionType, directionValue)
	if err != nil {
		return
	}
	fmt.Println(reflectFields)
	return
}

func StringToObjByReflect() {

}
