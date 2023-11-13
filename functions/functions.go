package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil //소수 판정 값, 정상데이터
}

func prime(n int) {	// 소수 판정 값 출력 함수
	p, err := isPrime(n)
	if err != nil { //err 체크
		fmt.Println(err)
		os.Exit(0)
	}

	if p {
		fmt.Println(n, "는(은) 소수입니다.")
	} else {
		fmt.Println(n, "는(은) 소수가 아닙니다.")
	}
}

func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil { //err 체크
			fmt.Println(err)
			os.Exit(0)
		}

		if p {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("\n")
}

// 소수 판정 및 구간 소수 판정 프로그램 v1.7
func main() {
	var menu int

	for true {
		fmt.Print("MENU 1) 소수판정 2) 구간 소수판정 0) 종료 : ")
		_, err := fmt.Scanln(&menu)

		if err != nil { //err 체크
			log.Fatal(err)
		}

		switch menu {
		case 1:
			var in int
			fmt.Print("정수 입력 : ")
			_, err := fmt.Scanln(&in)

			if err != nil { //err 체크
				log.Fatal(err)
			}
			prime(in)

		case 2:
			var n1, n2 int
			fmt.Print("정수 입력 : ")
			_, err := fmt.Scanln(&n1, &n2)

			if err != nil { //err 체크
				log.Fatal(err)
			}
			primeRange(n1, n2)

		default:
			fmt.Println("프로그램을 종료합니다.")
			os.Exit(1)
		}
	}

}

/*
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {	// Prime 변수 삭제, true/false로 return
	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil //소수 판정 값, 정상데이터
}

// 구간 소수 판정 프로그램 v1.4 : isPrime 함수 안의 변수를 하나 줄이고 return구문 추가, break 제거
func main() {
	var a, b int
	//2, 20
	//2 3 5 7 11 13 ... 19

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&a, &b)

	if err != nil { //err 체크
		log.Fatal(err)
	}

	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil { //err 체크
			fmt.Println(err)
			os.Exit(0)
		}

		if p {
			fmt.Print(i, " ")
		}
	}
}
*/

/*
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	prime := true

	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
		}
	}

	return prime, nil //소수 판정 값, 정상데이터
}

// 구간 소수 판정 프로그램 v1.3
func main() {
	var a, b int
	//2, 20
	//2 3 5 7 11 13 ... 19

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&a, &b)

	if err != nil { //err 체크
		log.Fatal(err)
	}

	if a > b {	// 입력 순서 상관 없이 작은 수 부터 큰 수 까지 작업
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil { //err 체크
			fmt.Println(err)
			os.Exit(0)
		}

		if p {
			fmt.Print(i, " ")
		}
	}
}
*/

/*
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	prime := true

	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
		}
	}

	return prime, nil //소수 판정 값, 정상데이터
}

// 소수 판정 프로그램 v1.1 : 함수 적용, error 리턴
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil { //err 체크
		log.Fatal(err)
	}

	p, err := isPrime(number)
	if err != nil { //err 체크
		fmt.Println(err)
		os.Exit(0)
	}

	if p {
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
	"os"
)

// isPrime() 함수 생성 후 적용
func isPrime(n int) bool {
	prime := true

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
		}
	}

	return prime //turn 리턴이면 소수, false면 소수
}

// 소수 판정 프로그램 v1.0 : 함수 적용
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil { //err 체크
		log.Fatal(err)
	}

	if number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
		os.Exit(0)
	}

	if isPrime(number) { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/* after (multi return)
package main

import "fmt"

func processScore(kor int, eng int, math int) (int, int) {
	totalScore := kor + eng + math
	average := totalScore / 3

	return totalScore, average	// 값을 두개 반환함
}

func main() {
	var t int
	var a int

	t, a = processScore(100, 90, 93)
	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길동", t, a)

	t, a = processScore(89, 91, 92)
	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길순", t, a)

}
*/

/* after
package main

import "fmt"

// 함수 생성
func processScore(name string, kor int, eng int, math int) {
	totalScore := kor + eng + math
	average := totalScore / 3

	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
}

func main() {
	processScore("홍길동", 100, 90, 93)

	processScore("홍길동", 89, 91, 92)

}
*/
/* before
package main

import "fmt"

func main() {
	kor := 100
	eng := 90
	math := 93
	name := "홍길동"

	fmt.Println(name, "의 총점은 ", (kor + eng + math), "입니다. 평균은 ", (kor+eng+math)/3.0, "입니다.")

	kor = 99
	eng = 91
	math = 92
	name = "홍길순"

	fmt.Println(name, "의 총점은 ", (kor + eng + math), "입니다. 평균은 ", (kor+eng+math)/3.0, "입니다.")

}
*/

/*
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello~")
}

func main() {
	sayHello()
}
*/
