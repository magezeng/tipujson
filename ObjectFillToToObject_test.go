package tipujson

import (
	"fmt"
	"testing"
)

func TestObjectFillToObject1(t *testing.T) {
	//层级关系：struct>slice>map
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Boy  bool   `json:"boy"`
	}
	type School struct {
		SchoolName  string    `json:"school_name"`
		SchoolStage string    `json:"school_stage"`
		Grade       int       `json:"grade"`
		Students    []Student `json:"students"`
	}
	var fromObject = map[string]interface{}{
		"school_name":  "成都信息工程大学实验小学",
		"school_stage": "小学",
		"grade":        3,
		"students": []map[string]interface{}{
			{"name": "小明", "age": 10, "boy": true},
			{"name": "小红", "age": 11, "boy": false},
			{"name": "小李", "age": 10, "boy": true},
		},
	}
	var slicePosition = []string{"students", "name"}
	var sliceFrom interface{} = fromObject
	school := School{}
	student := Student{}
	err := ObjectFillToObject(fromObject, &school, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		position = slicePosition
		fromSlice = sliceFrom
		directionSlice = student
		return
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(school)
}

func TestObjectFillToObject2(t *testing.T) {
	//层级关系:struct>slice>struct>
	type Hobby struct {
		Swim  bool   `json:"swim"`
		Other string `json:"other"`
	}
	type Student struct {
		Name    string `json:"name"`
		Hobbies Hobby  `json:"hobbies"`
	}
	type Class struct {
		Classname string    `json:"classname"`
		Students  []Student `json:"students"`
	}
	var fromObject = map[string]interface{}{
		"classname": "一年级二班",
		"students": []map[string]interface{}{
			{"name": "小红", "hobbies": map[string]interface{}{"swim": true, "other": "唱歌"}},
			{"name": "小明", "hobbies": map[string]interface{}{"swim": false, "other": "打篮球"}},
		},
	}
	var slicePosition = []string{"students", "hobbies", "other"}
	var sliceFrom interface{} = fromObject
	class := Class{}
	student := Student{}
	err := ObjectFillToObject(fromObject, &class, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		position = slicePosition
		fromSlice = sliceFrom
		directionSlice = student
		return
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(class)
}

func TestObjectFillToObject3(t *testing.T) {
	//层级关系:struct>struct>slice
	type Class struct {
		ClassName string `json:"class_name"`
		Students  int    `json:"students"`
	}
	type Grade struct {
		GradeName string  `json:"grade_name"`
		HaveClass []Class `json:"have_class"`
	}
	type School struct {
		SchoolName string `json:"school_name"`
		HaveGrade  Grade  `json:"have_grade"`
	}
	var fromObject = map[string]interface{}{
		"school_name": "没希望小学",
		"have_grade": map[string]interface{}{
			"grade_name": "一年级",
			"have_class": []map[string]interface{}{
				{"class_name": "一班", "students": 30},
				{"class_name": "二班", "students": 40},
				{"class_name": "三班", "students": 50},
			},
		},
	}
	var slicePosition = []string{"have_grade", "have_class", "students"}
	var sliceFrom interface{} = fromObject
	school := School{}
	grade := Grade{}
	err := ObjectFillToObject(fromObject, &school, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		position = slicePosition
		fromSlice = sliceFrom
		directionSlice = grade
		return
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(school)
}
