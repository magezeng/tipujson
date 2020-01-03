package tipujson

import (
	"testing"
)

func TestObjectToJsonString(t *testing.T) {
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
	//map类型的测试
	MapFromObject := map[string]interface{}{
		"school_name":  "成都信息工程大学实验小学",
		"school_stage": "小学",

		"grade": 3,
		"students": []map[string]interface{}{
			{"name": "小明", "age": 10, "boy": true},
			{"name": "小红", "age": 11, "boy": false},
			{"name": "小李", "age": 10, "boy": true},
		},
	}
	_, err := objectToJsonString(MapFromObject)
	if err != nil {
		//t.Error(ErrorMaker.GetErrorStringFromErr(err))
		t.Error("map转String失败")
	}
	//slice类型测试
	SliceFromObject := []Student{
		{
			Name: "小明",
			Age:  10,
		},
		{
			Name: "小红",
			Age:  11,
		},
		{
			Name: "小李",
			Age:  10,
		},
	}
	_, err = objectToJsonString(SliceFromObject)
	if err != nil {
		t.Error("slice转String失败")
	}
	//struct类型测试
	school := School{
		SchoolName:  "成都信息工程大学实验小学",
		SchoolStage: "小学",
		Grade:       4,
		Students: []Student{
			{
				Name: "小明",
				Age:  10,
			},
			{
				Name: "小红",
				Age:  11,
			},
			{
				Name: "小李",
				Age:  10,
			},
		},
	}
	_, err = objectToJsonString(school)
	if err != nil {
		t.Error("struct转String失败")
	}
}
