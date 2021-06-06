# promise

1. promise 有三种状态，进行中   成功(resolve) 失败(reject)
2. promise 只能是进行中到成功或者失败
3. promise 使用场景异步的时候使用

## 语法

```javascript
let test = new Promise((resolve, reject) => {
  console.log("我在new的时候执行");
  if (true) {
    resolve("成功"); // 成功的时候调用这个
  } else {
    reject("失败"); // 失败的时候调用这个
  }
});
test.then(
  (resolveVlaue) => {
    conosle.log(resolveVlaue);
  },
  (rejectValue) => {}
);
```

## 使用方法

```javascript
   function loadImg(url) {
       return new Promise((resolve, reject) => {
           let img = new Image();
           img.src = url;
           img.style.width = "300px"
           img.style.height = "300px"
           img.onload = function () {
               resolve(img)
           }
           img.onerror = function (event) {
               reject(event)
           }
       })
   }
   loadImg('https://i0.hdslb.com/bfs/sycp/creative_img/201810/1bdcd241c6e44c532235c1cd2fafeebd.jpg')
       .then((img) => {
           document.getElementsByTagName('body')[0].append(img)
       }).catch(()=>{}) //效果和失败函数一样
```

## promise.all()   作用(并发) 多个实例一起调用返回结果

```javascript
let a = Promise.resolve("aaa");
//快速创建实例 等于   new Promise((resolve,reject)=>{ resolve('aaa'))
let b = Promise.reject("bbb");
let c = Promise.all([a, a, a, b]);
c.then((value) => {
  console.log(value); //如果正确返回的是一个数组
}).catch((value) => {
  console.log(value); //如果错误返回的是错误的那个
});
```

## Promise.race()
```javascript
let a = Promise.resolve("aaa");
let b = Promise.reject("bbb");
let c = Promise.race([b, a, a, b]);
c.then((value) => {
  console.log(value);
});
//谁先运行完返回谁;
```

# 状态的传递  有点恶心

```javascript
p1 = new Promise(function (resolve, reject) {
  setTimeout(() => resolve("我是p1错误的回调"), 3000);
});

p2 = new Promise(function (resolve, reject) {
  setTimeout(() => resolve(p1), 1000); //p1的状态和参数都将传给p2
});
p2.then((value) => {
  //因为p1的状态是失败所以将执行catch ,err是重p1传过来的参数
  console.log(value);
  console.log("我是p2的回调");
}).catch((err) => {
  console.log(err); // 我是p1错误的回调
  console.log("我是p2错误的回调"); // 我是p2错误的回调
});
```
