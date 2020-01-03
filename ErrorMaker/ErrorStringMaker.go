package ErrorMaker

import (
	"fmt"
	"regexp"
	"runtime"
)

func GetErrorStringFromErr(err error) string {
	_, file, line, _ := runtime.Caller(1)
	re, _ := regexp.Compile(`(?:/|\\)([^(?:\.|/|\\)]*)\.go`)
	result := re.FindStringSubmatch(file)
	fileName := result[1]
	return fmt.Sprint(fileName, ".", line, ":", err)
}
