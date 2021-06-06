1. map 声明的语法
2. map 声明的几种方式
3. map 增，删，查，改

## 语法

map[key]value

```golang
	m1 := map[string]string{ // :=创建
		"name": "小明",
		"age":  "18",
	}

	m2 := make(map[string]int) //通过make创建

	var m3 map[string]*int //通过var
```

## 增，删，查，改

```golang
	// 遍历
	for k, v := range m {
		fmt.Println(v, k)
	}

	// 访问的时候会返回布尔值
	if name, ok := m["noName"]; ok { //如果noname存在ok就为true
		fmt.Println(name, ok)
	}

	// 增
	m["sex"] = "girl"

	//删
	delete(m, "name") //删除name

	// 改
	m["sex"] = "人妖"
```

## 瞎写

```golang
package main

import "fmt"

func main() {
	fmt.Println(lengthOfNonRepeationgSubstr("abcabcbb"))
	fmt.Println(lengthOfNonRepeationgSubstr("bbbbb"))
	fmt.Println(lengthOfNonRepeationgSubstr("pwwkew"))

}

func lengthOfNonRepeationgSubstr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我不会告诉你,我口袋有两块糖" //UTF-8编码
	// fmt.Println(len(s)) //打印字节数
	fmt.Printf("%s\n", []byte(s))

	for _, b := range []byte(s) { //b rune  16进制编码
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune  unicode编码
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count", utf8.RuneCountInString(s)) //返回字符串长度
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //返回首字符和字符大小
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) { //4个字节
		fmt.Printf("%d %c", i, ch)
	}
}

//改成支持中文版 用rune就行了
//其他字符串操作
// Fields,Split,join
// Containes,Cndex
// Tolower,Toupper
// Trim ,TrimRight,TrimLeft
```
