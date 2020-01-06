package tipujson

import (
	"github.com/magezeng/tipujson/ErrorMaker"
	"testing"
)

func TestObjectFillToObject_StructSliceStruct(t *testing.T) {
	//层级关系：struct>slice>struct
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Boy  bool   `json:"boy"`
	}
	type School struct {
		Students []Student `json:"students"`
	}
	var fromObject = map[string]interface{}{
		"students": []map[string]interface{}{
			{"name": "小明", "age": 10, "boy": true},
			{"name": "小红", "age": 11, "boy": false},
			{"name": "小李", "age": 10, "boy": true},
		},
	}
	school := School{}
	err := objectFillToObject(fromObject, &school, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if school.Students == nil {
		t.Error("填充结果的数组为空")
	}
	if len(school.Students) != 3 {
		t.Error("填充结果的数组长度不对")
	}
	if school.Students[0].Name != "小明" {
		t.Error("数组内第一项的Name填充错误")
	}
	if school.Students[0].Age != 10 {
		t.Error("数组内第一项的Age填充错误")
	}
	if school.Students[0].Boy != true {
		t.Error("数组内第一项的boy填充错误")
	}

	if school.Students[1].Name != "小红" {
		t.Error("数组内第二项的Name填充错误")
	}
	if school.Students[1].Age != 11 {
		t.Error("数组内第二项的Age填充错误")
	}
	if school.Students[1].Boy != false {
		t.Error("数组内第二项的boy填充错误")
	}
	if school.Students[2].Name != "小李" {
		t.Error("数组内第三项的Name填充错误")
	}
	if school.Students[2].Age != 10 {
		t.Error("数组内第三项的Age填充错误")
	}
	if school.Students[2].Boy != true {
		t.Error("数组内第三项的boy填充错误")
	}
}

func TestObjectFillToObject_StructSlicePtr(t *testing.T) {
	//层级关系：Struct/Slice/Pointer
	type Grade struct {
		GradeName string `json:"grade_name"`
		Boy       int    `json:"boy"`
		Girl      int    `json:"girl"`
	}
	type School struct {
		SchoolName string   `json:"school_name"`
		HaveGrade  []*Grade `json:"have_grade"`
	}
	var fromObject = map[string]interface{}{
		"school_name": "没希望小学",
		"have_grade": []*map[string]interface{}{
			{"grade_name": "一年级", "boy": 15, "girl": 10},
			{"grade_name": "二年级", "boy": 25, "girl": 20},
		},
	}
	school := School{}
	err := objectFillToObject(fromObject, &school, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if school.SchoolName != "没希望小学" {
		t.Error("最外层SchoolName填充错误")
	}
	if len(school.HaveGrade) != 2 {
		t.Error("填充数组长度不对")
	}
	if school.HaveGrade[0] == nil {
		t.Error("指针第一项为空")
	}
	if school.HaveGrade[0].GradeName != "一年级" {
		t.Error("指针第一项的GradeName填充错误")
	}
	if school.HaveGrade[0].Boy != 15 {
		t.Error("指针第一项的Boy填充错误")
	}
	if school.HaveGrade[0].Girl != 10 {
		t.Error("指针第一项的Girl填充错误")
	}
	if school.HaveGrade[1] == nil {
		t.Error("指针第二项为空")
	}
	if school.HaveGrade[1].GradeName != "二年级" {
		t.Error("指针第二项的GradeName填充错误")
	}
	if school.HaveGrade[1].Boy != 25 {
		t.Error("指针第二项的Boy填充错误")
	}
	if school.HaveGrade[1].Girl != 20 {
		t.Error("指针第二项的Girl填充错误")
	}
}

func TestObjectFillToObject_StructPtrSlice(t *testing.T) {
	//层级关系：Struct/Pointer/Slice
	type Student struct {
		Boy  []string `json:"boy"`
		Girl int      `json:"girl"`
	}
	type Grade struct {
		Class    string   `json:"class"`
		Ratio    float64  `json:"ratio"`
		Students *Student `json:"students"`
	}
	var fromObject = map[string]interface{}{
		"class": "一年级",
		"ratio": 0.16,
		"students": map[string]interface{}{
			"boy":  []string{"小明", "小红"},
			"girl": 30,
		},
	}
	grade := Grade{}
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade.Class != "一年级" {
		t.Error("最外层class填充错误")
	}
	if grade.Ratio != 0.16 {
		t.Error("最外层Ratio填充错误")
	}
	if grade.Students == nil {
		t.Error("指针为空")
	}
	if grade.Students.Girl != 30 {
		t.Error("指针第二项Girl填充错误")
	}
	if grade.Students.Boy == nil {
		t.Error("指针第一项Boy填充错误")
	}
	if grade.Students.Boy[0] != "小红" && grade.Students.Boy[0] != "小明" {
		t.Error("数组第一项填充错误")
	}
	if grade.Students.Boy[1] != "小红" && grade.Students.Boy[1] != "小明" {
		t.Error("数组第二项填充错误")
	}
}

func TestObjectFillToObject_StructPtrStruct(t *testing.T) {
	//层级关系：Struct/Pointer/Struct
	type Class struct {
		Name   string  `json:"name"`
		Age    float64 `json:"age"`
		Number int     `json:"number"`
	}
	type Grade struct {
		Classes *Class `json:"classes"`
	}
	var fromObject = map[string]interface{}{
		"classes": map[string]interface{}{
			"number": 45,
			"age":    10.5,
			"name":   "五年级",
		},
	}
	grade := Grade{}
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade.Classes == nil {
		t.Error("填充指针为空")
	}
	if grade.Classes.Number != 45 {
		t.Error("指针中Number项填充错误")
	}
	if grade.Classes.Name != "五年级" {
		t.Error("指针中Name项填充错误")
	}
	if grade.Classes.Age != 10.5 {
		t.Error("指针中Age项填充错误")
	}

}

func TestObjectFillToObject_StructStructSlice(t *testing.T) {
	//层级关系:Struct>Struct>Slice
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
			},
		},
	}
	school := School{}
	err := objectFillToObject(fromObject, &school, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if school.SchoolName != "没希望小学" {
		t.Error("最外层SchoolName填充错误")
	}
	if school.HaveGrade.GradeName != "一年级" {
		t.Error("内层GradeName填充错误")
	}
	if school.HaveGrade.HaveClass[0].ClassName != "一班" {
		t.Error("数组中第一项的ClassName填充错误")
	}
	if school.HaveGrade.HaveClass[0].Students != 30 {
		t.Error("数组中第一项的Students填充错误")
	}
	if school.HaveGrade.HaveClass[1].ClassName != "二班" {
		t.Error("数组中第二项的ClassName填充错误")
	}
	if school.HaveGrade.HaveClass[1].Students != 40 {
		t.Error("数组中第二项的Students填充错误")
	}
}

func TestObjectFillToObject_StructStructPtr(t *testing.T) {
	//层级关系:Struct>Struct>Ptr
	type Sex struct {
		Boy int  `json:"boy"`
		Max bool `json:"max"`
	}
	type Class struct {
		Name   string  `json:"name"`
		Age    float64 `json:"age"`
		Number *Sex    `json:"number"`
	}
	type Grade struct {
		Classes Class `json:"classes"`
	}
	var fromObject = map[string]interface{}{
		"classes": map[string]interface{}{
			"age":  10.5,
			"name": "五年级",
			"number": map[string]interface{}{
				"boy": 15,
				"max": true,
			},
		},
	}
	grade := Grade{}
	err := objectFillToObject(fromObject, &grade, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if grade.Classes.Age != 10.5 {
		t.Error("内层Age项填充出错")
	}
	if grade.Classes.Name != "五年级" {
		t.Error("内层Name项填充出错")
	}
	if grade.Classes.Number == nil {
		t.Error("指针填充为空")
	}
	if grade.Classes.Number.Boy != 15 {
		t.Error("指针中Boy项填充出错")
	}
	if grade.Classes.Number.Max != true {
		t.Error("指针中Max项填充出错")
	}

}

func TestObjectFillToObject_SlicePtr(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Boy  bool   `json:"boy"`
	}
	var fromObject = []*map[string]interface{}{
		{"name": "小明", "age": 10, "boy": true},
		{"name": "小红", "age": 11, "boy": false},
	}

	var students []*Student
	err := objectFillToObject(fromObject, &students, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if len(students) != 2 {
		t.Error("填充结果的数组长度不对")
	}
	if students[0] == nil {
		t.Error("数组中第一项指针为空")
	}
	if students[0].Name != "小明" {
		t.Error("第一个指针Name项填充出错")
	}
	if students[0].Boy != true {
		t.Error("第一个指针Boy项填充出错")
	}
	if students[0].Age != 10 {
		t.Error("第一个指针Age项填充出错")
	}
	if students[1] == nil {
		t.Error("数组中第一项指针为空")
	}
	if students[1].Name != "小红" {
		t.Error("第一个指针Name项填充出错")
	}
}

func TestObjectFillToObject_PtrSlice(t *testing.T) {
	type Student struct {
		Name []string `json:"name"`
		Age  int      `json:"age"`
		Boy  bool     `json:"boy"`
	}
	var fromObject = map[string]interface{}{
		"name": []string{"小明", "小红"},
		"age":  10,
		"boy":  true,
	}
	var students *Student
	err := objectFillToObject(fromObject, &students, func(position []string, fromSlice interface{}, directionSlice interface{}) (i interface{}, b bool) {
		return
	})
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	if students == nil {
		t.Error("填充指针为空")
	}
	if students.Age != 10 {
		t.Error("指针中Age项填充出错")
	}
	if students.Boy != true {
		t.Error("指针中Boy项填充出错")
	}
	if len(students.Name) != 2 {
		t.Error("填充结果数组长度不对")
	}
	if students.Name[0] != "小明" && students.Name[0] != "小红" {
		t.Error("数组中第一项填充出错")
	}
	if students.Name[1] != "小明" && students.Name[1] != "小红" {
		t.Error("数组中第二项填充出错")
	}

}
