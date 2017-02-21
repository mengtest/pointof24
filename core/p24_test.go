package core

import "testing"

var allTests []*Point24 = []*Point24{
	NewPoint24([]int{1,2,3,4}, 24),
}

func TestGetInputStrings(t *testing.T) {
	for _,v := range allTests {
		t.Log(v.GetInputStrings())
	}
}

func TestInputAtoi(t *testing.T) {
	for _,v := range allTests {
		t.Log(v.InputAtoi(v.GetInputStrings()))
	}
}

func TestCalcAll(t *testing.T) {
	for _,v := range allTests {
		v.CalcAll()
		//v.Display()
	}
}
