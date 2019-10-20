package TipuJson

import (
	"fmt"
	"testing"
)

func TestObjectToJsonString_struct(t *testing.T) {
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
