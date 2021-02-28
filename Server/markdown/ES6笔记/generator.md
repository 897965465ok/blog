# generator

## 语法

```
 function* test(one) {
        let two = yield `我是第${one}步`
        let three = yield `我是第${two}步`
        let four = yield `我是第${three}步`
        return four
    }
    let obj = test(1);
    let a = obj.next() //第一任何参数都无效
    let b = obj.next(2)//传的是上一个yield的返回值 
    let c = obj.next(3)//传的是上一个yield的返回值
    let d = obj.next(4)//传的是上一个yield的返回值  
    console.log(a, b, c, d)
```

![](https://upload-images.jianshu.io/upload_images/7265016-865b688fd1eab83e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=ABCD%E7%9A%84%E5%80%BC)

# yield是什么

1. yield 只能在generator函数里面使用
2. 第一个yield只能通过定义函数传参
3. yield 的返回值是下一个yield的参数

# generator.next()

1. generator对象是函数默认返回的值
2. test.next()是分段的,使用test.next()可以调用一个yield
3. 第一个test.next()传的任何参数都是无效的
4. test.next()传的参数都是上一个yield的返回值
5. test.next()的返回值是一个对象 value 是数据 done 代表是否运行结束布尔值
6. generator函数里面的return 将由最后一个test.next()返回 在value里

![](https://upload-images.jianshu.io/upload_images/7265016-536acfa17f5e3f0d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=%E6%9C%80%E5%90%8E%E4%B8%80%E6%AC%A1test.next%28%29)

#其他用法

```
for (let item of  generator ) { //对象方式调用
    console.log(item)
}
这样用 let [a, b, c] = gen() //直接返回解构赋值
```
