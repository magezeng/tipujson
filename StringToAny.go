package TipuJson

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToInt64(arg string) (value int64, err error) {
	value, err = strconv.ParseInt(arg, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s转Int64类型失败", arg))
		return
	}
	return
}

func StringToUInt64(arg string) (value uint64, err error) {
	value, err = strconv.ParseUint(arg, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s转UInt64类型失败", arg))
		return
	}
	return
}

func StringToBool(arg string) (value bool, err error) {
	value, err = strconv.ParseBool(arg)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s转Bool类型失败", arg))
		return
	}
	return
}

func StringToFloat64(arg string) (value float64, err error) {
	value, err = strconv.ParseFloat(arg, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s转Float64类型失败", arg))
		return
	}
	return
}
