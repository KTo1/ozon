// на вход подается строка содержащая только скобки ([{}])
// нужно определить является ли строка валидной
// строка валдина если:
//
//  1. все открытые скобки закрыты скобками того же типа
//  2. скобки закрыты в правильном порядке
//  3. каждой закрывающей скобке соотвествует открытая скобка того же типа

package main

import (
	"fmt"
)

func main() {
	//fmt.Println(isValid("()"))
	fmt.Println(isValid("()(]"))
}

func isValid(s string) bool {
	stack := []string{}
	openMap := make(map[string]string)
	closeMap := make(map[string]string)

	openMap["("] = ")"
	openMap["{"] = "}"
	openMap["["] = "]"

	closeMap[")"] = "("
	closeMap["}"] = "{"
	closeMap["]"] = "["

	for _, v := range s {
		symbol := string(v)
		if _, ok := openMap[symbol]; ok {
			stack = append(stack, symbol)
		}

		if closeSymbol, ok := closeMap[symbol]; ok {
			if stack[len(stack)-1] == closeSymbol {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
