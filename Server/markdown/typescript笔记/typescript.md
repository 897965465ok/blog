# typescript 笔记

## 基础类型：
``` typescript 
    Boolean: 布尔类型;   let isDone: boolean = false;
    Number: 数值类型;    let decLiteral: number = 6;
    String : 字符串类型; let name: string = "bob";
    Array: 数组类型;     let list: number[] = [1, 2, 3];    let list: Array<number> = [1, 2, 3];
    Tuple: 元祖;        let array: [string, number] = ["我",1]
    enum：枚举类型;      enum Color {Red, Green, Blue} let c: Color = Color.Green;
    any : 任意类型;      
    void：空类型函数没返回的时用;   let unusable: void = undefined|null;    
    Undefined|Null: 可以把 null和undefined赋值任意类型的变量;
    Never：永远不存在的类型 目前不了解 可能抛出错误是会用吧;
   
```
## 接口约束
``` typescript
    interface SquareConfig {

    color?: string; // ？表示可以不传

    readonly y: number; // 只读属性

    [propName: string]: any; // 表示这个对象可以动态添加属性

    funName:(source: string, subString: string): boolean; //函数约束

    numberKey:[index: number]: string; // 可索引约束 但是数字索引的返回值必须是字符串索引返回值类型的子类型
    stringKey:[string: number]: string; //  第二种
}

// 类接口
interface ClockInterface {
    currentTime: Date;
}
// implements关键字
class Clock implements ClockInterface {
    currentTime: Date;
    constructor(h: number, m: number) { }
}


```
# 需要记忆的东西
``` typescript
    ReadonlyArray<T>:只读数组类型;  
    protected:   受保护的 当前类 和子类可以访问
    private:     私有的
    public:      公有
    let a: number[] = [1, 2, 3, 4];
    let ro: ReadonlyArray<number> = a;
    ro[0] = 12; // error!
    ro.push(5); // error!
    ro.length = 100; // error!
    a = ro; // error!
    a = ro as number[]; // yes
```



# 函数
``` typescript

// 重载
function getInfo(name:String):String;

function getInfo(age:Number):number;

function getInfo(str:any ):any{ 

    if(typeof str=='string'){

        return str

    }else{

        return  str

    }

}

```



// console.log(getInfo('18'))

// function getInfo(name:string):string;

// function getInfo(name:string,age:number):string;

// function getInfo(name:any,age?:any):any{

//     if(age){

//         return '我叫:'+name+'我的年龄是'+age

//     }else{

//         return '我叫:'+name;

//     }

// }

// // alert(getInfo('张三'))

// alert(getInfo('张三',189))

// class Box {

//     name: string;

//     age: number;

//     constructor(name: string, age: number) {

//         this.name = name;

//         this.age = age;

//     }

//     print(): void {

//         console.log(this.name, this.age)

//     }

// }

// class Iphon extends Box {

//     width: string;

//     hight: number;

//     constructor(width: string, hight: number, name: string, age: number) {

//         super(name, age)

//         this.width = width;

//         this.hight = hight;

//     }

//     get() {

//         //    console.log(this.name,this.age,this.width,this.hight)

//         super.print()

//     }

// }

// let iphon7 = new Iphon('375px', 350, 'iohon7', 3);

// iphon7.get()

// iphon7.print()

// class Box {

//    name: String

//    age: Number

//     constructor(name: String, age: Number) {

//         this.name = name;

//         this.age = age;

//     }

//     protected printf(){

//         console.log(this.name , this.age)

//     }

//     arr (){

//         this.printf()

//     }

// }

// let box1 = new Box('小明',18)

// box1.arr()

//多态 父类定义一个方法不去实现 让继承它的子类去实现 每个一个子类有所不同

// class Animal {

//     name:string;

//     constructor(name:string) {

//      this.name = name;

//     }

//     eat() { //具体吃什么 不知道

//       console.log('吃')

//     }

// }

// class Dog extends Animal{

//      constructor(name:string){

//          super(name)

//      }

//      eat(){

//          return this.name+'吃肉'

//      }

// }

// class Cat extends Animal{

//     constructor(name:string){

//         super(name)

//     }

//     eat(){

//         return this.name+'吃鱼'

//     }

// }

//抽象类

//typescript 中的抽象类：它是提供其他类继承的基类,不能直接被实例化

//用abstract 关键字定义抽象方法,抽象类中的抽象方法不包含具体实现并且必须在派生类中实现

//abstract 抽象方法只能放在抽象类里面

//抽象类和抽象方法用来定义标准 , 标准：Animal 这个类要求这个子类必须包含eat方法

// abstract class Animal { //语法 抽象类只能用来定义标准 无法被实例化

//     public name: string;

//     constructor(name: string) {

//         this.name = name

//     }

//     abstract eat(): void;

// }

// class Dog extends Animal {

//     constructor(name: string) {

//         super(name)

//     }

//      //抽象类的子类必须实现抽象类里面的抽象方法

//     eat() {

//    console.log(this.name+'粮食')

//     }

// }

// let dog = new Dog('小黑')

// dog.eat()

/*

接口的做用：在面向对象的编程中,接口是一种规范的定义，它定义了行为和动作的规范，在程序设计里面，

接口是起到一种限制和规范的作用。接口的定义了某一批类所需要的遵守的规范，接口不关心这些类的内部状态数据，

也不关心这些类里方法和实现细节，它只规定这批类里面必须提供某些方法，提供这些方法的类可以满足实际的需要.

typescript中的接口类似于java，同时还增加了更灵活的接口类型,包括属性、函数、可索引和类等.

*/

/*

typescript中的接口

5.1 属性类的接口

5.2 函数类型的接口

5.3 可索引接口

5.4 类类型接口

5.5 接口扩展

*/

//1 属性接口 对json的约束

//ts中自定义方法传入参数对json进行约束

// function dog (food:{mlik:string}):void{ //这个平常的用法

//     console.log(food)

// }

// dog({mlik:'牛奶'}) //传入的时候必须传入上面约束好的mlik字段

//对批量方法传入参数进行约束.

//接口：行为和动作的规范， 对批量方法进行约束

// interface FullName{ //定义接口

//     name:string; //接口必须以;结束

//     age:number;

// }

// let obj = {

//     name:'小明',

//     age:18,

//     eat:'dog'

// }

// function printName(name:FullName){ //奇怪的使用方式

//     //必须传入对象  name 和age

//     console.log(name.name);

//     console.log(name.age);

// }

// // printName('小明'，123)错误的写法

// // printName({ //正确的写法  这样写必须一 一对应

// //     name:'张三',

// //     age:123

// // })

// printName(obj) //这样写不会报错

//接口可选属性

// interface FullName{

//     name:string;

//     age?:number; //加上这个可传可不传

// }

// function getName (info:FullName){

//  console.log(info.name,info.age)

// }

// getName({

//     name:'小明',

//     age:189

// })

// interface RequestCinfig {

//     type: string;

//     url: string;

//     ascync: boolean;

//     data?: object;

// }

// function myajax(config:RequestCinfig) {

//     let ajax: any;

//     if (XMLHttpRequest) {

//         ajax = new XMLHttpRequest()

//     } else {

//         ajax = new ActiveXObject('Microsoft.XMLTTP')

//     }

//     ajax.open(config.type,config.url,config.ascync)

//     ajax.send(config.data)

//     ajax.onreadystatechange = () => {

//         if (ajax.readyState == 4 && ajax.status == 200) {

//          console.log(ajax.responseText)

//         } else {

//             console.log(ajax.responseText)

//         }

//     }

// }

// myajax({

//     type:'get',

//     url:'www.baidu.com',

//     ascync:true,

//     data:{

//         name:'小明',

//         age:18

//     }

// })

//函数类型接口:对方法传入的参数 以及反回值进行约束

// interface fun{

//     (key:string,value:string):string;

// }

// let add:fun = function(key:string,value:string):string{ //要这样用

//     return ''

// }

//可索引接口:数组,对象的约束

// ts 定义数组的方式

// let arr:number[]=[123456,12,13,1,3,413,46]

// let arr1:Array=['123']

//可索引接口:数组,对象的约束

// interface userArr{

//     [index:number]:string

// }

// var arr:userArr=['12346','13246']

//可索引接口 对对象的约束

// interface UserObj{

//     [index:string]:string

// }

// let obj:UserObj={name:'小明'}

//类类型接口 对类的约束 和抽象类有点相似

// interface CLASS{

//     name:string;

//     eat(str:string):void;

// }

// class Dog implements CLASS{

//     name:string;

//     constructor(){

//         this.name = name;

//     }

//     eat(str:string){

//     }

// }

//接口扩展; 接口可以继承接口

// interface Animal {  //接口继承

//     eat(): void;

// }

// interface Person extends Animal {

//     work(): void;

// }

// class web implements Person {

//     constructor() {

//     }

//     eat() {

//     }

//     work() {

//     }

// }

// typescript中的泛型

// 1. 泛型的定义

// 2. 泛型的函数

// 3. 泛型类

// 4  泛型接口

/*

泛型:软件工程中 我们不仅要创建一致的定义良好的api,同时也要考虑到可重用性

组件不仅能够支持当前的数据类型。同时也能够支持未来的数据类型，这在创建大型系统时为你提供十分灵活的功能.

在想c#和java这样的语言中可以使用泛型来创建可重复用的组件,一个组件可以支持多种类型的数据,

以自己的数据类型来使用组件

通俗理解 泛型就是解决 类 接口 方法的复用性 以及对不特定数据类型的支持

*/

//只能返回string类型的数据

// function getData(value:string):string{

//     return value;

// }

//如果想不用any的情况的返回其他类型

//泛型：可以支持不特定的数据类型

//写法一

// function getData(value:T):T{ //语法

//     return value;

// }

// getData(123) //表示使用number类型

// getData('123')//很灵活

// function getData(value:T):number{ //语法

//     return 12346;

// }

// getData('你好') //传入的时候使用字符串 返回number类型

//泛型类

// class myajax{

//     public list: T[] = [];

//     add(value: T): void {

//         this.list.push(value);

//     }

// }

// var mi = new myajax(); //表示实例化对象并且顶住T的类型是number类型

//泛型接口

// interface Config{

//     (value:T,age:T):T;

// }

// var getData:Config = function(value:T):T{

//   return value

// }

// getData(1,2) // 定义了 number类型

//第二种泛型接口写法

// interface Config{

//     (value:T,age:T):T;

// }

// function getData(value:T):T{

//     return value

// }

// let mygetData:Config=getData // 定义了 number类型

// mygetData(1,2)

//TS命名空间

// 1. 定义命名空间（namespace）

// namespace modell {

//导出方法

//     export class printf {

//     }

// }

/*

导入模块 这种方式要在 html页面引入所以js文件不好

/// 

let mi = new modell.printf()

*/

//用ES6的吧

// CURD
