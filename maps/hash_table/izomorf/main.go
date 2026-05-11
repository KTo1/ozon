//Даны две строки s и t. Определите, являются ли они изоморфными.
//Две строки называются изоморфными, если можно заменить каждый символ строки s на некоторый символ так, чтобы получить строку t, при этом:
//Все вхождения одного символа из s должны заменяться на один и тот же символ.
//Разные символы s не могут заменяться на одинаковый символ t.
//Порядок символов сохраняется.

//"egg"  ,"add"	  true	e→a, g→d
//"foo"  ,"bar"	  false	o должен стать одновременно a и r
//"paper","title" true	p→t, a→i, e→l, r→e
//"ab"   ,"aa"	  false	a→a, b→a — два разных символа в один
//"badc" ,"baba"  false	b→b, a→a, d→b — конфликт (d и b в b)

package main

import "fmt"

func main() {
	fmt.Println(izomorf("egg", "add"))
	fmt.Println(izomorf("foo", "bar"))
	fmt.Println(izomorf("paper", "title"))
	fmt.Println(izomorf("ab", "aa"))
	fmt.Println(izomorf("badc", "baba"))
}

func izomorf(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)
	for i := 0; i < len(str1); i++ {
		if val, ok := map1[str1[i]]; ok {
			if val != str2[i] {
				return false
			}
		}

		if val, ok := map2[str2[i]]; ok {
			if val != str1[i] {
				return false
			}
		}

		map1[str1[i]] = str2[i]
		map2[str2[i]] = str1[i]
	}

	return true
}
