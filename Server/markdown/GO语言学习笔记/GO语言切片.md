# GO 语言切片

1. 切片的声明语法
2. 切片跟数组的关联

## 声明语法

```go
	//切片和数组最好的区分是切片在(定义的时候不需要指定长度)
	var One []int
	var Two = []string{"L", "o", "v", "e", "r"}
	Tree := []int{1, 2, 3}
	Four := make([]int, 5)                       // 声明和赋值 一个len为5的空切片
	Five := make([]int, 5, 10)                   // 声明和赋值 一个len为5,cap[底层映射长度]为10的空切片
	Array := [...]int{1, 2, 3, 4, 5}             //创建一个数组
	Six := Array[:]                              //根据数组创建slice
	fmt.Println(One, Two, Tree, Four, Five, Six) //[] [L o v e r] [1 2 3] [0 0 0 0 0] [0 0 0 0 0] [1 2 3 4 5]
```

## 切片跟数组的关联

```go
    var Array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	One := Array[:5] //将数组转成切片 从0~5 5为数
	Tow := One[1:6]  //目前这个切片只有5个，如果我们这样拿，会发生什么,让我们来打印一下
	fmt.Println(Tow) //[2 3 4 5 6]
	//因为映射的关系，如果往后越界了,不超过自己的cap()依然管用
	fmt.Println("One= len()", len(One), "One= cap()", cap(One)) //打印出切片长度 和切片底层长度 One= len() 5 One= cap() 10

	/* 如果超过cap长度, 如果往前越界， 系统将报错
	Tow = Tow[0:]
	fmt.Println(Tow) //eroor
	Tow = One[:11]
	fmt.Println(Tow) // eroor
	*/
```

```go
	var Array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	One := Array[:5] //将数组转成切片 从0~5 5为数
	One[0] = 100       //有趣的是修改这个切片 上面的数组也会受影响
	fmt.Println(Array) //[100 2 3 4 5 6 7 8 9 10]
	Array[0] = 200     //当然你修改这个数组的时候 切片也会受到影响
	fmt.Println(One)   // [200 2 3 4 5]
```

###没有修改前
![](https://upload-images.jianshu.io/upload_images/7265016-8b41bd6f773095b2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#align=left&display=inline&height=141&margin=%5Bobject%20Object%5D&originHeight=141&originWidth=641&status=done&style=none&width=641)

### 修改后

![](https://upload-images.jianshu.io/upload_images/7265016-29044b01f4cf7146.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#align=left&display=inline&height=71&margin=%5Bobject%20Object%5D&originHeight=71&originWidth=659&status=done&style=none&width=659)

#####由此得知数组和切片存在映射关系

# 分分合合 有时切片也会抛弃映射的数组

```go
	Array := [...]int{1, 2, 3}
	One := Array[:]
	Tow := One[1:]
	fmt.Println("这是从数组中创建的切片", Tow) //这是从数组中创建的切片 [2 3]
	Three := append(Tow, 4)         //如果append后超过了cap，那么系统生成新的切片,新的切片将拥有新底层数组，和旧底层数组没有任何关系了分手了~！
	Three[2] = 100
	fmt.Println(Three, Array) //[2 3 100] [1 2 3] 没有任何关联
```

![](https://upload-images.jianshu.io/upload_images/7265016-34b851512911c05c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#align=left&display=inline&height=252&margin=%5Bobject%20Object%5D&originHeight=252&originWidth=362&status=done&style=none&width=362)

# 一把掌甩到 Array 脸说,你的房子太小,系统已经给我更大的房子我不爱你了,要离开你
