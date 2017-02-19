package main

import "strings"

func main(){
	s := "4-(2+18/9*11)+(9/3)"
	println(strings.LastIndex(s, "("))
	println(strings.Index(s[strings.LastIndex(s, "("): ], ")")+strings.LastIndex(s, "("))
}
