package core

import (
	"testing"
	"strings"
	"fmt"
)

func TestPermute(t *testing.T) {
	in := "1 2 3"
	result := Permute(strings.Fields(in))
	fmt.Println(result)
}

func TestPermuteOperator(t *testing.T) {
	result := PermuteOperator(3)
	fmt.Println(result)
}
