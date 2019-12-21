package tipujson

import (
	"fmt"
	"testing"
)

func TestGetJsonObjectFromString(t *testing.T) {
	fromString := `{"school_name":"成都信息工程大学实验小学","school_stage":"小学","grade":3,"students":[{"name":"小明","age":10},{"name":"小张","age":10},{"name":"小李","age":10}]}`
	directionValue, err := getJsonObjectFromString([]byte(fromString))
	fmt.Println(directionValue, err)
}
