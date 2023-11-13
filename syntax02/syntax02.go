package main

import (
	"fmt"
	//"reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	//변수명은 영문자로 시작해야된다.
	//영문 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
	// var e string
	// var d bool
	// var c rune
	// var b float64 //값 저장 안하면 null이 아닌 '0'이 들어감
	// var a int
	// var studentId string
	// var i int32
	// //a := 7

	// //naming convention
	// fmt.Println(studentId)
	// fmt.Println(i)

	// //zero value
	// fmt.Println(a) //0
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d) //bool type의 기본값은 false
	// fmt.Println(e) //string type은 ""

	// fmt.Printf("%T\n", c)
	// fmt.Printf("%T\n", a)

	// fmt.Println(reflect.TypeOf(d))
	// fmt.Println(reflect.TypeOf(e))
}
