package main

import (
	"bufio"
	"fmt"
	"log" //Fatal function -> err 찍고 시스템 종료
	"os"
	"strconv" //TrimSpace
	"strings" //ParseInt
)

func main() {
	fmt.Print("단 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil {
		log.Fatal(err) //err 발생시 메세지 출력 후 종료됨
	}

	in = strings.TrimSpace(in) //공백 제거
	//dan, err := strconv.ParseInt(in, 10, 32)	//수동 설정
	dan, err := strconv.Atoi(in) //자동으로 설정해줌
	if err != nil {
		log.Fatal(err)
	}

	/*for 문 버전
	for i := 1; i < 10; i++ {
		// fmt.Println(dan, " X ", i, " = ", (int(dan) * i))	//dan의 비트가 안맞아 int 다시 지정해줌
		fmt.Println(dan, " X ", i, " = ", (dan * i))
	}
	*/

	//다른 언어의 while문 버전(go에는 while문이 없음)
	i := 1
	for i < 10 {
		fmt.Println(dan, " X ", i, " = ", (dan * i))
		i++
	}

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("숫자 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, _ := rd.ReadString('\n')	//ReadString이 변수와 에러를 같이 가져와서 변수 두개 필요
// 	fmt.Println(in)

// }
