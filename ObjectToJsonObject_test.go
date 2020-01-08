package tipujson

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestObjectToJsonObject_MapToJsonObject(t *testing.T) {
	//map类型的测试
	MapFromObject := map[string]interface{}{
		"school_name":  "成都信息工程大学实验小学",
		"school_stage": "小学",

		"grade": 3,
		"students": []map[string]interface{}{
			{"name": "小明", "age": 10, "boy": true},
		},
	}
	//测试结果
	testResult, err := objectToJsonObject(MapFromObject)
	//fmt.Println(reflect.TypeOf(testResult))
	if err != nil {
		//t.Error(ErrorMaker.GetErrorStringFromErr(err))
		t.Error("mapObject转jsonObject失败")
	}
	MapResult := testResult.(map[string]interface{})
	if _, ok := MapResult["school_name"]; !ok {
		t.Error("mapObject转jsonObject失败,结果中未包含key:school_name")
	}
	if MapResult["school_name"] != "成都信息工程大学实验小学" {
		t.Error("mapObject转jsonObject失败,结果中未包含value:成都信息工程大学实验小学")
	}
	if _, ok := MapResult["school_stage"]; !ok {
		t.Error("mapObject转jsonObject失败,结果中未包含key:school_stage")
	}
	if MapResult["school_stage"] != "小学" {
		t.Error("mapObject转jsonObject失败,结果中未包含value:小学")
	}
	if _, ok := MapResult["grade"]; !ok {
		t.Error("mapObject转jsonObject失败,结果中未包含key:grade")
	}
	if MapResult["grade"] != 3 {
		t.Error("mapObject转jsonObject失败,结果中未包含value:3")
	}
	if _, ok := MapResult["students"]; !ok {
		t.Error("mapObject转jsonObject失败,结果中未包含key:students")
	}
	if len(MapResult["students"].([]interface{})) != 1 {
		t.Error("mapObject转jsonObject失败,获取到键名为students键值里面不是长度为3的map切片")
	}
	for _, subMapResult := range MapResult["students"].([]interface{}) {
		value := map[string]interface{}{"name": "小明", "age": 10, "boy": true}
		if !reflect.DeepEqual(value, subMapResult) {
			t.Error(fmt.Sprintf("mapObject转jsonObject失败,未获取值%v", value))
		}
	}
}

//切片类型的测试
func TestObjectToJsonObject_SliceToJsonObject(t *testing.T) {
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
			Age:  12,
		},
	}
	//测试结果
	testResult, err := objectToJsonObject(SliceFromObject)
	if err != nil {
		t.Error("sliceObject转jsonObject失败")
	}
	//fmt.Println(reflect.TypeOf(testResult)
	sliceResult := testResult.([]interface{})
	if len(sliceResult) != 3 {
		t.Error("mapObject转jsonObject失败,未转换成正确的长度")
	}
	//定义一个长度为3的bool切片
	boolSlice := []bool{false, false, false}
	for _, subResult := range sliceResult {
		tempResult := subResult.(map[string]interface{})
		if tempResult["name"] == "小明" && !boolSlice[0] {
			if tempResult["age"] == 10 {
				boolSlice[0] = true
			}
		} else if tempResult["name"] == "小红" && !boolSlice[1] {
			if tempResult["age"] == 11 {
				boolSlice[1] = true
			}
		} else if tempResult["name"] == "小李" && !boolSlice[2] {
			if tempResult["age"] == 12 {
				boolSlice[2] = true
			}
		}
	}
	for _, v := range boolSlice {
		if !v {
			t.Error("mapObject转jsonObject失败,结果转换不完整！")
		}
	}
}

//结构体类型的测试
func TestObjectToJsonObject_StructToJsonObject(t *testing.T) {
	school := School{
		SchoolName:  "成都信息工程大学实验小学",
		SchoolStage: "小学",
		Grade:       4,
		Students: []Student{
			{
				Name: "小明",
				Age:  10,
			},
		},
	}
	testResult, err := objectToJsonObject(school)
	if err != nil {
		t.Error("structObject转jsonObject失败")
	}
	structResult := testResult.(map[string]interface{})
	if _, ok := structResult["school_name"]; !ok {
		t.Error("structObject转jsonObject失败,结果中未包含key:school_name")
	}
	if structResult["school_name"] != "成都信息工程大学实验小学" {
		t.Error("structObject转jsonObject失败,结果中未包含value:成都信息工程大学实验小学")
	}
	if _, ok := structResult["school_stage"]; !ok {
		t.Error("structObject转jsonObject失败,结果中未包含key:school_stage")
	}
	if structResult["school_stage"] != "小学" {
		t.Error("structObject转jsonObject失败,结果中未包含value:小学")
	}
	if _, ok := structResult["grade"]; !ok {
		t.Error("structObject转jsonObject失败,结果中未包含key:grade")
	}
	if structResult["grade"] != 4 {
		t.Error("structObject转jsonObject失败,结果中未包含value:3")
	}
	if _, ok := structResult["students"]; !ok {
		t.Error("structObject转jsonObject失败,结果中未包含key:students")
	}
	if len(structResult["students"].([]interface{})) != 1 {
		t.Error("structObject转jsonObject失败,获取到键名为students键值里面不是长度为3的map切片")
	}
	//
	boolSlice := []bool{false}
	for _, subResult := range structResult["students"].([]interface{}) {
		tempResult := subResult.(map[string]interface{})
		if tempResult["name"] == "小明" && !boolSlice[0] {
			if tempResult["age"] == 10 {
				boolSlice[0] = true
			}
		}
	}
	for _, v := range boolSlice {
		if !v {
			t.Error("structObject转jsonObject失败,结果转换不完整！")
		}
	}
}

//指针类型的测试
func TestObjectToJsonObject_PtrToJsonObject(t *testing.T) {
	var structPtr *School
	school := School{
		SchoolName:  "成都信息工程大学实验小学",
		SchoolStage: "小学",
		Grade:       4,
		Students: []Student{
			{
				Name: "小明",
				Age:  10,
			},
		},
	}
	structPtr = &school
	testResult, err := objectToJsonObject(structPtr)
	if err != nil {
		t.Error("ptrObject转jsonObject失败")
	}
	structResult := testResult.(map[string]interface{})
	if _, ok := structResult["school_name"]; !ok {
		t.Error("ptrObject转jsonObject失败,结果中未包含key:school_name")
	}
	if structResult["school_name"] != "成都信息工程大学实验小学" {
		t.Error("ptrObject转jsonObject失败,结果中未包含value:成都信息工程大学实验小学")
	}
	if _, ok := structResult["school_stage"]; !ok {
		t.Error("ptrObject转jsonObject失败,结果中未包含key:school_stage")
	}
	if structResult["school_stage"] != "小学" {
		t.Error("structObject转jsonObject失败,结果中未包含value:小学")
	}
	if _, ok := structResult["grade"]; !ok {
		t.Error("ptrObject转jsonObject失败,结果中未包含key:grade")
	}
	if structResult["grade"] != 4 {
		t.Error("ptrObject转jsonObject失败,结果中未包含value:3")
	}
	if _, ok := structResult["students"]; !ok {
		t.Error("ptrObject转jsonObject失败,结果中未包含key:students")
	}
	if len(structResult["students"].([]interface{})) != 1 {
		t.Error("ptrObject转jsonObject失败,获取到键名为students键值里面不是长度为3的map切片")
	}
	//
	boolSlice := []bool{false}
	for _, subResult := range structResult["students"].([]interface{}) {
		tempResult := subResult.(map[string]interface{})
		if tempResult["name"] == "小明" && !boolSlice[0] {
			if tempResult["age"] == 10 {
				boolSlice[0] = true
			}
		}
	}
	for _, v := range boolSlice {
		if !v {
			t.Error("structObject转jsonObject失败,结果转换不完整！")
		}
	}
}
