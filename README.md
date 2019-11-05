# Json Analysis Framework
Since the json of the Golang framework encountered many inconveniences in use, the development of Tippu Technology developed TipuJson for parsing json strings.
## Installation
1. Install
```sh
$ go get -u TipuJson
```
2. Import it in your code:

```go
import "github.com/magezeng/tipujson"
```
##Examples
```go

import (
	"fmt"
	"testing"
)
type Student struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
type SchoolStage string

const (
	Primary = "Primary" //小学
	Middle  = "Middle"  //初中
	High    = "High"    // 高中
)

type Classes struct {
	SchoolName  string      `json:"school_name"`
	SchoolStage SchoolStage `json:"school_stage"`
	Grade       uint        `json:"grade"`
	Students    []Student   `json:"students"`
}

func TestStringToObj(t *testing.T) {
	fromString := `{"school_name":"成都信息工程大学实验小学","school_stage":"小学","grade":3,"students":[{"name":"小明","age":10},{"name":"小张","age":10},{"name":"小李","age":10}]}`
	//result := []map[string]interface{}{}
	var result Classes
	err := StringToObj(fromString, &result)
	fmt.Println(err)
	fmt.Println(result)
}
```
