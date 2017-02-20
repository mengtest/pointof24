package main

import (
	"pointof24/core"
)

func main(){
	result := 24
	//inputs := []int{5,5,6,6}
	//inputs := []int{3,2,3,4}
	//inputs := []int{3,3,6,6}
	//inputs := []int{12,13,16,16}
	inputs := []int{6,6,2,2}
	//inputs := []int{2,6,2,6}
	p24obj := core.NewPoint24(inputs, result)
	p24obj.Display()
}

