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

	//对变量进行处理，得到Type和Value
	directionType := reflect.TypeOf(direction).Elem()
	directionValue := reflect.ValueOf(direction).Elem()
	return StringToObjByReflect(src, directionType, directionValue)
}

func StringToObjByReflect(src string, directionType reflect.Type, directionValue reflect.Value) (err error) {
	//首先运算表达式序列
	expression, err := JsonScan.ScanJsonExpressions([]byte(src))
	if err != nil {
		return
	}

	//根据表达书序列  生成持有树
	var srcJsonField *Modles.JsonField
	var lastExpression *Modles.JsonExpression
	kind := expression.Type
	switch kind {
	case Modles.JsonExpressionTypeListStart:
		lastExpression, srcJsonField, err = JsonScan.GetJsonListField(expression)
	case Modles.JsonExpressionTypeMapStart:
		lastExpression, srcJsonField, err = JsonScan.GetJsonMapField(expression)
	default:
		err = errors.New("Json顶层必须是一个字典或者数组")
	}
	if srcJsonField == nil {
		err = errors.New("检索字符串得到的内容为空")
		return
	}
	if lastExpression != nil {
		err = errors.New("形成独立Json时间早于字符串结束(字符串前面一部分已经形成了完整Json)")
		return
	}

	switchDirectionType := directionType
	for switchDirectionType.Kind() == reflect.Ptr {
		switchDirectionType = switchDirectionType.Elem()
	}
	// 分类型进行扫描反射字段,并映射值到目标
	switch switchDirectionType.Kind() {
	case reflect.Struct, reflect.Map, reflect.Interface, reflect.Slice:
		err = Tools.ReflectJsonFieldToAnyType(srcJsonField, directionType, directionValue)
		return
	default:
		err = errors.New(fmt.Sprintf("暂不支持%V的映射", directionType))
	}

	return

}
