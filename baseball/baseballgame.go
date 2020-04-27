package main

import (
	"fmt"
	"math/rand"
)

func main(){
	// Make number for play
	gameNum := makeNumber()
	fmt.Println("gameNumber=",gameNum)
	// get input from player
	inputNum := getInputNumber()

	// show result
	showResult(inputNum)
}

func makeNumber() [3]int{
	var gameNum [3]int
	for i := 0 ; i < 3; i++{
		gameNum[i] = rand.Intn(10) + 1
	}
	return gameNum
}

func getInputNumber() [3]int{

	return [3]int{}
}

func showResult(inputNum [3]int)  {

	fmt.Println("inputNum=", inputNum)
}
