package main

import (
	"TipuJson/RunParamsUnitType"
	"TipuJson/TipuJson"
	"TipuJson/TipuJson/Tools/JsonScan"
	"fmt"
)

type RunParamsUnit struct {
	NickName    string                              `json:"nick_name"`   //在配置界面展示给用户的名称
	Name        string                              `json:"name"`        //参数的名称，在策略内会以此名称进行调用
	Type        RunParamsUnitType.RunParamsUnitType `json:"type"`        //参数的类型，类型分为   字符串  数字  布尔值  列表选择  四种
	Default     interface{}                         `json:"default"`     //参数默认值
	Description string                              `json:"description"` //参数描述
	Index       int                                 `json:"index"`       //参数在展示的时候的排序位置
}

func main() {
	fromString := `{"nick_name":"参数111111","name":"params1","type":1,"default":"haha","description":"该参数是在测试","index":1}`
	var result RunParamsUnit

	err := TipuJson.StringToObj(fromString, &result)
	fmt.Println(err)

	expressions, _ := JsonScan.ScanJsonExpressions([]byte(fromString))
	for ; expressions != nil; expressions = expressions.Next {
		fmt.Println(expressions.Content)
	}

}
