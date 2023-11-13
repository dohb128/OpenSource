package main

import "fmt"

/*
shadowing : 자료타입, 패키지, 함수명 등을 변수명으로 선언하면
더이상 자료형, 패키지, 함수로 작동하지 않음
*/
func main() {
	var test1 float64 = 9.1
	var test2 float64 = 7.9
	fmt.Println(test1 + test2)

	var univ string = "inha"
	fmt.Println(univ)

	var f1 string = "func"
	var f2 = append([]string{}, "함수")
	fmt.Println(f1)
	fmt.Println(f2)

	// 자료타입을 변수명으로 사용
	// var float64 float64 = 9.1
	// var test float64 = 7.9
	// fmt.Println(float64)

	// 패키지를 변수명으로 사용
	// var fmt string = "inha"
	// fmt.Println()

	// 함수를 변수명으로 사용
	// var append string = "func"
	// var f = append([]string{}, "함수")
}

/*
package main

import (
	"fmt"
	"log"
	"os"
)

// 입력(0과 1 처리)된 수의 소수 판정 프로그램 v0.9
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil { //err 체크
		log.Fatal(err)
	}

	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
		os.Exit(1)
	}

	isPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
		}
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
package main

import (
	"fmt"
	"log"
)

// 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.8
func main() {
	var number int
	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	//fmt.Println(n)
	if err != nil { //err 체크
		log.Fatal(err)
	}

	isPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
		}
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 입력된 수의 소수 판정 프로그램 v0.7
func main() {
	fmt.Print("정수 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil {
		log.Fatal(err) //err 발생시 메세지 출력 후 종료됨
	}

	in = strings.TrimSpace(in)      //공백 제거
	number, err := strconv.Atoi(in) //자동으로 설정해줌
	if err != nil {
		log.Fatal(err)
	}

	isPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
		}
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/
