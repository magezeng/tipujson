package tipujson

import (
	"fmt"
	"testing"
)

func TestTipuJsonObjectToObject(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}
	type SchoolStage string
	type Classes struct {
		SchoolName  string      `json:"school_name"`
		SchoolStage SchoolStage `json:"school_stage"`
		Grade       uint        `json:"grade"`
		Students    []Student   `json:"students"`
	}
	fromString := `{"school_name":"成都信息工程大学实验小学","school_stage":"小学","grade":3,"students":[{"name":"小明","age":10},{"name":"小张","age":10},{"name":"小李","age":10}]}`
	testMap := map[string]interface{}{}
	err := JsonStringToObject(fromString, &testMap)
	if err != nil {
		return
	}
	//testMap := map[interface{}]interface{}{"school_name":"成都信息工程大学实验小学","school_stage":"小学"}
	var resultClass Classes
	err1 := ObjectToObject(testMap, &resultClass)
	fmt.Println(err1)
	fmt.Println(resultClass)
}
