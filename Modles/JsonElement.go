package Modles

import "github.com/magezeng/tipujson/Modles/JsonElementType"

type JsonElement struct {
	Type     JsonElementType.JsonElementType
	Contents []string
}
