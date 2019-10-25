package TipuJson

import (
	"fmt"
	"github.com/magezeng/TipuJson/Modles"
	"testing"
)

func TestGetJsonFieldFromString(t *testing.T) {
	fromString := `{"a":2.2}`
	field, err := GetJsonFieldFromString(fromString)
	fmt.Println(field.Content.(map[string]*Modles.JsonField)["a"].Content)
	print(field, err)
}
