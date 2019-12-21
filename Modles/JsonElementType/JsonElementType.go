package JsonElementType

type JsonElementType string

const (
	MapStart  JsonElementType = "{"
	MapEnd                    = "}"
	ListStart                 = "["
	ListEnd                   = "]"
	String                    = "String"
	Colon                     = ":"
	Comma                     = ","
	Num                       = "Num"
	Bool                      = "Bool"
	Null                      = "Null"
)
