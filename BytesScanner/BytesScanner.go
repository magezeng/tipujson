package BytesScanner

import (
	"errors"
	"strings"
)

type BytesScanner struct {
	Bytes  []byte
	Cursor int
}

func (scanner *BytesScanner) GetMarkString(position int, mark string) (errString string) {
	startPosition := position - 20
	if startPosition <= 0 {
		startPosition = 0
	}
	endPosition := position + 20
	if endPosition >= len(scanner.Bytes) {
		endPosition = len(scanner.Bytes)
	}
	return string(scanner.Bytes[startPosition:position]) + mark + string(scanner.Bytes[position:endPosition])
}
func (scanner *BytesScanner) CurrentValue() (currentByte byte) {
	if scanner.Cursor >= len(scanner.Bytes) {
		panic(errors.New("字节扫描器异常,访问超出长度限制"))
	}
	currentByte = scanner.Bytes[scanner.Cursor]
	return
}

func (scanner *BytesScanner) GetNextValue() (currentByte byte, err error) {
	if scanner.Cursor+1 >= len(scanner.Bytes) {
		panic(errors.New("字节扫描器异常,访问超出长度限制"))
	}
	currentByte = scanner.Bytes[scanner.Cursor+1]
	return
}

func (scanner *BytesScanner) BackMove() {
	scanner.BackMoveDistance(1)
}
func (scanner *BytesScanner) BackMoveDistance(distance int) {
	scanner.Cursor += distance
}

func (scanner *BytesScanner) BackMoveTo(to byte) {
	for ; scanner.CurrentValue() != to; scanner.Cursor++ {
	}
}

func (scanner *BytesScanner) BackMoveToNotNull() {
	for ; scanner.CurrentValue() == ' ' || scanner.CurrentValue() == '\t' || scanner.CurrentValue() == '\n'; scanner.BackMove() {
	}
}

func (scanner *BytesScanner) GetSubStringTo(endPosition int) string {
	return string(scanner.Bytes[scanner.Cursor:endPosition])
}

func (scanner *BytesScanner) ScanString() (result string) {

	scanner.BackMoveToNotNull()
	//字符串必须以引号开始
	if scanner.CurrentValue() == '"' {
		scanner.BackMove()
	} else {
		panic(errors.New(""))
	}
	stringStartPosition := scanner.Cursor
	scanner.BackMoveTo('"')
	result = string(scanner.Bytes[stringStartPosition:scanner.Cursor])
	scanner.BackMove()
	return
}

func (scanner *BytesScanner) ScanNumberString() (result string, isBool bool) {
	scanner.BackMoveToNotNull()
	for {
		if (scanner.CurrentValue() >= '0' && scanner.CurrentValue() <= '9') || scanner.CurrentValue() == '.' {
			isBool = false
			break
		} else if scanner.CurrentValue() == 'T' || scanner.CurrentValue() == 'F' || scanner.CurrentValue() == 't' || scanner.CurrentValue() == 'f' {
			isBool = true
			break
		} else {
			panic(errors.New(""))
		}
	}

	stringStartPosition := scanner.Cursor
	if isBool {
		if strings.ToLower(string(scanner.GetSubStringTo(scanner.Cursor+4))) == "true" {
			scanner.BackMoveDistance(4)
		} else if strings.ToLower(string(scanner.GetSubStringTo(scanner.Cursor+5))) == "false" {
			scanner.BackMoveDistance(5)
		}
	} else {
		for {
			scanner.BackMove()
			if !(scanner.CurrentValue() >= '0' && scanner.CurrentValue() <= '9') || scanner.CurrentValue() == '.' {
				break
			}
		}
	}
	result = string(scanner.Bytes[stringStartPosition:scanner.Cursor])
	return
}
