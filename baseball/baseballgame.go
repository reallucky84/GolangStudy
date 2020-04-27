package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Make number for play
	gameNum := makeNumber()
	fmt.Println("gameNumber=", gameNum)

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

	return [3]int{}
}

func showResult(inputNum [3]int) {

	fmt.Println("inputNum=", inputNum)
}
