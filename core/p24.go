package core

import (
	"strconv"
	"fmt"
)

type Point24 struct {
	inputNums []int
	result int
	//缓存所有满足条件的表达式
	stored []*Expression
}

func NewPoint24(inputNums []int, result int) *Point24{
	return &Point24{inputNums, result, make([]*Expression, 0)}
}

func (this *Point24)GetInputStrings() []string{
	results := make([]string, len(this.inputNums))
	for i, v := range this.inputNums{
		results[i] = strconv.Itoa(v)
	}
	return results
}

func (this *Point24)InputAtoi(inputs []string) []int{
	results := make([]int, len(inputs))
	for i, v := range inputs{
		if vint, ok := strconv.Atoi(v); ok == nil{
			results[i] = vint
		}
	}
	return results
}

/*
遍历所有情况
*/
func (this *Point24)CalcAll(){
	allInput := Permute(this.GetInputStrings())
	allOperations := PermuteOperator(len(this.inputNums) - 1)
	for _, input := range allInput{
		for _, opt := range allOperations{
			this.AddKuoHao(this.InputAtoi(input), opt, make([][2]int, len(input)), 0, len(input) - 1)
		}
	}
}

/*
遍历所有加括号的情况
*/
func (this *Point24)AddKuoHao(inputNums []int, operations []string, parenthesis [][2]int, start, end int){
	//计算表达式的值
	expObj, err := NewExpression(inputNums, operations, parenthesis)
	if err == nil{
		//fmt.Println(expObj.String())
		v, _ := expObj.Value()
		if v == this.result{
			//this.stored = append(this.stored, expObj)
			this.AddNoRepeat(expObj)
		}
	}else{
		fmt.Println(err)
		fmt.Println("错误的表达式: ")
		fmt.Println(inputNums)
		fmt.Println(operations)
		fmt.Println(parenthesis)
	}

	if end - start <= 1{
		return
	}else{
		for i := start; i < end; i++ {
			parenthesisCopy := make([][2]int, len(parenthesis))
			copy(parenthesisCopy, parenthesis)
			//右边部分
			if i+1 < end{
				//加左括号
				parenthesisCopy[i+1][0] += 1
				//加右括号
				parenthesisCopy[end][1] += 1
			}
			//左边部分
			if i > start{
				//加左括号
				parenthesisCopy[start][0] += 1
				//加右括号
				parenthesisCopy[i][1] += 1
			}
			//递归加括号
			this.AddKuoHao(inputNums, operations, parenthesisCopy, start, i)
			this.AddKuoHao(inputNums, operations, parenthesisCopy, i+1, end)
		}
	}
}

/*
去重
*/
func (this *Point24)AddNoRepeat(exp *Expression){
	for _, expobj := range this.stored{
		if expobj.Equal(exp){
			return
		}
	}
	this.stored = append(this.stored, exp)
}

func (this *Point24)Display(){
	this.CalcAll()
	if len(this.stored) > 0{
		fmt.Println(fmt.Sprintf("%v = %d. 所有满足条件表达式: ", this.inputNums, this.result))
		for _, exp := range this.stored{
			exp.Display()
		}
	}else{
		fmt.Println(fmt.Sprintf("%v = %d. 无解", this.inputNums, this.result))
	}
}
