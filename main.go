package main

import (
	"pointof24/core"
	"math/rand"
)

func main(){
	result := 24
	//inputs := []int{5,5,6,6}
	//inputs := []int{3,2,3,4}
	//inputs := []int{3,3,6,6}
	//inputs := []int{12,13,16,rand.Intn(100) + 1}
	//inputs := []int{6,6,2,2}
	//inputs := []int{2,6,2,6}
	//p24obj := core.NewPoint24(inputs, result)
	//p24obj.Display()

	//test
	for {
		inputs := make([]int, 4)
		for i := 0; i < 4; i++ {
			inputs[i] = rand.Intn(100) + 1
		}
		p24obj := core.NewPoint24(inputs, result)
		p24obj.Display()
	}
}

