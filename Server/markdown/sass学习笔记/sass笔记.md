# sass笔记

```
***sass 支持变量***
$nav-color:red;
nav{
    color: $nav-color;
}
//
nav {
  color: red;
}
```

```
sass支持嵌套
$nav-color:red;
nav{
    color: $nav-color;
    ul{
        
        li{
          display:inline-block;
          line-height:50px;    
          a{
              display: inline-block;
              line-height:50px;  
              :hover{
                color: aqua; 
              }  
          }  
        }
    }
}
//
nav {
  color: red;
}
nav ul li {
  display: inline-block;
  line-height: 50px;
}
nav ul li a {
  display: inline-block;
  line-height: 50px;
}
nav ul li a :hover {
  color: aqua;
}
```

```
sass支持函数
@mixin定义函数 @include使用函数  % 站位符继承用的 @extend继承 
$nav-color:red;
@mixin alert($set-color, $set-width) {
    li {
        width: $set-width;
        a {
            color: darken($set-color, 20%);
        }
    }
}

ul {
    @include alert(red, 100px);
}

.ui-set {
    @include alert($set-color: blue, $set-width: 100%)
}
//
ul li {
  width: 100px;
}
ul li a {
  color: #990000;
}

.ui-set li {
  width: 100%;
}
.ui-set li a {
  color: #000099;
}
```
