package main

func main() {
	s := "test"
	println(s[0]) // тут будут байты символа Т - 115
	// s[0] = "R" -строки неизменяемы тип так сделать нельзя

	//как изменить? Привести в рунам
	r := []rune(s)
	r[0] = 'R'

	println(string(r))
}
