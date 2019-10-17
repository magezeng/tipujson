package TipuJson

import (
	"fmt"
	"testing"
)

type TempStruct struct {
	a         int
	b         int
	subStruct TempSubStruct
	subSlice  []string
	subMap    map[string]float64
}
type TempSubStruct struct {
	s int
	b string
}

func TestObjectToJsonString_struct(t *testing.T) {
	result, err := ObjectToJsonString(TempStruct{
		a: 10,
		b: 30,
		subStruct: TempSubStruct{
			s: 3,
			b: "我是谁",
		},
		subSlice: []string{"123", "345"},
		subMap:   map[string]float64{"kk": 12.3, "BB": 45.9},
	})
	fmt.Println(result, err)
}
