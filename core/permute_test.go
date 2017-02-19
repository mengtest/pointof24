package core

import (
	"testing"
	"strings"
	"fmt"
)

func TestPermute(t *testing.T) {
	in := "0 1 2 3"
	result := Permute(strings.Fields(in))
	//dictSort(result)
	//s := format(result)
	fmt.Println(result)
}

func TestPermuteOperator(t *testing.T) {
	result := PermuteOperator(4)
	//dictSort(result)
	//s := format(result)
	fmt.Println(result)
}
