package tipujson

import (
	"github.com/magezeng/tipujson/ErrorMaker"
	"github.com/magezeng/tipujson/Modles"
	"github.com/magezeng/tipujson/Modles/JsonElementType"
	"github.com/magezeng/tipujson/Modles/JsonElementTypeRegexp"
	"regexp"
	"strconv"
)

func getJsonObjectFromString(jsonString []byte) (jsonObject interface{}, err error) {
	element, nowIndex, err := getJsonElementSymbol(0, jsonString)
	if err != nil {
		return
	}
	switch element.Type {
	case JsonElementType.MapStart:
		tempMap := map[string]interface{}{}
		nowIndex, err = fillMapFromJsonString(&tempMap, nowIndex, jsonString)
		jsonObject = tempMap
		return
	case JsonElementType.ListStart:
		var tempList []interface{}
		nowIndex, err = fillListFromJsonString(&tempList, nowIndex, jsonString)
		jsonObject = tempList
		return
	default:
		err = ErrorMaker.GetError("字符串不是一个Json字符串")
		return
	}
}

var jsonRegexpByElementTypeMap = map[JsonElementType.JsonElementType]JsonElementTypeRegexp.JsonElementTypeRegexp{
	JsonElementType.MapStart:  JsonElementTypeRegexp.MapStart,
	JsonElementType.MapEnd:    JsonElementTypeRegexp.MapEnd,
	JsonElementType.ListStart: JsonElementTypeRegexp.ListStart,
	JsonElementType.ListEnd:   JsonElementTypeRegexp.ListEnd,
	JsonElementType.String:    JsonElementTypeRegexp.String,
	JsonElementType.Colon:     JsonElementTypeRegexp.Colon,
	JsonElementType.Comma:     JsonElementTypeRegexp.Comma,
	JsonElementType.Num:       JsonElementTypeRegexp.Num,
	JsonElementType.Bool:      JsonElementTypeRegexp.Bool,
	JsonElementType.Null:      JsonElementTypeRegexp.Null,
}

func getJsonElementSymbol(nowIndex int, jsonString []byte) (element Modles.JsonElement, nextIndex int, err error) {
	defer func() {
		nextIndex = nowIndex
	}()
	var re *regexp.Regexp
	//依次判定获取到的内容是否有对应的类型予以匹配
	for tempElementType, tempElementRegexp := range jsonRegexpByElementTypeMap {
		re, err = regexp.Compile(string(tempElementRegexp))
		if err != nil {
			panic(err)
		}
		subStringIndexes := re.FindSubmatchIndex(jsonString[nowIndex:])
		if len(subStringIndexes) >= 2 {
			element.Type = tempElementType
			newContents := make([]string, len(subStringIndexes)/2)
			for tempIndex, _ := range newContents {
				newContents[tempIndex] = string(jsonString[nowIndex+subStringIndexes[tempIndex*2] : nowIndex+subStringIndexes[tempIndex*2+1]])
			}
			element.Contents = newContents
			nowIndex = nowIndex + subStringIndexes[1]
			return
		}
	}
	err = ErrorMaker.GetError("在", nowIndex, "未找到一个Json元素")
	return
}

func fillMapFromJsonString(directionMap *map[string]interface{}, nowIndex int, jsonString []byte) (nextIndex int, err error) {
	defer func() {
		nextIndex = nowIndex
	}()
	for {
		var tempKey string
		var tempValue interface{}
		var tempElement Modles.JsonElement
		//从原始字符串中找到一个字符串作为字典的Key
		tempElement, nowIndex, err = getJsonElementSymbol(nowIndex, jsonString)
		if err != nil {
			return
		}
		switch tempElement.Type {
		case JsonElementType.String:
			tempKey = tempElement.Contents[1]
		default:
			err = ErrorMaker.GetError("字典解析错误")
			return
		}

		//从原始字符串中找到一个":"
		tempElement, nowIndex, err = getJsonElementSymbol(nowIndex, jsonString)
		if err != nil {
			return
		}
		if tempElement.Type != JsonElementType.Colon {
			err = ErrorMaker.GetError("字典解析错误")
			return
		}
		//从原始字符串中找到一个对象作为字典的Value
		tempElement, nowIndex, err = getJsonElementSymbol(nowIndex, jsonString)
		if err != nil {
			return
		}
		switch tempElement.Type {
		case JsonElementType.MapStart:
			tempMap := map[string]interface{}{}
			nowIndex, err = fillMapFromJsonString(&tempMap, nowIndex, jsonString)
			if err != nil {
				return
			}
			tempValue = tempMap
		case JsonElementType.ListStart:
			var tempList []interface{}
			nowIndex, err = fillListFromJsonString(&tempList, nowIndex, jsonString)
			if err != nil {
				return
			}
			tempValue = tempList
		case JsonElementType.String:
			tempValue = tempElement.Contents[1]
		case JsonElementType.Num:
			tempValue, err = strconv.ParseFloat(tempElement.Contents[1], 64)
			if err != nil {
				return
			}
		case JsonElementType.Bool:
			tempValue, err = strconv.ParseBool(tempElement.Contents[1])
			if err != nil {
				return
			}
		case JsonElementType.Null:
			tempValue = nil
		default:
			err = ErrorMaker.GetError(tempElement.Contents[0] + "  不构成一个可使用的对象")
			return
		}

		//将值存入目标字典
		if directionMap == nil {
			directionMap = &map[string]interface{}{}
		}
		(*directionMap)[tempKey] = tempValue

		tempElement, nowIndex, err = getJsonElementSymbol(nowIndex, jsonString)
		switch tempElement.Type {
		case JsonElementType.Comma:
			//逗号说明接下来还有值需要扫描到字典
			continue
		case JsonElementType.MapEnd:
			return
		default:
			err = ErrorMaker.GetError("扫描Map时遇到一个未知错误")
			return
		}
	}
}

func fillListFromJsonString(directionList *[]interface{}, nowIndex int, jsonString []byte) (nextIndex int, err error) {
	defer func() {
		nextIndex = nowIndex
	}()
	for {
		var tempValue interface{}
		var tempElement Modles.JsonElement
		tempElement, nowIndex, err = getJsonElementSymbol(nowIndex, jsonString)
		switch tempElement.Type {
		case JsonElementType.String:
			tempValue = tempElement.Contents[1]
		case JsonElementType.Num:
			tempValue, err = strconv.ParseFloat(tempElement.Contents[1], 64)
			if err != nil {
				return
			}
		case JsonElementType.Bool:
			tempValue, err = strconv.ParseBool(tempElement.Contents[1])
			if err != nil {
				return
			}
		case JsonElementType.MapStart:
			tempMap := map[string]interface{}{}
			nowIndex, err = fillMapFromJsonString(&tempMap, nowIndex, jsonString)
			if err != nil {
				return
			}
			tempValue = tempMap
		case JsonElementType.ListStart:
			var tempList []interface{}
			nowIndex, err = fillListFromJsonString(&tempList, nowIndex, jsonString)
			if err != nil {
				return
			}
			tempValue = tempList
		case JsonElementType.Null:
			tempValue = nil
		default:
			err = ErrorMaker.GetError("类型不能被添加到数组内")
		}
		if directionList == nil {
			*directionList = []interface{}{tempValue}
		} else {
			tempList := append(*directionList, tempValue)
			*directionList = tempList
		}
		tempElement, nowIndex, err = getJsonElementSymbol(nowIndex, jsonString)
		switch tempElement.Type {
		case JsonElementType.ListEnd:
			return
		case JsonElementType.Comma:
			continue
		default:
			err = ErrorMaker.GetError("List内容错误")
		}
	}
}
