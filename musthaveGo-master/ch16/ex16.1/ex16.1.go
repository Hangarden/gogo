// ch16/ex16.1/ex16.1.go
package main

//패키지는 코드를 묶는 가장 큰 단위
import ( // ❶ 둘 이상의 패키지는 소괄호로 묶어줍니다.
	"fmt"
	"math/rand" // ❷ 패키지명은 rand입니다.
)

//외부 노출 여부는 구조체명, 변수명, 함수명의 첫글자가 대문자 소문자인지를 보면 된다. 대문자 : 노출 소문자 : 노출 X

func main() {
	fmt.Println(rand.Int()) // ❸ 랜덤한 숫자값을 출력합니다.
}

//main패키지에는 main함수, 그 외함수
//구조체, 그외로 구성되어있다.
