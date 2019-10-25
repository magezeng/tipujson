package TipuJson

import (
	"errors"
	"fmt"
	. "github.com/magezeng/TipuJson/Modles"
)

func ToMapElement(field *JsonField) (result interface{}, err error) {
	switch field.Type {
	case JsonFieldTypeString:
		result = field.Content
		return
	case JsonFieldTypeNumber:
		result, err = StringToFloat64(field.Content.(string))
		return
	case JsonFieldTypeBool:
		result, err = StringToBool(field.Content.(string))
		return
	case JsonFieldTypeList:
		content := field.Content.([]*JsonField)
		length := len(content)
		tempResult := make([]interface{}, length)
		for index, value := range content {
			tempResult[index], err = ToMapElement(value)
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
			tempResult[key], err = ToMapElement(value)
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
