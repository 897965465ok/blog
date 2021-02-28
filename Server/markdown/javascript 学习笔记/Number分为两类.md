# Number分为两类

## (1) 第一类是4种原始值

1. null 返回0
2. undefined 返回NAN
3. 数字只是简单的传入返回
4. boolean true 是1 false 是0

## (2) 第二类是字符串和对象

### 字符串又分5种

1. 第一种整数在多字节的情况下[多字词首] 有(正或负符号)(零和多个零)(正或负后面加一个或多个零)
2. 第二种是浮点数规则和整形一样, 不过有一点和整形不一样, 浮点数可以解析一个小数点
3. 第三种情况是0X只能在词首第一位有效, (0x) 后面可以加整数字符串, 如果在加一位ox会被当做非数字处理 => 0xa0 xb NaN
4. 第四种情况如果字符串是空返回0
5. 第五种情况如果出现以上非上述情况返回NaN

### 对象

1.先调用valueOf() 在依照上面4种普通情况转换 如果转换结果为NAN,那就调用toString() 依照字符串规则转换

# parseInt("",10)

1. 在多字节的情况下[多字词首] 有(正或负符号)(零和多个零)(正或负后面加一个或多个零) 会被忽略掉
2. 和Number不一样的地方(空字符串返回NaN)(parseInt重左到右解析直到不能解析为止返回被解析的值)(可以传参)
3. Number如果字符串中有一个非数字直接返回NaN

### parsefloat()(只能解析一位小数点)(没有参数) 其他和parserint一样

# 各类型转换

**其他类型转数字**

```
Boolean  true =>1  false => 0
null =>0
undefined => NaN
String的话调用Number()函数
```

**其他类型转字符串**

```
Boolean  true => "true" false => "0"
Number   0=> "0"  "1234"  "" => ""
undefined =>"undefined"
null => "null"
```

**其他类型转换成布尔值**

```
Number  1=>true 0=>false NaN=>false
String  "非空算真"=>true  ""=>false 
undefined =>false
null =>false
对象转为布尔都为 true
```

**对象到字符串** 


```
如果对象有 toString() 方法，就调用 toString() 方法。如果该方法返回原始值，就将这个值转化为字符串。
如果对象没有 toString() 方法或者 该方法返回的不是原始值，就会调用该对象的 valueOf() 方法。如果存在就调用这个方法，如果返回值是原始值，就转化为字符串。
否则就报错
```

**对象到数字** 


```
对象转化为数字做了跟对象转化为字符串做了想同的事儿，不同的是后者是先调用 valueOf 方法，如果调用失败或者返回不是原始值，就调用 toString 方法。
补充。一些常用内置对象 toString 方法和 valueOf 的转换规则
```

## 逻辑AMD

```
 1. 只要“&&”前面是true，结果都将返“&&”后面的值;
 2. 只要“&&”前面是false，结果都将返“&&”前面的值;
 3. 如果触发短路 下面有这些值的 各自类型返回 null undefined NaN
```

## 逻辑OR

```
1. 只要“||”前面为true,都返回“||”前面的值。
2. 只要“||”前面为false,都返回“||”后面的值。
3. 如果触发短路 下面有这些值的 各自类型返回 null undefined NaN
```

```
一元运算符 正号负号 ++  -- * / % 如果碰到非数字类型就会调用Number()，
```

## 加法

```
加法碰到非数字类型会调用String()
```

## 减法

1. 如果有一个操作数是字符串、布尔值、null 或 undefined，则先在后台调用 Number()函数将

其转换为数值，然后再根据前面的规则执行减法计算。如果转换的结果是 NaN，则减法的结果

就是 NaN；
2. 如果有一个操作数是对象，则调用对象的 valueOf()方法以取得表示该对象的数值。如果得到

的值是 NaN，则减法的结果就是 NaN。如果对象没有 valueOf()方法，则调用其 toString()

方法并将得到的字符串转换为数值。

## 小于（<）、大于（>）、小于等于（<=）和大于等于（>=）

```
1、如果两个操作数都是数值，则执行数值比较。
2、如果两个操作数都是字符串，则比较两个字符串对应的字符编码值。
3、如果有一个是非数字则调用Number()
```

## (==) (!=)

```
转换操作数（通常称为强制转型），然后再比较它们的相等性。
在转换不同的数据类型时，相等和不相等操作符遵循下列基本规则：
1、如果有一个操作数是布尔值，则在比较相等性之前先将其转换为数值——false 转换为 0，而
true 转换为 1；
2、 如果一个操作数是字符串，另一个操作数是数值，在比较相等性之前先将字符串转换为数值；
3、 如果一个操作数是对象，另一个操作数不是，则调用对象的 valueOf()方法，用得到的基本类
型值按照前面的规则进行比较；
这两个操作符在进行比较时则要遵循下列规则。
null 和 undefined 是相等的。
要比较相等性之前，不能将 null 和 undefined 转换成其他任何值。
 如果有一个操作数是 NaN，则相等操作符返回 false，而不相等操作符返回 true。重要提示：
即使两个操作数都是 NaN，相等操作符也返回 false；因为按照规则，NaN 不等于 NaN。
如果两个操作数都是对象，则比较它们是不是同一个对象。如果两个操作数都指向同一个对象，
则相等操作符返回 true；否则，返回 false。
```

## Math常用方法

```
Math.abs:取绝对值
Math.ceil:向上取整
Math.floor:向下取整
Math.roud:四舍五入负数中5包含在向下
Math.random:获取0到1之间的随即小数
需求：获取3~15的随机整数
Math.round(Math.random()*12+3)
Max && min 获取最大最小
Math.PI:获取圆周率
Math.pow:获取一个值的多少次方
Math.sqrt:开平方
```

## String常用方法

```
charCodeAt&&fromCharCode()
charAt(索引):返回指定索引位置的字符,当索引不存在的时候返回""
charCodeAt(索引):返回索引字符的unicode编码值(对应ascll码表)
fromCharCode((unicode编码)):把编码转回字符串
substr&&substring&&slice [实现字符串截取]
substr(n,m):重索引n开始,截取M个字符串
substring(n,m):重索引n开始,截取到M不包含M
slice():和substring语法一样不过这个支持负数 原理用字符串总长度加上负数
toUpperCase && toLowerCase
toUpperCase:字符转大写
toLowerCase:字符转小写
indexOf&&lastIndexOf
indexOf:字符在字符串中的第一次出现的位置索引
lastIndexOf:获取字符串最后一次出现位置的索引
split&&join
split:按照谋一个字符把字符串拆分 列:split("|")
join:按照谋一个字符把字符串拼接
replace:替换字符执行一次只能替换一次列 replace('a','b') 返回改变后的字符串,原有字符串不改
trim&&trimRight&&trimLetft1清空空格2清空右边空格3清空左边空格
需求：queryURLParameter:获取url
for循环可以遍历私有属性 而for in循环可以遍历公有属性
```

## 数组常用方法

```
1.方法的形参
2.方法的返回值
3.原来的数组是否改变
push:数组后面增加加元素  参数(元素) 返回数组lenght  原有数组改变
unshift:向数组开头增加元素 其他push一样
pop:删除数组最后一项 参数无 返回被删的内容 原有数组改变
shift:删除数组第一项 其他和pop一样
splice(n,m):重索引n开始删除M个元素(M不写删除到末尾) 返回被删除的内容 不写索引(删除全部)
splice(0)删除全部 splice()一个不删返回数组
splice(n,m,x)在原有删除的基础上,用x代替删除的内容 返回被删除的内容
```

## 作用域

```
1. 分词/词法分析 这个过程会将由字符组成的字符串分解成,有意义的代码块，这些代码块被称为词法单元,无状态的叫分词 有状态的叫词法分析
2. 解析/语法解析  先把这些词法单元流（数组）转换成一颗名叫AST树
3. 代码生成将 AST 转换为可执行代码的过程称被称为代码生成
```

ES6

设计图

1.计算基准值 的方法 设计师给设计图  先用图宽比如 750/10=75px 那么这个75px 就是基准值了

2.基于基准值75px 转换rem应该怎么转换呢 如果你想设置一个宽为300px的div 那公式就是用 300px/基准值75px=4rem 那4rem就是300px了

let

没有变量提升

不可以重复声明

有块局作用域

不会给window增加属性

---

const 和上面一样

const一旦声明必须赋值 这个是常量

---

块局作用域

{}在行首会被当成代码块

如果你想当成对象可以这样（{}）解决

if(1){

function aa(){}

}

在if里面函数在预处理的时候只能是被声明

---

解构赋值：

1.左右两边解构必须一样

2.右边必须是个东西

3.声明必须和赋值一起

4.可以设置默认值只有后面的值是undifine的时候才会走默认值

5.省略赋值 列:let[,,m2]=[1,2,3] 只有M2是有值一一对应

6.不定参数赋值 列let[a,b,...c]=[1,2,3,4,5,6] 将后面的项放在一个数组中赋值给C

7.对象解析赋值 属性名匹配属性名 列： var{name:name1,age:age2}={name:"小江",age:18}  var{name,age}={name:'ok',age:18} 简写

先到对象里面的name属性取name属性的值取出来赋值给name1变量,

先到对象里面的age属性取age属性的值取出来赋值给age1变量

8.对象解析赋值默认值 列： let{ok:a=15,one:b=25}={ok:10,one:20} //a =10 b=20  简写let{ok=10,one=20}={ok:15,one:50}

9.解构赋值时，如果等号右边是数值和布尔值，则会先转为对象

10.函数也可以使用解构赋值

函数使用解构赋值规则 如果传值不走右边，不传的话走右边

---

es6字符串最新方法

includes：查询字符串是否有某个字符 参数(a,索引) 返回值Boolean

startsWith:判断字符串是否以某个字符串开头， 参数("字符"，索引) 返回值Boolean

endsWith:判断字符串是否以某个字符串结尾，其他和startsWith一样

repeat:将字符串复制N次 参数(n) ,返回值新的字符串

padStart: 将字符填充到字符串前面并限定长度 列：[//str.padStart](//str.padStart)(6,"lover")"love小明" 只保留6位

padEnd:将字符填充到字符串后面并限定长度其他和padStart一样

以后使用字符串可以这样用了!

let str=`哈哈哈`;

`你敢说你没笑？${str}`  //你敢说你没笑？哈哈哈

ES6数组方法

copyWithin

fill

fillter

find

findIndex

includes

every

some

reduceRight

keys

entries

---

node.js   遵守commonjs规范  一个js 就是一个模块

导入 require

导出 module.exports

---

es6 moduel

导出 exports

导入 import  form

语法 列:

// 写法一

export var m = 1;

// 写法二

var m = 1;

export {m};

// 写法三

var n = 1;

export {n as m};
