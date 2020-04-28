package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	msgInput          = "0~9 사이의 겹치지않는 3자리 숫자를 입력하세요."
	errorMsgWrong     = "잘못 입력하셨습니다."
	errorMsgDuplicate = "겹치지않는 3자리 숫자를 입력하세요."
)

func startGame() int {
	gameNum := makeNumber()
	fmt.Println("Game Start!", gameNum)
	fmt.Println(msgInput)

	for {
		// get input from player
		inputNum := getInputNumber()
		strike, ball := checkResult(inputNum, gameNum)
		if strike > 0 || ball > 0 {
			if strike == 3 {
				fmt.Println("You WIN!!")
				return 1
			} else {
				fmt.Printf("%dS %dB\n", strike, ball)
			}
		} else {
			fmt.Println("OUT!!")
		}
	}

}

func main() {
	outloop:
	for {
		if gameOver := startGame(); gameOver == 1 {
			fmt.Println("다른 게임을 시작하시겠습니까?\n1.네\n2.아니오")
			for {
				var restart int
				if _, err := fmt.Scanln(&restart); err == nil {
					if restart == 1 {
						fmt.Println()
						continue outloop
					} else if restart == 2 {
						fmt.Println("Game Over..")
						break outloop
					} else {
						fmt.Println("다시 입력하세요.")
					}
				} else {
					fmt.Println("다시 입력하세요.")
				}
			}
		}
	}
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
	var resultNum = [3]int{-1, -1, -1}

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

		for i, val := range input {
			if tmp := int(val) - '0'; -1 < tmp && tmp < 10 {
				if !checkIfUsableNum(tmp, resultNum) {
					fmt.Println(errorMsgDuplicate)
					continue outloop
				}
				resultNum[i] = tmp
			} else {
				fmt.Println(errorMsgWrong)
				continue outloop
			}
		}
		break
	}
	return resultNum
}

func checkResult(inputNum [3]int, gameNum [3]int) (int, int) {
	var strike int
	var ball int
	//fmt.Println(inputNum, gameNum)
	for i, val := range inputNum {
		for j, gameVal := range gameNum {
			if i == j && val == gameVal {
				strike++
			} else if i != j && val == gameVal {
				ball++
			}
		}
	}

	return strike, ball
}
