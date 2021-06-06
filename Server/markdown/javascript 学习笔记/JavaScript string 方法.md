# JavaScript string 方法

```js
var tostring = "你为什么要说这种话，蓝瘦香菇"
console.log(tostring.charAt(0))//=> 返回下标的值  (你)
console.log(tostring.charCodeAt(3))//=> 返回下标值 ascll码表(比如大写的A就返回65) 
console.log(String.fromCharCode(65)) // =>把ascll码表转成字符串(和上面相反65 转A)
console.log(tostring.substr(0,4)) //=> 重N开始 截取M个字符串(重下标0开始,截取4个字符串)
console.log(tostring.substring(3,4)) //=>重N开始截取到M(不包括M) substring(n,m)
console.log(tostring.slice(0,20))//=> 重N开始截取到M(不包括M) slice(n,m)这个支持负数
console.log(tostring.toUpperCase())//=>tostring.toUpperCase()将字符串转成大写
console.log(tostring.toLowerCase())//=>tostring.toLowerCase()将字符串转成小写
console.log(tostring.indexOf("A"))//=> tostring.indexOf("A")返回第一次出现的下标
console.log(tostring.lastIndexOf("A"))//=> tostring.lastIndexOf("A")返回最后出现的下标
console.log(tostring.split("A"))//=>以字符为间隔将字符串拆成数组
console.log(tostring.join("A"))//=>将字符A做间隔合成字符串
console.log(tostring.replace("A","B"))//=>tostring.replace()替换A换成B
console.log(tostring.trimLeft("A","B"))//=>去除左空格
console.log(tostring.trimRight("A","B"))//=>去除左右空格
tostring.trim()//=> 去除全部空格
```
