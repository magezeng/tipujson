package Modles

type JsonFieldType string

const (
	JsonFieldTypeMap    JsonFieldType = "Map"
	JsonFieldTypeList                 = "List"
	JsonFieldTypeString               = "String"
	JsonFieldTypeNumber               = "Number"
	JsonFieldTypeBool                 = "Bool"
)

type JsonField struct {
	Type    JsonFieldType
	Content interface{} //类型为Map时此字段是Map 类型为List时 此字段是Slice String和Number类型时   直接为值
}
