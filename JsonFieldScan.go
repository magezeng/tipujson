package TipuJson

import (
	"errors"
	"github.com/magezeng/TipuJson/BytesScanner"
	. "github.com/magezeng/TipuJson/Modles"
)

func GetJsonFieldFromString(jsonString string) (field *JsonField, err error) {
	defer func() {
		if err == nil {
			if tempErr := recover(); tempErr != nil {
				err = tempErr.(error)
			}
		}
	}()
	bytes := []byte(jsonString)
	scanner := BytesScanner.BytesScanner{bytes, 0}
	field = getJsonObjectFieldFromScanner(&scanner)
	return
}

func getJsonObjectFieldFromScanner(scanner *BytesScanner.BytesScanner) (result *JsonField) {
	scanner.BackMoveToNotNull()
	switch scanner.CurrentValue() {
	case '[':
		return getJsonListFieldFromScanner(scanner)
	case '{':
		return getJsonMapFieldFromScanner(scanner)
	case '"':
		defer scanErrorDescriptionDefer(scanner, scanner.Cursor, "未扫描到一个完整的字符串")
		return &JsonField{
			Type:    JsonFieldTypeString,
			Content: scanner.ScanString(),
		}
	default:
		defer scanErrorDescriptionDefer(scanner, scanner.Cursor, "未扫描到一个完整的数值字符串")
		result, isBool := scanner.ScanNumberString()
		if isBool {
			return &JsonField{
				Type:    JsonFieldTypeBool,
				Content: result,
			}
		} else {
			return &JsonField{
				Type:    JsonFieldTypeNumber,
				Content: result,
			}
		}
	}
}

func getJsonListFieldFromScanner(scanner *BytesScanner.BytesScanner) (field *JsonField) {
	content := []*JsonField{}
	scanner.BackMove()
	scanner.BackMoveToNotNull()
	for {
		value, isEnd := getListValueOrListEndFromScanner(scanner)
		if isEnd {
			break
		}
		content = append(content, value)
	}
	return &JsonField{
		Type:    JsonFieldTypeList,
		Content: content,
	}
}

func getListValueOrListEndFromScanner(scanner *BytesScanner.BytesScanner) (value *JsonField, isEnd bool) {
	scanner.BackMoveToNotNull()
	if scanner.CurrentValue() == ']' {
		isEnd = true
		scanner.BackMove()
		return
	}
	value = getJsonObjectFieldFromScanner(scanner)
	isEnd = false
	if scanner.CurrentValue() == ',' {
		scanner.BackMove()
	}
	return
}

func getJsonMapFieldFromScanner(scanner *BytesScanner.BytesScanner) (field *JsonField) {
	content := map[string]*JsonField{}
	scanner.BackMove()
	for {
		key, value, isEnd := getMapKeyValueOrMapEndFromScanner(scanner)
		if isEnd {
			break
		}
		content[key] = value
	}
	return &JsonField{
		Type:    JsonFieldTypeMap,
		Content: content,
	}
}

func getMapKeyValueOrMapEndFromScanner(scanner *BytesScanner.BytesScanner) (key string, value *JsonField, isEnd bool) {
	defer scanErrorDescriptionDefer(scanner, scanner.Cursor, "未找到一个正常的键值对")
	scanner.BackMoveToNotNull()
	if scanner.CurrentValue() == '}' {
		isEnd = true
		scanner.BackMove()
		return
	} else if scanner.CurrentValue() == '"' {
		isEnd = false
		key = scanner.ScanString()
		scanner.BackMoveToNotNull()
		if scanner.CurrentValue() != ':' {
			panic(errors.New(""))
		}
		scanner.BackMove()
		scanner.BackMoveToNotNull()
		value = getJsonObjectFieldFromScanner(scanner)
		if scanner.CurrentValue() == ',' {
			scanner.BackMove()
		}
		return
	} else {
		panic(errors.New(""))
	}
}

func scanErrorDescriptionDefer(scanner *BytesScanner.BytesScanner, position int, errReason string) {
	if err := recover(); err != nil {
		tempErr, isError := err.(error)
		if isError && len(tempErr.Error()) > 0 {
			panic(tempErr)
		} else {
			panic(
				errors.New(
					"从:" +
						scanner.GetMarkString(position, "<--该位置-->") +
						"  " + errReason,
				),
			)
		}
	}
}
