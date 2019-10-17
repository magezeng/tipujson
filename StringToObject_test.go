package TipuJson

import (
	"fmt"
	"testing"
)

type RunParamsUnitType uint

const (
	TypeString   RunParamsUnitType = 1
	TypeNumber   RunParamsUnitType = 2
	TypeBool     RunParamsUnitType = 3
	TypeList     RunParamsUnitType = 4
	TypeExchange RunParamsUnitType = 5
)

type RunParamsUnit struct {
	Params      *Params           `json:"params"`
	NickName    string            `json:"nick_name"`   //在配置界面展示给用户的名称
	Name        string            `json:"name"`        //参数的名称，在策略内会以此名称进行调用
	Type        RunParamsUnitType `json:"type"`        //参数的类型，类型分为   字符串  数字  布尔值  列表选择  四种
	Description string            `json:"description"` //参数描述
	Index       int               `json:"index"`       //参数在展示的时候的排序位置
}

type Params struct {
	NickNames string            `json:"nick_names"` //在配置界面展示给用户的名称
	Names     string            `json:"names"`      //参数的名称，在策略内会以此名称进行调用
	Types     RunParamsUnitType `json:"types"`      //参数的类型，类型分为   字符串  数字  布尔
}

func TestStringToObj(t *testing.T) {
	fromString := `[{"params":{"nick_names":"哈哈","names":"子对象","types":2},"nick_name":"参数111111","name":"params1","type":1,"description":"该参数是在测试","index":1}]`
	//result := []map[string]interface{}{}
	var result []map[string]interface{}
	//result := new(RunParamsUnit)
	//var result **RunParamsUnit
	err := StringToObj(fromString, &result)
	fmt.Println(err)
	fmt.Println(result)
}
