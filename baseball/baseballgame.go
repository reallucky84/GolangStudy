package main

import (
	"fmt"
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
	var inputNum int
	fmt.Println("숫자를 입력하세요.")
	for{
		_, err := fmt.Scanf("%d", &inputNum)
		if err != nil {
			fmt.Println("inputNum=", inputNum)
			break
		}
	}
	var resultNum [3]int
	return resultNum
}

func showResult(inputNum [3]int) {

	fmt.Println("inputNum=", inputNum)
}
