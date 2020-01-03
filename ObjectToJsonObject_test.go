package tipujson

import (
	"github.com/magezeng/tipujson/ErrorMaker"
	"testing"
)

func TestObjectToJsonObject(t *testing.T) {
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
	MapfromObject := map[string]interface{}{
		"school_name":  "成都信息工程大学实验小学",
		"school_stage": "小学",
		"grade":        3,
		"students": []map[string]interface{}{
			{"name": "小明", "age": 10, "boy": true},
			{"name": "小红", "age": 11, "boy": false},
			{"name": "小李", "age": 10, "boy": true},
		},
	}
	_, err := objectToJsonObject(MapfromObject)
	if err != nil {
		t.Error(ErrorMaker.GetErrorStringFromErr(err))
	}
	////slice类型测试
	//SlicefromObject:=
	//

}
