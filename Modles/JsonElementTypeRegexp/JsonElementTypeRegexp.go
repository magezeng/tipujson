package JsonElementTypeRegexp

type JsonElementTypeRegexp string

const (
	MapStart  JsonElementTypeRegexp = `^\s*(\{)`
	MapEnd                          = `^\s*(\})`
	ListStart                       = `^\s*(\[)`
	ListEnd                         = `^\s*(\])`
	String                          = `^\s*\"((?:[^(?:\"|\\)]|\\.)*)\"`
	Colon                           = `^\s*(:)`
	Comma                           = `^\s*(,)`
	Num                             = `^\s*((?:-|\+)?\d+(?:\.\d*)?)`
	Bool                            = `^\s*((?:F|f|T|t)(?:alse|ALSE|rue|RUE))`
	Null                            = `^\s*((?:N|n)(?:ull|ULL))`
)
