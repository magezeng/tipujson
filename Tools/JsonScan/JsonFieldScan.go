package JsonScan

import (
	"errors"
	. "github.com/magezeng/TipuJson/Modles"
	"strings"
)

func GetJsonListField(expression *JsonExpression) (lastExpression *JsonExpression, field *JsonField, err error) {
	field = new(JsonField)
	field.Type = JsonFieldTypeList
	contents := []*JsonField{}
	lastExpression = expression
	lastExpression = lastExpression.GetNext()
	if err != nil {
		return
	}
	for true {

		var tempField *JsonField
		switch lastExpression.Type {
		case JsonExpressionTypeListEnd:
			//list结束了
			lastExpression = lastExpression.GetNext()
			field.Content = contents
			return
		case JsonExpressionTypeMapStart:
			lastExpression, tempField, err = GetJsonMapField(lastExpression)
		case JsonExpressionTypeListStart:
			lastExpression, tempField, err = GetJsonListField(lastExpression)
		case JsonExpressionTypeStringMark:
			lastExpression, tempField, err = GetJsonStringField(lastExpression)
		case JsonExpressionTypeValue:
			lastExpression, tempField, err = GetJsonNumberField(lastExpression)
		default:
			err = errors.New("数组内只允许存放Number，Map，List，String")
		}
		if err != nil {
			return
		}
		tempField.Parents = field
		contents = append(contents, tempField)
	}
	return
}

func GetJsonMapField(expression *JsonExpression) (lastExpression *JsonExpression, field *JsonField, err error) {
	field = new(JsonField)
	field.Type = JsonFieldTypeMap
	contents := map[string]*JsonField{}
	lastExpression = expression
	lastExpression = lastExpression.GetNext()
	if err != nil {
		return
	}
	for true {
		if lastExpression.Type == JsonExpressionTypeMapEnd {
			//map结束了
			lastExpression = lastExpression.GetNext()
			field.Content = contents
			return
		}
		if lastExpression.Type != JsonExpressionTypeStringMark {
			err = errors.New("Map的键必须是字符串")
			return
		}
		var key string
		var tempField *JsonField
		lastExpression, key, tempField, err = GetJsonKeyValue(lastExpression)
		if err != nil {
			return
		}
		tempField.Parents = field
		contents[key] = tempField
	}
	return
}

func GetJsonKeyValue(expression *JsonExpression) (lastExpression *JsonExpression, key string, field *JsonField, err error) {
	lastExpression, key, err = GetJsonStringContent(expression)
	if err != nil {
		return
	}
	if lastExpression.Type != JsonExpressionTypeColon {
		err = errors.New("Json字典内Key后面必须跟随一个':'")
		return
	}
	lastExpression = lastExpression.GetNext()
	if err != nil {
		return
	}
	switch lastExpression.Type {
	case JsonExpressionTypeStringMark:
		lastExpression, field, err = GetJsonStringField(lastExpression)
	case JsonExpressionTypeValue:
		lastExpression, field, err = GetJsonNumberField(lastExpression)
	case JsonExpressionTypeListStart:
		lastExpression, field, err = GetJsonListField(lastExpression)
	case JsonExpressionTypeMapStart:
		lastExpression, field, err = GetJsonMapField(lastExpression)
	default:
		err = errors.New("键值对的值只能是数字或者字符串")
		return
	}
	if lastExpression.Type == JsonExpressionTypeComma {
		lastExpression = lastExpression.GetNext()
	}
	return
}
func GetJsonNumberField(expression *JsonExpression) (lastExpression *JsonExpression, field *JsonField, err error) {
	field = new(JsonField)
	field.Type = JsonFieldTypeNumber
	lastExpression = expression
	field.Content = lastExpression.Content
	lastExpression = lastExpression.Next

	tempContent := strings.ToLower(lastExpression.Content)
	if tempContent == "true" || tempContent == "false" {
		field.Type = JsonFieldTypeBool
	}
	return
}
func GetJsonStringField(expression *JsonExpression) (lastExpression *JsonExpression, field *JsonField, err error) {
	field = new(JsonField)
	field.Type = JsonFieldTypeString
	lastExpression, field.Content, err = GetJsonStringContent(expression)
	return
}

func GetJsonStringContent(expression *JsonExpression) (lastExpression *JsonExpression, content string, err error) {
	lastExpression = expression
	lastExpression = lastExpression.Next
	if err != nil {
		return
	}
	content = lastExpression.Content
	lastExpression = lastExpression.Next
	if err != nil {
		return
	}
	if lastExpression.Type != JsonExpressionTypeStringMark {
		err = errors.New("字符串未以\"结束")
		return
	}
	lastExpression = lastExpression.Next
	return
}
