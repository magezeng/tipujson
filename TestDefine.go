package TipuJson

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
