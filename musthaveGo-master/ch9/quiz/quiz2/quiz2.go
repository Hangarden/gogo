package main

import "fmt"

func main() {

	//조건1 : 낮 최고 기온 25도 이상 강수 확률 80% 이상일 시
	//-> fmt.Println("덥고 비가 옵니다.")
	//조건2 : 낮 최고 기온 25도 이상 강수 확률 20% 이상일 시
	//-> fmt.Println("덥고 습합니다.")
	//조건3 : 낮 최고 기온 25도 이상 강수 확률 20% 미만일 시
	//-> fmt.Println("야외 활동 하기 좋습니다..")
	//조건4 : 낮 최고 기온 25도 이상이 아닐 시
	//-> fmt.Println("덥고 습합니다.")
	//조건5 : 강수확률 80% 이상일 시
	//-> fmt.Println("야외 활동 하기 좋습니다..")
	//

	temp := 30
	rain := 40
	if temp >= 25 && rain >= 80 {
		fmt.Println("덥고 비가 옵니다.")
	} else if temp >= 25 && rain >= 20 {
		fmt.Println("덥고 습합니다.")
	} else if temp >= 25 && rain < 20 {
		fmt.Println("야외 활동 하기 좋습니다..")
	} else if temp < 25 || rain >= 80 {
		fmt.Println("야외 활동 하기 좋지 않습니다..")
	} else {
		fmt.Println("좋은 날씨입니다.")
	}
}
