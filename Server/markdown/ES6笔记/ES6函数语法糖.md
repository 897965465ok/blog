# ES6函数语法糖

## 箭头函数


1. 如果是匿名函数可以省略function



```
() => {

}
```


2. 如果是只有一个参数可以省略括号



```
message => {

}
```


3. 如果函数体只有一个表达式可以省略大括号 ,如果没有大括号将自动return



```
message =>  message * 3 
message = > ({  //注意这个是代码块返回的是一个对象
    age : 18,
    name : '小明'
})
```


4. 箭头函数中的this将指向他外层的function的this



```
   setInterval//在默认情况下指向windows
   let test = {
        set() {
            setInterval(()=>{
                console.log(this)
            }, 5000)
        }
    }

    test.set()// test
```
