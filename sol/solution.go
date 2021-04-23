package sol

import (
	"strconv"
	"unicode"
)

func isNumber(s string) (bool, int) {
	val, err := strconv.Atoi(s)
	return err == nil, val
}

func decodeString(s string) string {
	result := ""
	inputBuffer := ""
	lenS := len(s) - 1
	tokenStack := make([]string, 0)
	numberStack := make([]int, 0)
	for idx, r := range s {
		isNumValue := unicode.IsNumber(r)
		switch r {
		case '[':
			if inputBuffer != "" { // do Push
				isNum, val := isNumber(inputBuffer)
				if isNum {
					numberStack = append(numberStack, val)
				} else {
					tokenStack = append(tokenStack, inputBuffer)
				}
				inputBuffer = ""
			}
		case ']':
			top := len(tokenStack) - 1
			topNum := len(numberStack) - 1
			num := numberStack[topNum]
			targetToken := ""
			appendToken := ""
			appendResult := ""
			if inputBuffer != "" {
				targetToken = inputBuffer
				appendToken = tokenStack[top]
				tokenStack = tokenStack[:top]
				inputBuffer = ""
			} else {
				targetToken = tokenStack[top]
				appendToken = tokenStack[top-1]
				tokenStack = tokenStack[:top-1]
			}
			for i := 0; i < num; i++ {
				appendResult += targetToken
			}
			if idx == lenS {
				result = appendToken + appendResult
			} else {
				tokenStack = append(tokenStack, appendToken+appendResult)
			}
			numberStack = numberStack[:topNum]
		default:
			if isNumValue { // do Push if inputBuffer not ""
				isNum, _ := isNumber(inputBuffer)
				if isNum == false && inputBuffer != "" {
					tokenStack = append(tokenStack, inputBuffer)
					inputBuffer = ""
				}
			}
			inputBuffer += string(r)
		}
	}
	return result
}
