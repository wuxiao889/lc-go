package largestNumber

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numsString := make([]string, len(nums))
	for i := range nums {
		numsString[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(numsString, func(i, j int) bool {
		//递减排序
		return !less(numsString[i], numsString[j])
	})
	//特殊情况，最大值是0
	if numsString[0] == "0" {
		return "0"
	}
	return strings.Join(numsString, "")
}

//比较规则
//字符串的比较，go语言不用strings.Compare
//直接用符号比较

//按ASCII码逐位比较，如果len(num1)<len(num2)，并且num1==num2[:len(num1)],那么num2大

//但是这题
//比如说正常规则下3 < 30，但是330>303,所以3>30
//正常, 3 < 34 ， 334 < 343 ， 3 < 34
//所以 num1 < num2 等价于 num1 + num2 < num2 + num1
func less(num1, num2 string) bool {
	return num1+num2 < num2+num1
}
