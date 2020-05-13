package main

import (
	"fmt"
)
	
type StructA struct{
	val1 string
	val2 string
}

// func (s *StructA) AAA(){
// 	fmt.Println(s.val1)
// }

func (s StructA) AAA(){
	s.val2 = "333"
	fmt.Println(s.val2)
}

func main(){
	
	s := &StructA{val1:"111", val2:"222"} 
	// s := StructA{val1:"111", val2:"222"}
	s.AAA()
	fmt.Println(s)

}