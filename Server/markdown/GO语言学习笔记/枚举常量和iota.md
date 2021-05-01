# 枚举常量和iota

- 常量赋值后不能改变
- 枚举 列举  一周有几个星期天 1~5天这个就是枚举
- iota 配合枚举使用


 ## const 和  iota



```golang
	const fileName string = "abc.txt" //定义常量
    //定义后无法在次修改  	fileName = "123" 错误

const ( //当做枚举来用
cpp = iota //iota 表示自增值Z
java
ptyhon
glang
      )
fmt.Println(cpp, java, ptyhon) //0123

const (
		b = 1 << (10 * iota) //套用这个公式
		kb
		mb
		gb
		tB
		pb
)
	fmt.Println(b, kb, mb, gb, tB, pb)  //1 1024 1048576 1073741824 1099511627776 1125899906842624

	const (
		a string  = "1"
		c bool    = false
		d float64 = 4.12354
	)
	fmt.Println(a, d, c) // 1 4.12354 false
```
