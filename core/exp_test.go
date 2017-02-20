package core
import (
	"testing"
	"fmt"
)

type TestExpression struct {
	Exp *Expression
	//表达式
	ExpStr string
	//结果
	Result int
}

func ExpressionFactory(optNumbers []int, operators []string, parenthesis [][2]int) *Expression{
	obj, err := NewExpression(optNumbers, operators, parenthesis)
	if err != nil{
		fmt.Println(err)
		panic("ExpressionFactory Error!!!!!!!!!")
	}else{
		return obj
	}
}

var addTests = []TestExpression{
	//1+2+3+4=10
	TestExpression{
		ExpressionFactory([]int{1,2,3,4}, []string{"+", "+", "+"}, [][2]int{[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}),
		"1+2+3+4",
		10,
	},
	//1+3+2+4=10
	TestExpression{
		ExpressionFactory([]int{1,3,2,4}, []string{"+", "+", "+"}, [][2]int{[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}),
		"1+3+2+4",
		10,
	},
	//(1+2)*(3+4)=21
	TestExpression{
		ExpressionFactory([]int{1,2,3,4}, []string{"+", "*", "+"}, [][2]int{[2]int{1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, 1}}),
		"(1+2)*(3+4)",
		21,
	},
	//1+2*3+4=11
	TestExpression{
		ExpressionFactory([]int{1,2,3,4}, []string{"+", "*", "+"}, [][2]int{[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}),
		"1+2*3+4",
		11,
	},
	//4-2+18/9*11+9/3
	TestExpression{
		ExpressionFactory([]int{4,2,18,9,11,9,3}, []string{"-", "+", "/", "*", "+", "/"}, [][2]int{[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}),
		"4-2+18/9*11+9/3",
		27,
	},
	//4-(2+18/9*11)+9/3
	TestExpression{
		ExpressionFactory([]int{4,2,18,9,11,9,3}, []string{"-", "+", "/", "*", "+", "/"}, [][2]int{[2]int{0, 0}, [2]int{1, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 1}, [2]int{0, 0}, [2]int{0, 0}}),
		"4-(2+18/9*11)+9/3",
		-17,
	},
	//2*6+6*2=24
	TestExpression{
		ExpressionFactory([]int{2,6,6,2}, []string{"*", "+", "*"}, [][2]int{[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}),
		"2*6+2*6",
		24,
	},
}

func TestExpString(t *testing.T) {
	for _,v := range addTests {
		if v.Exp.String() != v.ExpStr {
			t.Error("TestExpString Error!!!!")
			t.Error(v.Exp.String())
			t.Fail()
		}
	}
}

func TestExpValue(t *testing.T) {
	for _,v := range addTests {
		if expv, err := v.Exp.Value(); expv != v.Result {
			v.Exp.Display()
			t.Error("TestExpValue Error!!!!", err)
			t.Fail()
		}
	}
}

func TestExpDisplay(t *testing.T) {
	for _,v := range addTests {
		v.Exp.Display()
	}
}

func TestExpCalc(t *testing.T) {
	opts := []string {"+", "-", "*", "/"}
	testData := [][4]int{
		[4]int{ 12, 2, 0, 14},
		[4]int{ 12, 2, 1, 10},
		[4]int{ 12, 2, 2, 24},
		[4]int{ 12, 2, 3, 6},
	}
	for _,v := range addTests {
		for _, d := range testData{
			if v.Exp.Calc(d[0], d[1], opts[d[2]]) != d[3]{
				t.Error("TestExpCalc Error!!!!")
				t.Fail()
			}
		}
		break
	}
}

func TestEqual(t *testing.T) {
	if len(addTests) >= 2{
		if !addTests[0].Exp.Equal(addTests[1].Exp){
			t.Error("TestEqual Error!!!!")
			t.Fail()
		}
	}
}