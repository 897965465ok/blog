- 声明方式
- 数组给函数传递参数的方式
- 遍历


> > ## 声明方式



```
// 长度在前，类型在后
	//定义数组的几种方式
	var One [5]int                                   //声明len为5类型为int的的空数组
	var Tow = [2]int{1, 2}                           //声明并且赋值len为2数据为int(有值的数组)必须在后面加 = 号
	var Three = [...]string{"L", "o", "v", "e", "r"} //可以让系统计算有多少位
	var Four [2][2]*int                              //声明指针数组
	Five := [...]int{1, 2, 3, 4}                     //也可以这样声明和赋值但后面必须带花括号
	fmt.Println(One, Tow, Three, Four, Five)
```


> > ## 数组给函数传递参数的方式



```
func print(value [2]int, pointer *[2]int) { //一个是值接受(拷贝)，一个是指针(引用传递|地址传递)
	pointer[0] = 100                    //指针传递在修改后原数据会改变
	value[1] = 200                      //而值传递修改后原数组不会有任何改变
	for value, index := range pointer { //遍历方式
		fmt.Println(value, index)
	}
}

func main() {
  	var Tow = [2]int{1, 2}    
    print(Tow, &Tow)
	fmt.Println(Tow) //[100 2]
}
```


> > ## 遍历



```
	for value, index := range pointer { //遍历方式
		fmt.Println(value, index)
	}
```
