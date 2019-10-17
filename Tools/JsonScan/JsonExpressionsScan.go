package JsonScan

import (
	"github.com/magezeng/TipuJson/Modles"
)

type JsonScanner struct {
	Content []byte
}

func ScanJsonExpressions(jsonString []byte) (jsonExpressionsHead *Modles.JsonExpression, err error) {
	var cursor = 0
	var jsonExpressionsTail *Modles.JsonExpression = nil
	length := len(jsonString)

	for cursor < length {
		var expression Modles.JsonExpression
		expression, err = getOneExpression(jsonString, length, &cursor)
		if err != nil {
			return
		}
		if jsonExpressionsHead == nil {
			jsonExpressionsHead = &expression
			jsonExpressionsTail = &expression
		} else {
			jsonExpressionsTail.Next = &expression
			expression.Pre = jsonExpressionsTail
			jsonExpressionsTail = &expression
		}
	}
	return
}
func getOneExpression(jsonString []byte, length int, cursor *int) (expression Modles.JsonExpression, err error) {
	CharMap := map[byte]Modles.JsonExpressionType{
		'{': Modles.JsonExpressionTypeMapStart,
		'}': Modles.JsonExpressionTypeMapEnd,
		'[': Modles.JsonExpressionTypeListStart,
		']': Modles.JsonExpressionTypeListEnd,
		'"': Modles.JsonExpressionTypeStringMark,
		',': Modles.JsonExpressionTypeComma,
		':': Modles.JsonExpressionTypeColon,
	}
	defer func() {
		*cursor++
	}()
	//获取非值类型的表达式
	startCursor := *cursor //起始游标，主要用于组装字符串的时候用
	for *cursor < length {
		currentChar := jsonString[*cursor]
		switch currentChar {
		case ' ':
			if startCursor == *cursor {
				startCursor += 1
			}
			continue
		case '\\':
			//转义符需要特殊处理，直接把转义符后面一个字节默认成普通字节，所以游标先加一   等循环末尾再加一
			*cursor++
		case '{', '}', '[', ']', '"', ',', ':':
			if startCursor == *cursor { //起始游标和当前游标还一致  说明了之前没找到普通字符串，则直接返回一个特殊表达式
				expression = Modles.JsonExpression{
					Type:          CharMap[jsonString[*cursor]],
					Content:       string(CharMap[jsonString[*cursor]]),
					StartPosition: startCursor,
					EndPosition:   *cursor,
					Pre:           nil,
					Next:          nil,
				}
			} else {
				expression = Modles.JsonExpression{
					Type:          Modles.JsonExpressionTypeValue,
					Content:       string(jsonString[startCursor:*cursor]),
					StartPosition: startCursor,
					EndPosition:   *cursor,
					Pre:           nil,
					Next:          nil,
				}
				*cursor -= 1 //因为已经遍历到了下一个特殊字符了  所以需要把游标移到字符串结束的位置
			}
			return
		}
		*cursor++
	}
	return
}
