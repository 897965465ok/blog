# js-预解释

###javascript 的预解释也叫变量提升是我见过最无节操的

但是掌握预解释并不难

1. 
所有带var 的先声明单不赋值

2. 
function 是直接声明和赋值


3.所有声明过的浏览器绝不声明第二次

举个栗子

```
第一步开始声明所有带var的变量,注意是声明声明声明，
第二步开始声明+赋值所有带function的变量，注意是声明+赋值声明+赋值声明+赋值，
然后开始重上往下执行，
   开始预解释
   第一步 和第二部同时进行  
   先预解释父函数
   就一个function a (){}
   接下来到子函数
   var a;  var b;  function a(){} 覆盖了var a； function b(){} 预解释完成!
 然后开始运行  
   (function(
		var a=10; 
		function a() {

		};
           alert(a);//10
           function b(){

           };
           b=6;
           alert(b);//6
           var c=d=b;
       })();

       alert(d);//6
       alert(c); //报错
       a=5;  
       function a(){

        };
      console.log(a)
```

4. 如有形参 形参名字和函数里内属性名字相同 ,将忽略 变量声明  但是函数还是会覆盖的

```
var name ="小明";
function box(name){
var name; //将不进行预处理
alert(name);//小明
name = 10;
alert(name) //10
}
box(name);
```
