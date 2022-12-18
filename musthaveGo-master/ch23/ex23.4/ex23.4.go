// ch23/ex23.4/ex23.4.go
package main //에러를 감싸 새로운 에러를 만들어야 할 때 에러 래핑을 사용

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // ❶ 스캐너 생성 한단어씩 끊어 읽는 bufio패키지의 스캐너
	// io.Reader 인터페이스를 인수로 받아 string을 string.NewReader를 통해 형변환을 시도함 위 구문은 한 단어씩 끊어 읽고자할 때 자주 사용
	scanner.Split(bufio.ScanWords) // ❷ ScanWords를 사용하여한 단어씩 끊어읽기 ScanLine을 사용시 한 줄 씩 끊어 읽음

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err) // ➏ 에러 감싸기
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환합니다.
// 변환된 숫자, 읽은 글자수, 에러를 반환합니다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // ❸ 단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) // ❹ 문자열을 숫자로 변환 Itoa는 숫자를 문자열로 바꾸어줌
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err) // ➎ 원래 에러는 NumberError 타입 에러 반환
		// Errorf, %w를통해 에러 감싸기
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { // ➐ 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 10")
	readEq("123 51 abc")
}
