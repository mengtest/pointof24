package core

import (
	"errors"
	"strings"
	"fmt"
	"strconv"
)

/*
定义表达式结构及逻辑
*/

type Expression struct {
	//操作数从左到右
	optNumbers []int
	//操作符从左到右
	operators []string
	//括号对应操作数的左右括号
	parenthesis [][2]int
}

func NewExpression(optNumbers []int, operators []string, parenthesis [][2]int) (*Expression, error){
	//参数检查
	if len(optNumbers) != len(parenthesis) || len(optNumbers) != len(operators) + 1 {
		//TODO
		//其他检查
		return nil, errors.New("参数有误")
	}else{
		return &Expression{optNumbers, operators, parenthesis}, nil
	}
}

func (this *Expression)Display(){
	if v, err := this.Value(); err == nil{
		fmt.Println(fmt.Sprintf("%s=%d", this.String(), v))
	}else{
		fmt.Println(fmt.Sprintf("表达式：%s无法计算，原因：%s", this.String(), err))
	}

}

/*
表达式字符串
*/
func (this *Expression)String() string{
	r := make([]string, 0, len(this.optNumbers) * 4)
	lOptNum := len(this.optNumbers)
	for i, v := range this.parenthesis{
		//左括号
		for _, sLeft := range strings.Repeat("(", v[0]){
			r = append(r, string(sLeft))
		}
		//操作数
		r = append(r, strconv.Itoa(this.optNumbers[i]))
		//右括号
		for _, sRight := range strings.Repeat(")", v[1]){
			r = append(r, string(sRight))
		}
		//操作符
		if i + 1 < lOptNum{
			//不是最后一个操作数
			r = append(r, this.operators[i])
		}
	}
	return strings.Join(r, "")
}

/*
计算表达式的值(带括号)
思想: 循环遍历表达式字符串，每次结合一个最右边最内层的括号，直到表达式没有括号为止
*/
func (this *Expression)Value() (int, error){
	s := this.String()
	for {
		//每次遍历需要消除一对括号
		//先找左括号
		leftPos := strings.LastIndex(s, "(")
		if leftPos == -1{
			//没有了
			break
		}else{
			//找配对的右括号
			rightPos := strings.Index(s[leftPos:], ")") + leftPos
			//计算括号内表达式的值
			v, err := this.ValueOfNoParenthesis(s[leftPos + 1: rightPos])
			if err != nil{
				return 0, err
			}
			//s 重新整理
			s = s[: leftPos] + strconv.Itoa(v) + s[rightPos + 1:]
		}
	}
	return this.ValueOfNoParenthesis(s)
}

/*
计算表达式的值(不带括号)
思想:
    双队列思想求解带括号表达式，从左到右遍历表达式，操作数入数据队列，操作符入
符号队列，遇到作括号忽略，遇到右括号立即结合左边第一个操作符并且求出结果，两个
操作数出队列，结果入队列，操作符出队列,然后按照操作符优先级顺序结算结果
*/
func (this *Expression)ValueOfNoParenthesis(s string) (int, error){
	//fmt.Println(s)
	//操作数队列
	optNumbers := make([]int, 0)
	//操作符队列
	operators := make([]string, 0)
	//操作数缓存
	digistStr := []string{}
	//操作数是否为负数
	isminus := false
	for i, ss := range s{
		if _, ok := strconv.Atoi(string(ss)); ok == nil{
			//操作数
			//optNumbers = append(optNumbers, v)
			//支持多位操作数, 操作数缓存
			digistStr = append(digistStr, string(ss))
			//最后的操作数直接入队
			if i == len(s) - 1{
				//操作数缓存入队
				if v, ok := strconv.Atoi(strings.Join(digistStr, "")); ok == nil{
					digistStr = []string{}
					if isminus{
						v = 0 - v
						isminus = false
					}
					optNumbers = append(optNumbers, v)
				}
			}
		}else{
			//这里要处理负数，重要
			if string(ss) == "-" && len(digistStr) == 0{
				//操作数是负数
				isminus = true
				continue
			}
			//操作符
			operators = append(operators, string(ss))
			//操作数缓存入队
			if v, ok := strconv.Atoi(strings.Join(digistStr, "")); ok == nil{
				digistStr = []string{}
				if isminus{
					v = 0 - v
					isminus = false
				}
				optNumbers = append(optNumbers, v)
			}
		}

	}
	//遍历操作符队列，按照优先级结合
	//fmt.Println(optNumbers)
	//fmt.Println(operators)
	Loop1: for{
			for i, opt := range operators{
				if opt == "*" ||  opt == "/"{
					if opt == "/" && optNumbers[i+1] == 0{
						return 0, errors.New("error: integer divide by zero")
					}
					operators = append(operators[:i], operators[i+1:]...)
					optNumbersTemp := []int{}
					//防止越界
					if i + 2 < len(optNumbers){
						optNumbersTemp = append(optNumbersTemp, optNumbers[i+2:]...)
					}
					optNumbers = append(optNumbers[:i], this.Calc(optNumbers[i], optNumbers[i+1], opt))
					optNumbers = append(optNumbers, optNumbersTemp...)
					if i == len(operators){
						break Loop1
					}else{
						break
					}
				}
				if i == len(operators) - 1{
					break Loop1
				}
			}
			if len(operators) == 0{
				break Loop1
			}
		}
	//fmt.Println(this.optNumbers)
	//fmt.Println(this.operators)
	//fmt.Println(this.parenthesis)
	//fmt.Println(optNumbers)
	//fmt.Println(operators)
	if len(operators) > 0 {
		Loop2: for{
			for i, opt := range operators{
				operators = append(operators[:i], operators[i+1:]...)
				optNumbersTemp := []int{}
				//防止越界
				if i + 2 < len(optNumbers){
					optNumbersTemp = append(optNumbersTemp, optNumbers[i+2:]...)
				}
				optNumbers = append(optNumbers[:i], this.Calc(optNumbers[i], optNumbers[i+1], opt))
				optNumbers = append(optNumbers, optNumbersTemp...)
				if i == len(operators){
					break Loop2
				}else{
					break
				}
			}
		}
	}

	return optNumbers[0], nil
}

func (this *Expression)Calc(optD1, optD2 int, opt string) int{
	r := 0
	switch opt {
	case "+":
		r = optD1 + optD2
	case "-":
		r = optD1 - optD2
	case "*":
		r = optD1 * optD2
	case "/":
		r = optD1 / optD2
	}
	return r
}

/*
求取表达式相似度
*/
func (this *Expression)Equal(exp *Expression) bool{
	s1 := this.String()
	s2 := exp.String()
	if len(s1) != len(s2){
		return false
	}else{
		if s1 == s2{
			return true
		}
		for i, v := range this.operators{
			if v == exp.operators[i] && (v == "+" || v == "*"){
				//fmt.Println(exp.optNumbers)
				//fmt.Println(this.optNumbers)
				//fmt.Println(i)
				if this.optNumbers[i] == exp.optNumbers[i + 1] && this.optNumbers[i + 1] == exp.optNumbers[i]{
					//有可能是重复的
					if this.parenthesis[i][1] == exp.parenthesis[i][1] &&
						this.parenthesis[i + 1][0] == exp.parenthesis[i + 1][0]&&
						this.parenthesis[i][1]==0&&this.parenthesis[i + 1][0]==0{
						//左边的数没有右括号，右边的数没有左括号
						//判断其他部分是否相同
						matchs := [2]string {fmt.Sprintf("%d%s%d", this.optNumbers[i], v, this.optNumbers[i+1]),
							fmt.Sprintf("%d%s%d", exp.optNumbers[i], v, exp.optNumbers[i+1])}
						for _, m := range matchs{
							s1 = strings.Replace(s1, m, "", 1)
							s2 = strings.Replace(s2, m, "", 1)
						}
						if s1 == s2{
							return true
						}else{
							return false
						}
					}
				}
			}
		}
		return false
	}
}

