package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var a int
	//a = 9
	//fmt.Println(a)

	a := 9 //value 단축어
	var b float32 = 2.71
	c := 'Z'
	d, e := 4.10, 3.99
	f := "문자열"
	fmt.Println(a, reflect.TypeOf(a), b, reflect.TypeOf(b), c, reflect.TypeOf(c), d, reflect.TypeOf(d), e, reflect.TypeOf(e), f, reflect.TypeOf(f))

	var g int
	var h float32
	//var i bool
	var j string
	fmt.Println("%s %d %f\n", j, g, h)

	//fmt.Println(reflect.TypeOf('Z'))
	//fmt.Println('A', 'a', '0', '김') //65, 97, 48, 44608
	//fmt.Println(math.Ceil(3.17))
	//fmt.Println(math.Floor(3.17))
	//fmt.Println(strings.Title("오픈소스 프로그래밍~"))
	//fmt.Println(strings.Title("open source\tProgramming~~\n\"go\"!"))
}
