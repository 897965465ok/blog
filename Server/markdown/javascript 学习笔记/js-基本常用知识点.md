# js-基本常用知识点

```
F12控制台常用
console.dir()//=>比console.log()详细一点,
console.table()//=>吧json数据输出成表格，
怎么多数据js如何判断
typeof:检测数据类型的运算符，
instanceof：检测某个实例是否属于这个类，
constructor:返回当前实例的构造器,
Object.prototype.toString.call:获取当前实例所属类信息，
isNaN("12")当我们用isNaN检测值时浏览器先会把值转换成number类型,再去检测
把字符串转换成number类型的方法
Number("12px")=>NaN
parseInt("12px")=>12
parseInt("12px13")=>12提取规则:从左到右
依次查找有效数字字符,直到遇见非有效数字字符为止（不管后面是否还有，都不找了），把找到的转换为数字
parseInt("px12")=>12
```

---


获取数组最大最小值:

方法1：//利用 max ,min 会把数组当字符串传过去

var ary=[23 ,34 , 24 , 12 ,35 ,36 ,14 ,25];

var max = Math.max.apply(null , ary);

var min = Math.min.apply(null , ary);

console.log(min,max)

方法2;

//先排序取第一个是最小值

//取最后一个是最大的

ary.sort(function(a,b){

return a-b;

})

var min =ary[0];

var max = ary[ary.length-1];

console.log(min,max)

方法三：

var max = eval("Math.min("+ ary.toString()+")");

var max = eval("Math.max("+ ary.toString()+")");

console.log(max)


判断类型的方法

var o = {};

var a = [];

o instanceof Array // false

a instanceof Array // true

typeof null // "object"

typeof window // "object"

typeof {} // "object"

typeof [] // "object"

undefined == null

// true


Array.push() //=>往数组后面加返回值是新增后数组的长度

Array.unshift() //=>往数组开头加返回值是新增后数组的长度

Array.pop()//=>删除最后一个返回被删除的那一项内容

Array.shift() //=>删除数组第一项返回被删除的那一项

delete Array[0]; //=>删除指定索引对象

Array.splice(n,m) //=>重索引开始N开始删除M个返回的是被删除的内容

Array.splice(0) //=>清空数组

Array.splice()//=>一个都不删除

Array.splice(n,m,x)//=>重N开始用X代替M =>Array.splice(0,M,100)下标0开始删除0个在M之前插入100

Array.slice(n,m)//=>重索引N开始找到索引为M处不包含M返回新数组

Array.slice(0)//=>数组克隆

Array.concat(n,m)//=>返回新数组

Array.concat()//=>实现数组克隆

Array.join("+")//=>指定分隔符+号

eval(Array.join("+"))//=>把字符串变成字符串表达执行和求和

Array.reverse()//=>返回排序后的数组原有的改变

// Array.sort(function(a,b){//=>数组排序

//  return a-b;//升s

//  return b-a;//降

// })


indexOf/iastIndexOf:获取当前项在数组在中第一次

或者最后一次出现位置的索引

在数组中ie6-ie8不兼容

在字符串中这两个方法兼容

如果数组中没有这一项返回-1我们根据这一点可以见证数组中是否

包含这一项

if(Array.indexOf(12)>-1){

//->是否包含

}

myindexOF(value){

var result = -1;

for (var i = 0; i < this.length; i++) {

if(this[i]==value){

resilt = i;

break

}

}

return resilt;

}


便利方法ie6-ie8不兼容

Array.forEach(function(value,index){

//数组中有多少项,当前回调函数会执行多少次

//每一次穿丢进来的value就是当前遍历数组这一项的值

//index 就是遍历这一项的索引

})


// map 遍历数组中的每一项在forEach的的基础上，

// 可以修改每一项的值

Array.map(function(value ,index){

//数组中有多少项,当前回调函数会执行多少次

//每一次穿丢进来的value就是当前遍历数组这一项的值

//index 就是遍历这一项的索引

//return xxx ;  return 后面返回的结果就是当前遍历的这一项修改为XXX

})


数组去重//

> 方案1：

遍历数组中的每一项，拿每一项和它后面的项依次比较

如果相同,则把相同的一项在原来数组中删除

for (var i = 0; i < Things.length-1; i++) {

var cur= Things[i];

for (var j = i+1; i < Things.length; j++) {	

cur===Things[j]?Things.splice(j,1):j++;

}

}

// >方案2：利用linuxOf来验证当前数组是否包含某一项，

包含把当前项删除掉(不兼容ie6-ie8)

for (var i = 0; i < Things.length; i++) {

var cur =Things[i];

var curNextAry = Things.slice(i+1);

if(curNextAry.indexOf(cur)>-1){

Things.splice(i,1);

i--;

}

}

方案3

遍历数组中的每一项，把每一项作为新对象的属性名和属性值存起来

例如：当前项1，对象中存储的{1:1}


在每一次向对象中存储之前，首先看一下原有对象中是否包含了这个属性

(typeof obj[xx]==="undefined",说明当前对象中没有这个XXX属性)，如果已经存在了这个属性

说明数组中的当前项是重复的（1-原有数组中删除这一项，2-不在向对象中储存这个结果）

，如果不存在（把当前对项作为对象的属性名和属性值存储进去即可）

var ary = [1,2,3,54,6,4,56,14,2,12,1,2,12,12,1,21,3,123,1];

var obj ={};

for (var i = 0; i < ary.length; i++) {

var cur=ary[i];

```
   if (typeof boj[cur]!=="undefined") {
   	ary[i]=ary[ary.length-1];
   	ary.length--;
   	i--;
   	continue;
   }
```

obj[cur]=cur;

}


F1 instanceof Box 检查一个实例属于那个类

例子

var a[];

console.log(a instanceof Array)

in 检查某个属性是否存在对象中不管共有私有

f1.hasOwnProperty("getx")查某个私有属性是否存在对象中

function haspubproperty(ogj, attr){

return (attr in obj) && !obj.hasOwnProperty

}


for(var key in obj){

//for in 循环在遍历的时候会默认的话吧自己私有的

//和在它所属类的原型上扩展的属性和方法都可以遍历到

//但是一般情况下，我们遍历一个对象只需要遍历私有的即可

//我们可以使用以下方法

//

if (obj.propertyisEnumerable(key)) {

console.log(obj)

}

}


propertyisEnumerable()//判断是是否私有属性

for(var key in obj){

//判断自己私有的

if (obj.hasOwnProperty(key)) {

console.log(obj)

}

}

var opop =Object.creat(boj)//创建一个对象把obj做原型

Array.filter()

Array.find()

Array.reduce()

Array.every()

//随机算法

function shuffle(arr) {

let i = arr.length;

while (i) {

let j = Math.floor(Math.random() * i--);

[arr[j], arr[i]] = [arr[i], arr[j]];

}

}

