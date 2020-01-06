package tipujson

import (
	"regexp"
	"strings"
	"testing"
)

func TestObjectToJsonString_mapToJsonString(t *testing.T) {
	MapFromObject := map[string]interface{}{
		"school_name":  "成都信息工程大学实验小学",
		"school_stage": "小学",
		"grade":        3,
		"students": []map[string]interface{}{
			{"name": "小明", "age": 10, "boy": true},
			{"name": "小红", "age": 11, "boy": false},
		},
	}
	//测试结果
	testResult, err := objectToJsonString(MapFromObject)
	if err != nil {
		t.Error("mapObject转jsonString失败")
	}
	var contain bool
	var subRexString string
	subRexString = `"school_name"\s*:\s*"成都信息工程大学实验小学"`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"school_stage"\s*:\s*"小学"`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"grade"\s*:\s*3`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"age"\s*:\s*10`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"age"\s*:\s*11`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"name"\s*:\s*"小明"`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"name"\s*:\s*"小红"`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"boy"\s*:\s*true`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"boy"\s*:\s*false`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*,\s*,\s*,\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*\{\s*,\s*,\s*\}\s*,\s*\{\s*,\s*,\s*\}\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `"students"\s*:\s*\[\s*\]`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*\{\s*\}\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	if len(testResult) > 0 {
		t.Error("MapObject获取的结果字符串取缔掉所有应该匹配的对象之后，字符串结果仍然不为空")
	}
}
func TestObjectToJsonString_SliceToJsonString(t *testing.T) {
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
	testResult, err := objectToJsonString(SliceFromObject)
	if err != nil {
		t.Error("sliceObject转jsonString失败")
	}
	var contain bool
	var subRexString string
	subRexString = `\{\s*"name"\s*:\s*"小明"\s*,\s*"age"\s*:\s*10\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\{\s*"name"\s*:\s*"小红"\s*,\s*"age"\s*:\s*11\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\{\s*"name"\s*:\s*"小李"\s*,\s*"age"\s*:\s*12\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\[\s*,\s*,\s*\]`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	if len(testResult) > 0 {
		t.Error("sliceObject获取的结果字符串取缔掉所有应该匹配的对象之后，字符串结果仍然不为空")
	}
}

func TestObjectToJsonString_StructToJsonString(t *testing.T) {
	structFromObject := School{
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
				Age:  12,
			},
		},
	}
	//测试结果
	testResult, err := objectToJsonString(structFromObject)
	if err != nil {
		t.Error("structObject转jsonString失败")
	}
	var contain bool
	var subRexString string
	subRexString = `\{\s*"name"\s*:\s*"小明"\s*,\s*"age"\s*:\s*10\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\{\s*"name"\s*:\s*"小红"\s*,\s*"age"\s*:\s*11\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\{\s*"name"\s*:\s*"小李"\s*,\s*"age"\s*:\s*12\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*"school_name"\s*:\s*"成都信息工程大学实验小学"\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*"school_stage"\s*:\s*"小学"\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*"grade"\s*:\s*4\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*\[\s*,\s*,\s*\]`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*,\s*,\s*,`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\s*"students"\s*:\s*`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	subRexString = `\{\s*\}`
	if contain, testResult = containSubString(testResult, subRexString); !contain {
		t.Error("结果内未包含" + subRexString)
	}
	if len(testResult) > 0 {
		t.Error("structObject获取的结果字符串取缔掉所有应该匹配的对象之后，字符串结果仍然不为空")
	}
}

func containSubString(ori string, subRexString string) (contain bool, removed string) {
	re, _ := regexp.Compile(subRexString)
	result := re.FindStringSubmatch(ori)
	contain = len(result) == 1
	if !contain {
		return
	}
	removed = strings.Replace(ori, result[0], "", 1)
	return
}
