package core
/*
排列
*/

func Permute(trainsNums []string) [][]string {
	COUNT := len(trainsNums)
	if COUNT == 0{
		panic("不合法的参数")
	}
	//如果只有一个数，则直接返回
	if COUNT == 1 {
		return [][]string {[]string{trainsNums[0]}}
	}

	//递归插入排列的的所有位置
	return insert(Permute(trainsNums[:COUNT-1]), trainsNums[COUNT-1])
}

func insert(res [][]string, insertNum string) [][]string {
	//保存结果的slice
	result := make([][]string, len(res)*(len(res[0])+1))

	index := 0
	for _, v := range res {
		for i := 0; i < len(v); i++ {
			//在v的每一个元素前面插入
			result[index] = append(result[index], v[:i]...)
			result[index] = append(result[index] , insertNum)
			result[index] = append(result[index] , v[i:]...)
			index++
		}

		//在v最后面插入
		result[index] = append(result[index], v...)
		result[index] = append(result[index] , insertNum)
		index++
	}

	return result
}

var OPERATORS []string = []string{"+", "-", "*", "/"}

/*
操作符的组合
*/
func PermuteOperator(m int) [][]string{
	if m <= 0{
		panic("参数错误")
	}
     if m == 1{
	     return [][]string {
		     []string{OPERATORS[0]},
		     []string{OPERATORS[1]},
		     []string{OPERATORS[2]},
		     []string{OPERATORS[3]},
	     }
     }else{
	return appendOnce(PermuteOperator(m-1))
     }
}

func appendOnce(old [][]string)[][]string{
	results := make([][]string, 4*len(old))
	for i, v := range old{
		for ii, p := range OPERATORS{
			results[4*i+ii] = append(results[4*i+ii], v...)
			results[4*i+ii] = append(results[4*i+ii], p)
		}
	}
	return results
}