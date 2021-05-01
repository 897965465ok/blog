# Symbol
```
const s = Symbol('www.baidu.com')

 console.log(s.toString())

 let obj = {

     [s]:'list'

 }

 Object.getOwnPropertySymbols(obj)遍历symbols

 Reflect.ownKeys(obj)遍历symbols

const s8 = Symbol.for('li')会在全局查找有没有一样的

 const s9 = Symbol.for('li')

 const s10 = Symbol.for('ok')

console.log(s8s9)  true

 console.log(s8s10)  false

 console.log(Symbol.keyFor(s8)) 返回Symbol的标识 li

const _func2 = () =>{} 模拟私有方法

 class Iphone {

      fun1 (){

          _func2.call(this.fun1)

      }

 }
```