# ES6 模块化

## 语法

## 导出方式

``` javascript
//a.js
//导出方式1
export let firstName = "Michael";
export let lastName = "Jackson";
export let year = 1958;
//导出方式2
let oox = 10;
let name = "小明";
export { oox, name };
```

## 相应的接收方式
``` javascript
// index.js
import { year, firstName, lastName, oox, name } from "./a.js";
console.log(year, firstName, lastName, oox, name);
```

# 别名 as

``` javascript
//在导出端用法
let a = 10;
export {
    a as b
    a as c
}
 // 在导入端用法
 import {b,c} form './a.js'
 conosl.log(b,c)
```

# export default

```javascript
//规则只能导出一个
export default a
export default {} //对象
// export 能导出多个
相应接受规则 可以自定义变量
 import a form './a.js'
```
# import()  按需导入的时候用到,函数返回一个promise