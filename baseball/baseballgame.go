package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Make number for play
	gameNum := makeNumber()
	fmt.Println("Game Start!", gameNum)

	// get input from player
	inputNum := getInputNumber()

	// show result
	showResult(inputNum)

}

func makeNumber() [3]int {
	var gameNum [3]int
	for i := 0; i < 3; i++ {
		for {
			rand.Seed(time.Now().UnixNano())
			tmpNum := rand.Intn(9) + 1
			if checkIfUsableNum(tmpNum, gameNum) {
				gameNum[i] = tmpNum
				break
			}
		}
	}
	return gameNum
}

func checkIfUsableNum(num int, gameNum [3]int) bool {
	for i := 0; i < 3; i++ {
		if gameNum[i] == num {
			return false
		}
	}
	return true
}

func getInputNumber() [3]int {
	fmt.Println("숫자를 입력하세요.")
	var resultNum [3]int
	var isOk bool
	for {
		var inputNum int
		_, err := fmt.Scanf("%d", &inputNum)

		if err != nil {
			isOk, resultNum = validateInput(inputNum)
			if isOk {
				break
			}
			fmt.Println("inputNum=", inputNum)
		} else {
			log.Print(err)
		}
	}

	return resultNum
}

func validateInput(inputNum int) (bool, [3]int) {
	var resultNum [3]int
	if 0 < inputNum && inputNum <= 999 {
		i := 0
		for inputNum > 9 {
			resultNum[i] %= 10
			inputNum /= 10
			i++
		}
		return true, resultNum
	} else {
		fmt.Println("잘못 입력하셨습니다.")
		return false, resultNum
	}
}

func showResult(inputNum [3]int) {

	fmt.Println("inputNum=", inputNum)
}
