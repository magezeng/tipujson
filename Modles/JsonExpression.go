package Modles

type JsonExpressionType string

const (
	JsonExpressionTypeMapStart   JsonExpressionType = "{"
	JsonExpressionTypeMapEnd                        = "}"
	JsonExpressionTypeListStart                     = "["
	JsonExpressionTypeListEnd                       = "]"
	JsonExpressionTypeStringMark                    = "\""
	JsonExpressionTypeComma                         = ","
	JsonExpressionTypeColon                         = ":"
	JsonExpressionTypeValue                         = "Value"
)

type JsonExpression struct {
	Type          JsonExpressionType
	Content       string
	StartPosition int
	EndPosition   int
	Pre           *JsonExpression
	Next          *JsonExpression
}
