package tipujson

import (
	"github.com/magezeng/tipujson/ErrorMaker"
	"testing"
)

type Class struct {
	Name   string  `json:"name"`
	Number int     `json:"number"`
	Boy    bool    `json:"boy"`
	Age    float64 `json:"age"`
}
type Grade struct {
	Name    string   `json:"name"`
	Classes *Class   `json:"classes"`
	Lesson  []string `json:"lesson"`
}

func TestObjectFillToObject_MapToMap(t *testing.T) {
	var fromObject = map[string]interface{}{
		"name":   "没希望的一年级",
		"lesson": []string{"语文", "英语"},
		"classes": &map[string]interface{}{
			"name":   "小红",
			"age":    15.5,
			"number": 156,
			"boy":    false,
		},
	}
	var grade map[string]interface{}
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充Map对象为空")
	}
	if grade["name"] != "没希望的一年级" {
		t.Error("Map对象的name字段填充错误")
	}

	lessons := grade["lesson"].([]interface{})
	if len(lessons) != 2 {
		t.Error("Map对象的lesson字段长度填充错误")
	}
	if lessons[0] == nil {
		t.Error("lesson数组第一项为空")
	}
	if lessons[1] == nil {
		t.Error("lesson数组第二项为空")
	}
	if (lessons[0] == "语文" && lessons[1] != "英语") || (lessons[0] == "英语" && lessons[1] != "语文") {
		t.Error("结构体中的数组内容填充错误")
	}
	if grade["classes"] == nil {
		t.Error("Map对象的classes字段填充为空")
	}
	classes := grade["classes"].(map[string]interface{})
	if classes["age"] != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if classes["boy"] != false {
		t.Error("Map对象的classes字段boy项填充错误")
	}
	if classes["name"] != "小红" {
		t.Error("Map对象的classes字段name项填充错误")
	}
	if classes["number"] != 156 {
		t.Error("Map对象的classes字段number项填充错误")
	}

}

func TestObjectFillToObject_MapToStruct(t *testing.T) {
	var fromObject = map[string]interface{}{
		"name":   "没希望的一年级",
		"lesson": []string{"语文", "英语"},
		"classes": &map[string]interface{}{
			"name":   "小红",
			"age":    15.5,
			"number": 156,
			"boy":    false,
		},
	}
	grade := Grade{}
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade.Name != "没希望的一年级" {
		t.Error("结构体的Name项填充错误")
	}
	if grade.Lesson == nil {
		t.Error("结构体中的数组填充为空")
	}
	if len(grade.Lesson) != 2 {
		t.Error("结构体中的数组填充长度不对")
	}
	if (grade.Lesson[0] == "语文" && grade.Lesson[1] != "英语") || (grade.Lesson[0] == "英语" && grade.Lesson[1] != "语文") {
		t.Error("结构体中的数组内容填充错误")
	}
	if grade.Classes == nil {
		t.Error("结构体中的指针填充为空")
	}
	if grade.Classes.Name != "小红" {
		t.Error("结构体中的指针的Name项填充错误")
	}
	if grade.Classes.Number != 156 {
		t.Error("结构体中的指针的Number项填充错误")
	}
	if grade.Classes.Age != 15.5 {
		t.Error("结构体中的指针的Age项填充错误")
	}
	if grade.Classes.Boy != false {
		t.Error("结构体中的指针的Boy项填充错误")
	}
}

func TestObjectFillToObject_MapToPtr(t *testing.T) {
	var fromObject = map[string]interface{}{
		"name":   "没希望的一年级",
		"lesson": []string{"语文", "英语"},
		"classes": &map[string]interface{}{
			"name":   "小红",
			"age":    15.5,
			"number": 156,
			"boy":    false,
		},
	}
	var grade *Grade
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充指针为空")
	}
	if grade.Name != "没希望的一年级" {
		t.Error("数组中Name项填充错误")
	}
	if grade.Lesson == nil {
		t.Error("数组中Lesson项填充为空")
	}
	if len(grade.Lesson) != 2 {
		t.Error("数组中Lesson项长度填充错误")
	}
	if (grade.Lesson[0] == "语文" && grade.Lesson[1] != "英语") || (grade.Lesson[0] == "英语" && grade.Lesson[1] != "语文") {
		t.Error("数组中Lesson项内容填充错误")
	}
	if grade.Classes == nil {
		t.Error("数组中Classes项填充为空")
	}
	if grade.Classes.Age != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Boy != false {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Name != "小红" {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Number != 156 {
		t.Error("Map对象的classes字段age项填充错误")
	}
}

func TestObjectFillToObject_StructToStruct(t *testing.T) {
	var fromObject = Grade{
		Name:   "没希望的一年级",
		Lesson: []string{"数学", "英语"},
		Classes: &Class{
			Name:   "小红",
			Age:    15.5,
			Number: 156,
			Boy:    false,
		},
	}
	var grade Grade
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade.Name != "没希望的一年级" {
		t.Error("结构体的Name项填充错误")
	}
	if grade.Lesson == nil {
		t.Error("结构体中的数组填充为空")
	}
	if len(grade.Lesson) != 2 {
		t.Error("结构体中的数组填充长度不对")
	}
	if (grade.Lesson[0] == "语文" && grade.Lesson[1] != "英语") || (grade.Lesson[0] == "英语" && grade.Lesson[1] != "语文") {
		t.Error("结构体中的数组内容填充错误")
	}
	if grade.Classes == nil {
		t.Error("结构体中的指针填充为空")
	}
	if grade.Classes.Name != "小红" {
		t.Error("结构体中的指针的Name项填充错误")
	}
	if grade.Classes.Number != 156 {
		t.Error("结构体中的指针的Number项填充错误")
	}
	if grade.Classes.Age != 15.5 {
		t.Error("结构体中的指针的Age项填充错误")
	}
	if grade.Classes.Boy != false {
		t.Error("结构体中的指针的Boy项填充错误")
	}
}

func TestObjectFillToObject_StructToMap(t *testing.T) {
	var fromObject = Grade{
		Name:   "没希望的一年级",
		Lesson: []string{"数学", "英语"},
		Classes: &Class{
			Name:   "小红",
			Age:    15.5,
			Number: 156,
			Boy:    false,
		},
	}
	var grade map[string]interface{}
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充Map对象为空")
	}
	if grade["name"] != "没希望的一年级" {
		t.Error("Map对象的name字段填充错误")
	}
	lessons := grade["lesson"].([]interface{})
	if len(lessons) != 2 {
		t.Error("Map对象的lesson字段长度填充错误")
	}
	if lessons[0] == nil {
		t.Error("lesson数组第一项为空")
	}
	if lessons[1] == nil {
		t.Error("lesson数组第二项为空")
	}
	if (lessons[0] == "语文" && lessons[1] != "英语") || (lessons[0] == "英语" && lessons[1] != "语文") {
		t.Error("结构体中的数组内容填充错误")
	}
	if grade["classes"] == nil {
		t.Error("Map对象的classes字段填充为空")
	}
	classes := grade["classes"].(map[string]interface{})
	if classes["age"] != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if classes["boy"] != false {
		t.Error("Map对象的classes字段boy项填充错误")
	}
	if classes["name"] != "小红" {
		t.Error("Map对象的classes字段name项填充错误")
	}
	if classes["number"] != 156 {
		t.Error("Map对象的classes字段number项填充错误")
	}
}

func TestObjectFillToObject_StructToPtr(t *testing.T) {
	var fromObject = Grade{
		Name:   "没希望的一年级",
		Lesson: []string{"数学", "英语"},
		Classes: &Class{
			Name:   "小红",
			Age:    15.5,
			Number: 156,
			Boy:    false,
		},
	}
	var grade *Grade
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充指针为空")
	}
	if grade.Name != "没希望的一年级" {
		t.Error("数组中Name项填充错误")
	}
	if grade.Lesson == nil {
		t.Error("数组中Lesson项填充为空")
	}
	if len(grade.Lesson) != 2 {
		t.Error("数组中Lesson项长度填充错误")
	}
	if (grade.Lesson[0] == "语文" && grade.Lesson[1] != "英语") || (grade.Lesson[0] == "英语" && grade.Lesson[1] != "语文") {
		t.Error("数组中Lesson项内容填充错误")
	}
	if grade.Classes == nil {
		t.Error("数组中Classes项填充为空")
	}
	if grade.Classes.Age != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Boy != false {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Name != "小红" {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Number != 156 {
		t.Error("Map对象的classes字段age项填充错误")
	}

}

func TestObjectFillToObject_SliceToSlice(t *testing.T) {
	var fromObject = []Grade{
		{
			Name:   "没希望的一年级",
			Lesson: []string{"数学", "英语"},
			Classes: &Class{
				Name:   "小红",
				Age:    15.5,
				Number: 156,
				Boy:    false,
			},
		},
	}
	var grade []Grade
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充数组为空")
	}
	if grade[0].Name != "没希望的一年级" {
		t.Error("数组中Name项填充错误")
	}
	if grade[0].Lesson == nil {
		t.Error("数组中Lesson项填充为空")
	}
	if len(grade[0].Lesson) != 2 {
		t.Error("数组中Lesson项长度填充错误")
	}
	if (grade[0].Lesson[0] == "语文" && grade[0].Lesson[1] != "英语") || (grade[0].Lesson[0] == "英语" && grade[0].Lesson[1] != "语文") {
		t.Error("数组中Lesson项内容填充错误")
	}
	if grade[0].Classes == nil {
		t.Error("数组中Classes项填充为空")
	}
	if grade[0].Classes.Age != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade[0].Classes.Boy != false {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade[0].Classes.Name != "小红" {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade[0].Classes.Number != 156 {
		t.Error("Map对象的classes字段age项填充错误")
	}
}

func TestObjectFillToObject_SliceToPtr(t *testing.T) {
	var fromObject = []Grade{
		{
			Name:   "没希望的一年级",
			Lesson: []string{"数学", "英语"},
			Classes: &Class{
				Name:   "小红",
				Age:    15.5,
				Number: 156,
				Boy:    false,
			},
		},
	}
	var grade *[]Grade
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充指针为空")
	}
	currentGrade := *grade
	if currentGrade == nil {
		t.Error("被填充数组为空")
	}
	if currentGrade[0].Name != "没希望的一年级" {
		t.Error("数组中Name项填充错误")
	}
	if currentGrade[0].Lesson == nil {
		t.Error("数组中Lesson项填充为空")
	}
	if len(currentGrade[0].Lesson) != 2 {
		t.Error("数组中Lesson项长度填充错误")
	}
	if (currentGrade[0].Lesson[0] == "语文" && currentGrade[0].Lesson[1] != "英语") || (currentGrade[0].Lesson[0] == "英语" && currentGrade[0].Lesson[1] != "语文") {
		t.Error("数组中Lesson项内容填充错误")
	}
	if currentGrade[0].Classes == nil {
		t.Error("数组中Classes项填充为空")
	}
	if currentGrade[0].Classes.Age != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if currentGrade[0].Classes.Boy != false {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if currentGrade[0].Classes.Name != "小红" {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if currentGrade[0].Classes.Number != 156 {
		t.Error("Map对象的classes字段age项填充错误")
	}
}

func TestObjectFillToObject_PtrToPtr(t *testing.T) {
	var fromObject = &Grade{
		Name:   "没希望的一年级",
		Lesson: []string{"数学", "英语"},
		Classes: &Class{
			Name:   "小红",
			Age:    15.5,
			Number: 156,
			Boy:    false,
		},
	}
	var grade *Grade
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade == nil {
		t.Error("被填充指针为空")
	}
	if grade.Name != "没希望的一年级" {
		t.Error("数组中Name项填充错误")
	}
	if grade.Lesson == nil {
		t.Error("数组中Lesson项填充为空")
	}
	if len(grade.Lesson) != 2 {
		t.Error("数组中Lesson项长度填充错误")
	}
	if (grade.Lesson[0] == "语文" && grade.Lesson[1] != "英语") || (grade.Lesson[0] == "英语" && grade.Lesson[1] != "语文") {
		t.Error("数组中Lesson项内容填充错误")
	}
	if grade.Classes == nil {
		t.Error("数组中Classes项填充为空")
	}
	if grade.Classes.Age != 15.5 {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Boy != false {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Name != "小红" {
		t.Error("Map对象的classes字段age项填充错误")
	}
	if grade.Classes.Number != 156 {
		t.Error("Map对象的classes字段age项填充错误")
	}
}
