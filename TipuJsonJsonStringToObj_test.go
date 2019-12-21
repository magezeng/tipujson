package tipujson

import (
	"fmt"
	"testing"
)

func TestStringToObj(t *testing.T) {
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
	//result := []map[string]interface{}{}
	var result Classes
	err := JsonStringToObject(fromString, &result)
	fmt.Println(err)
	fmt.Println(result)
}

func TestStringToObj_NormalStruct(t *testing.T) {
	type School struct {
		SchoolName  string `json:"school_name"`
		SchoolStage string `json:"school_stage"`
	}
	fromString := `{"school_name":"成都信息工程大学实验小学","school_stage":"小学"}`
	//result := []map[string]interface{}{}
	var result School
	err := JsonStringToObject(fromString, &result)
	fmt.Println(err)
	fmt.Println(result)
}

func TestStringToObj_TypeInType(t *testing.T) {
	type School struct {
		SchoolName  string `json:"school_name"`
		SchoolStage string `json:"school_stage"`
	}
	type Classes struct {
		School
	}
	fromString := `{"school_name":"成都信息工程大学实验小学","school_stage":"小学"}`
	//result := []map[string]interface{}{}
	var result Classes
	err := JsonStringToObject(fromString, &result)
	fmt.Println(err)
	fmt.Println(result)
}
