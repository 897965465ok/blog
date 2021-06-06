# go语言初体验

- 引入包
- 变量的几种声明方法
- 变量的类型

### 1. 引入包:

```golang
- 每一个入口程序必须有一个主函数:main函数

- 每一个go文件必须声明包名 声明包名字语法:  package 名字

- 导入其他包:语法 import  包名
```

#### 列子:

```golang
package main //声明包名

import "fmt" //导入其他包

func main (){ //主函数

 fmt.Println("Hello, World!")

}
```

### 2. 类型:

#### 布尔类型：


- 布尔型的值只可以是常量 true 或者 false 例子：var b bool = true

#### 字符串类型:


- string  例子：var a string = "我是字符串类型"

#### 派生类型:


- 指针类型 (Pointer)  例子： var a  *int = &b
- 数组类型 Array    例子:  var a [10]int{1,2,3,4,5}
- 结构化类型(struct)
- Channel 类型
- 函数类型 (function)
- 切片类型 (Slice)
- 接口类型（interface）
- Map 类型

#### 数字类型:


- uint8 无符号 8 位整型 (0 到 255)
- uint16 无符号 16 位整型 (0 到 65535)
- uint32 无符号 32 位整型 (0 到 4294967295)
- uint64 无符号 64 位整型 (0 到 18446744073709551615)
- int8 有符号 8 位整型 (-128 到 127)
- int16 有符号 16 位整型 (-32768 到 32767)
- int32 有符号 32 位整型 (-2147483648 到 2147483
- int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)


 #### 浮点型




- float32 IEEE-754 32位浮点型数
- float64 IEEE-754 64位浮点型数
- complex64 32 位实数和虚数
- complex128 64 位实数和虚数

#### 其他数字类型


- byte 类似 uint8
- rune 类似 int32
- uint 32 或 64 位
- int 与 uint 一样大小
- uintptr 无符号整型，用于存放一个指针

### 3. 定义变量的几种方式

```golang

//第一种，指定变量类型，如果没有初始化，则变量默认为零值。

// var identifier type

var name string
var age int

// 第二种，根据值自行判定变量类型。

var Number =  1;
var string = "1";

// 第三种 只能在函数内用，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式： 

var intVal int 
intVal :=1 // 这时候会产生编译错误
intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
```

#### 多变量声明

```golang
//类型相同多个变量, 非全局变量
var a,b,c = 1,'1',true   // 和 python 很像,不需要显示声明类型，自动推断

q, w, e, r := 1, 2, 3, 4 //也可以这样用 在函数外面不能这样用

   var ( //也可以这样写  等于 a=1,b=2,c=3
   	 q = 1
   	 w= 2
   	 e= 3
   )
```

#### 瞎写

```golang
package main //主包

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	euler()
	triangle()
}

func euler() {
	c := 3 + 4i
	fmt.Println("其实我不知道这个是什么意思")
	fmt.Println(cmplx.Abs(c)) //取绝对值
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //只有强转
	fmt.Println(c)
}
```
