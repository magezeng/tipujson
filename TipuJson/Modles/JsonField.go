package Modles

import (
	"TipuJson/TipuJson/Tools/StringToAny"
	"errors"
	"fmt"
)

type JsonFieldType string

const (
	JsonFieldTypeMap    JsonFieldType = "Map"
	JsonFieldTypeList                 = "List"
	JsonFieldTypeString               = "String"
	JsonFieldTypeNumber               = "Number"
	JsonFieldTypeBool                 = "Bool"
)

type JsonField struct {
	Type    JsonFieldType
	Content interface{} //类型为Map时此字段是Map 类型为List时 此字段是Slice String和Number类型时   直接为值
	Parents *JsonField  // 父对象指针
}

func (field *JsonField) ToMapElement() (result interface{}, err error) {
	switch field.Type {
	case JsonFieldTypeString:
		result = field.Content
		return
	case JsonFieldTypeNumber:
		result, err = StringToAny.StringToFloat64(field.Content.(string))
		return
	case JsonFieldTypeBool:
		result, err = StringToAny.StringToBool(field.Content.(string))
		return
	case JsonFieldTypeList:
		content := field.Content.([]*JsonField)
		length := len(content)
		tempResult := make([]interface{}, length)
		for index, value := range content {
			tempResult[index], err = value.ToMapElement()
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
			tempResult[key], err = value.ToMapElement()
			if err != nil {
				return
			}
		}
		result = tempResult
		return
	default:
		//不可能有其他类型出现
		panic(errors.New(fmt.Sprintf("JsonField.ToMapElement出现了不支持的类型%s", field.Type)))
		return
	}
}
