package tipujson

import (
	"testing"
)

func TestGetJsonFieldFromString(t *testing.T) {
	fromString := `[{"nick_name":"AccessKey","name":"access_key","type":"string","sub_type":"","description":"[{\\n\\\"userid\\\": 3.0,\\\"index\\\":\\\"1\\\",\\\"username\\\": \\\"3号用户\\\"}]","index":0,"have_default":false,"default":null},{"nick_name":"AccessKey","name":"access_key","type":"slice","sub_type":"number","description":"参数序列","index":0,"have_default":false,"default":"用户请选择"}]`
	field, err := getJsonFieldFromString(fromString)
	print(field, err)
}
