package TipuJson

import (
	"fmt"
	"testing"
)

func TestObjectToJsonString_struct(t *testing.T) {
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
	result, _ := ObjectToJsonString(Classes{
		SchoolName:  "成都信息工程大学实验小学",
		SchoolStage: "小学",
		Grade:       3,
		Students: []Student{{
			Name: "小明",
			Age:  10,
		}, {
			Name: "小张",
			Age:  10,
		}, {
			Name: "小李",
			Age:  10,
		}},
	})
	fmt.Println(result)
}
