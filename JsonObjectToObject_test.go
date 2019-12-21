package tipujson

import (
	"errors"
	"fmt"
	"testing"
)

func TestJsonObjectToObject(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type School struct {
		SchoolName  string    `json:"school_name"`
		SchoolStage string    `json:"school_stage"`
		Grade       int       `json:"grade"`
		Students    []Student `json:"students"`
	}
	fromString := `{"school_name":"成都信息工程大学实验小学","school_stage":"小学","grade":3,"students":[{"name":"小明","age":10},{"name":"小张","age":10},{"name":"小李","age":10}]}`
	jsonObject, err := getJsonObjectFromString([]byte(fromString))
	if err != nil {
		panic(errors.New("字符串转换为Json对象失败"))
	}
	school := School{}
	err = jsonObjectToObject(jsonObject, &school)
	if err != nil {
		panic(errors.New("Json对象转换为对象失败"))
	}
	fmt.Println(school)
}
