# TS 高级类型

//高级类型

```
const mergeFunc = <T,U>(message:T , message2:U): T&U =>{
let result =   <T & U> {}
result = Object.assign(message,message2)
return result;
}
mergeFunc(1,2)

```
