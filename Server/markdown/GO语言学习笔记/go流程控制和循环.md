## 知识点

1. for ,if else , switch ,range

> > ## for ,if else, range

```golang
package main

import "fmt"

func fors() {
package main

import "fmt"

func fors() {
	//声明数组
	var array = [...]int{1, 2, 3, 4, 5}

	//方式一
	for i := 0; i < len(array); i++ {
		if i == 2 {
			continue //跳过本轮循环
		} else {
			fmt.Println(i)
		}
	}
	/*  结果
	0
	1
	3
	4
	*/

	//方式二
	// range返回两个变量一个是索引 一个是值
	// 多变量赋值可以用:=方式赋值因为GO可以判断类型
	for index, value := range array {
		fmt.Printf("index=%d \n,value=%d \n", index, value)
	}

	/*   结果
	index=0
	,value=1
	index=1
	,value=2
	index=2
	,value=3
	index=3
	,value=4
	index=4
	,value=5
	*/

	//如果只想获取index,或value,另外一个可以用下划线 _ 符号代替一切略过的变量
	for _, value := range array {
		fmt.Printf("value=%d \n", value)
	}

	/*  结果
	value=1
	value=2
	value=3
	value=4
	value=5
	*/
}

func main() {
	fors()
}
```

> ## switch

```golang
package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "text.go"
	//readFile返回两个参数  if 支持一个初始表达式 nil 等于null
	if constents, err := ioutil.ReadFile(filename); err != nil { //fuck 奇怪的语法  这样呢变量作用域在if里面了
		fmt.Println(err) //错误打印错误
	} else {
		fmt.Printf("%s\n", constents) //打印内容
	}
}

func switchs(a, b int, option string) (int, int) {
	//go语言里面自动 break ,  除非使用fallthrough
	var result int
	switch option { //判断option是什么然后自动对号入座
	case "+":
		result = a + b
		//使用fallthrough 将不会break
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator" + option) //终止代码
	}
	return result, a
}

func main() {
	result, _ := switchs(1, 2, "1")
	fmt.Println(result)
	readFile()
}
```
