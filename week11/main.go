package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Printf(primes, primes[1])

	// var primes [3]int = [3]int{2, 3, 5}
	// fmt.Printf(primes, primes[1])	// primes 배열을 배열 리터럴로 초기화

	primes := [3]int{2, 3, 5} // 단축 연산자
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) // boolean 타입의 제로 값, false

	i := 0
	//for i < 4 {	// 4번 실행할 primes 객체가 부족해 오류
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}

	for _, prime := range primes { //값을 출력하려 했으나 인덱스가 출력됨
		fmt.Println(prime)
	}

	fmt.Println("--------")
	fmt.Printf("%#v\n", test)   // test의 타입, 크기, 값 출력
	fmt.Println(test)           // test의 값 출력
	fmt.Printf("%#v\n", primes) // primes의 타입, 크기, 값 출력
	fmt.Println(primes)         // primes의 값 출력
}
