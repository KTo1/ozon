package main

import (
	"fmt"
	"github.com/KTo1/ozon/O/mapper"
)

type A struct {
	Name    string
	Code    int
	StructC *C
}

type B struct {
	Name    string
	Code    int
	StructC *C
}

type C struct {
	NameC string
	CodeC int
}

func main() {
	a := &A{
		Name: "Jonh",
		Code: 1,
		StructC: &C{
			NameC: "JonhC",
			CodeC: 2,
		},
	}

	b := &B{
		Name: "",
		Code: 0,
	}

	err := mapper.CopyFieldByPath(a, "Name", b, "Name")
	if err != nil {
		fmt.Println(err)
	}

	err = mapper.CopyFieldByPath(a, "StructC.NameC", b, "StructC.NameC")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
	fmt.Println(b)

	aa := []int{10, 20, 30, 40}
	bb := aa[:2]         // [10 20]
	cc := append(bb, 99) // что происходит с a?
	aa[1] = 777          // влияет ли это на b и c?
	cc = append(cc, 555) // а теперь?

	fmt.Println("a:", aa)
	fmt.Println("b:", bb)
	fmt.Println("c:", cc)
}
