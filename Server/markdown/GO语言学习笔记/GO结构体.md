# GO 结构体

#### 1. 结构是由 struct 和 type 定义的

```golang
type Proson struct {  //定义结构体
	name string
	age  int
	sex  string
}
```

#### 2. 结构体的声明和定义

```golang
	//定义方法1
	var john Proson
	john.name = "john"
	john.age = 28
	john.sex = "男"
	fmt.Println(john.name)

	//定义方法2
	ben := Proson{name: "ben", age: 18, sex: "男"}
	fmt.Println(ben)

	//定义方法3
	sofia := Proson{"sofia", 17, "女孩"}
	fmt.Println(sofia)

	// 定义方法4
	var jiang = new(Proson) // 等于 jiang := &Proson
	jiang.age = 18
	jiang.name = "江"
	jiang.sex = "男"
	fmt.Println(jiang)

	//匿名结构体 定义方式5
	unknown := struct {
		name string
		age  int
	}{
		name: "无名",
		age:  28,
	}
	fmt.Print(unknown)

	// 匿名字段 定义方式6
	type Unkey struct {
		string
		int
	}
	unkey := Unkey{"匿名字段", 01}
	fmt.Print(unkey)

	//切片结构体 定义方式6
	nodes := []Proson{{"One",nil,"男"}}
	fmt.Println(nodes)
```

#### 3. 为结构体添加方法

```golang
//给结构体声明方法
func (node *Proson) setvalue(age int) { //node 这个参数传的是地址
	node.age = age
}
john.setValue(123) //调用方法
```
