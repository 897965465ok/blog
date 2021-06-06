# golang interface(接口) 笔记

### 概念

```golang
在Go中，接口是一组方法签名。
接口定义行为 ，具体的实现行为取决于对象
接口就像是一个约束 满足约束的 就属于此类接口
```

### 知识点

1. 规定与实现
2. 嵌套
3. 断言

## 规定与实现

```golang
package main

import "fmt"

// 定义吃的行为
type A interface {
	eat()
}

// 定义喝的行为
type B interface {
	drink()
}

type C struct{}

// C实现吃的行为
func (this C) eat() {
	fmt.Println("吃")
}

// C实现吃的行为
func (this C) drink() {
	fmt.Println("喝")
}

type D struct{}

// D实现喝的行为
func (this D) drink() {
	fmt.Println("喝")
}

//  接口校验 A
func verify__A(a A) {
	a.eat()
}

//  接口校验 B
func verify__B(a B) {
	a.drink()
}

func main() {
	var c = new(C)
	var d = C{}
	// C满足了A和B接口
	verify__A(c) // 吃
	verify__B(c) // 喝
	// D只满足了B接口，所以接口A会报错
	// verify__A(d)  cannot use d (variable of type D) as A value in argument to verify__A: missing method eat
	verify__B(d) // 喝

}
```

## 嵌套

```golang
package main

import "fmt"

// 定义吃的行为
type A interface {
	eat()
}

// 定义喝的行为
type B interface {
	A // 接口嵌套, 可以理解成B继承A的定义
	drink()
}

type C struct{}

// C实现吃的行为
func (this C) eat() {
	fmt.Println("吃")
}

// C实现吃的行为
func (this C) drink() {
	fmt.Println("喝")
}

type D struct{}

// C实现吃的行为
func (this D) eat() {
	fmt.Println("吃")
}

//  接口校验 A
func verify__A(a A) {
	a.eat()
}

//  接口校验 B
func verify__B(a B) {
	a.drink()
	a.eat()
}

func main() {
	var d D = D{}
	var c C = C{}

	verify__A(d)
	// 因为C结构体满足了B接口的所有方法
	verify__A(c)
	verify__B(c)

}
```

### 空接口断言

```golang
	/ 定义一个空接口
	var ovid interface{}
	var four int = 5
	//  将被断言的变量赋值给ovid
	ovid = four
	//  现在可以用ovid.(type)开始断言了
	value, err := ovid.(float32)
	// 如果正确返回 值和ture否者返回0和false
	fmt.Println(value, err)
```
