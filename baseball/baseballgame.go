package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	errorMsgInput     = "0~9 사이의 겹치지않는 3자리 숫자를 입력하세요."
	errorMsgWrong     = "잘못 입력하셨습니다."
	errorMsgDuplicate = "겹치지않는 3자리 숫자를 입력하세요."
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
	gameNum := [3]int{-1, -1, -1}
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		for {
			tmpNum := rand.Intn(10)
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
	fmt.Println(errorMsgInput)
	var resultNum [3]int

outloop:
	for {
		var input string
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println(errorMsgWrong)
			continue
		}

		if len(input) < 3 || 3 < len(input) {
			fmt.Println(errorMsgWrong)
			continue
		}

		inputNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(errorMsgWrong)
			continue
		}

		if 0 < inputNum && inputNum <= 999 {
			i := 0
			for inputNum > 0 {
				tmpNum := inputNum % 10
				if !checkIfUsableNum(tmpNum, resultNum) {
					fmt.Println(errorMsgDuplicate)
					continue outloop
				}
				resultNum[i] = tmpNum
				inputNum /= 10
				i++
			}
			resultNum[0], resultNum[2] = resultNum[2], resultNum[0]
			break
		} else {
			fmt.Println(errorMsgDuplicate)
		}
	}
	return resultNum
}

func showResult(inputNum [3]int) {

	fmt.Println("inputNum=", inputNum)
}
