package ErrorMaker

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"
)

func GetError(args ...interface{}) error {
	_, file, line, _ := runtime.Caller(1)
	re, _ := regexp.Compile(`(?:/|\\)([^(?:\.|/|\\)]*)\.go`)
	result := re.FindStringSubmatch(file)
	fileName := result[1]
	return errors.New(fmt.Sprint(append([]interface{}{fileName, ".", line, ":"}, args...)...))
}
