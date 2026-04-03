// var s *string тут есть указание на тип данных
// сама переменная будет нил, но если это забросить через интерфейс, то указани на тип в интерфейсе будет.
// и переменная будет не равно нил

package main

import "fmt"

type as struct {
}

func check(i interface{}) {
	fmt.Println(i == nil)
}

func main() {
	var s *string
	var a *as

	fmt.Println(s == nil) // нил

	fmt.Println(a == nil) // нил
	check(a)              // не нил

	var i interface{}

	fmt.Println(i == nil) // нил

	i = s
	fmt.Println(i == nil) // не нил

	p := ""
	s = &p
	s = nil
	fmt.Println(s == nil) //нил
}
