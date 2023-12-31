package main

import (
	"fmt"
	"math/rand"
	"time" //Seed 생성용 패키지
)

// 난수 추출된 수의 소수 판정 프로그램 v0.6
// 소수 : 1과 자기 자신 말고는 나누어 떨어지지 않는 수(0과 1은 제외)
func main() {
	//seed 설정(실행 할 때마다 랜덤 값 설정을 위함)
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2 //0과 1 제외, 2~151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

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

/*
// 난수 추출된 수의 소수 판정 프로그램 v0.4
func main() {
	//seed 설정(실행 할 때마다 랜덤 값 설정을 위함)
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			//count = count + 1
		}
		fmt.Print(i, " ")
	}

	if isPrime { //비교 연산자 제거(true는 필요 없음)
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
// 난수 추출된 수의 소수 판정 프로그램 v0.3
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false	//count 변수를 지우고 isPrime 변수를 이용해 간략화
		}
	}

	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"time" //Seed 생성용 패키지
)
*/

/*
// 난수 추출된 수의 소수 판정 프로그램 v0.2
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ { // 1과 자기 자신일때 loop 돌지 않음
		if number%i == 0 {
			//count++
			count = count + 1
		}
	}

	if count == 0 {	//1과 자신을 제외한 값이 0
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)
*/

/*
// 난수 추출된 수의 소수 판정 프로그램 v0.1
// 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {	// count 값 : 1과 자기 자신
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	for i := 1; i < 6; i++ {
		dice := rand.Intn(6) + 1
		fmt.Println(dice)
	}
}
*/
