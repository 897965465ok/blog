# node.js调用golang生成的动态库

## golang

1. 设置编译环境为64位 GOARCH=amd64 ，如果想设置为32位 GOARCH=386
2. 需要生成动态库的文件必须要符合以下条件

   1. import "C"
   2. main函数
   3. 导出的函数首字母必须大写
   4. 函数前必须要放//export (函数名)
3. 生成 dll 或者 os 文件

## 生成动态库

```
awesome.dll 是windows 动态库
go build -o awesome.dll -buildmode=c-shared awesome.go
awesome.so 是mac 动态库
go build -o awesome.so -buildmode=c-shared awesome.go
```

### 列子:

```
package main

import "C"

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int {
	return a + b
}

//export Cosine
func Cosine(x float64) float64 {
	return math.Cos(x)
}

//export Sort
func Sort(vals []int) {
	sort.Ints(vals)
}

//export SortPtr
func SortPtr(vals *[]int) {
	Sort(*vals)
}

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

//export LogPtr
func LogPtr(msg *string) int {
	return Log(*msg)
}
func main() {}
```

## node.js

1. 切换node.js到8.17版本,ps 可以用nvm版本管理工具
2. npm install --global --production windows-build-tools
3. npm install -g node-gyp
4. npm install ffi ref ref-struct ref-array

### 列子:

```
var ref = require("ref");
var ffi = require("ffi");
var Struct = require("ref-struct")
var ArrayType = require("ref-array")
var path = require('path')
var LongArray = ArrayType(ref.types.longlong);
var GoSlice = Struct({
  data: LongArray,
  len:  "longlong",
  cap: "longlong"
});
var GoString = Struct({
  p: "string",
  n: "longlong"
});

var awesome = ffi.Library(path.resolve(__dirname,'./awesome.dll'), {
  Add: ["longlong", ["longlong", "longlong"]],
  Cosine: ["double", ["double"]],
  Sort: ["void", [GoSlice]],
  Log: ["longlong", [GoString]]
});
console.log("awesome.Add(12, 99) = ", awesome.Add(12, 99));
console.log("awesome.Cosine(1) = ", awesome.Cosine(1));
nums = LongArray([12,54,0,423,9]);
var slice = new GoSlice();
slice["data"] = nums;
slice["len"] = 5;
slice["cap"] = 5;
awesome.Sort(slice);
str = new GoString();
str["p"] = "Hello Node!";
str["n"] = 11;
awesome.Log(str);
//  awesome.dll 是windows   awesome.so 是mac
// go build -o awesome.so -buildmode=c-shared awesome.go
```
