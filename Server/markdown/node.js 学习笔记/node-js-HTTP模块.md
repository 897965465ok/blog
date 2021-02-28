# node-js-HTTP模块

今天来写一下日记，一个原因是老是忘记，第2个当复习一下,下次可以当说明书看了，废话少说上代码。

---

```node.js
var http =require("http")//引入http模块
http.createServer((req,res)=>{

res.end("hello world");

}).listen(3800,"127.0.0.1");
```

//listen(3800,"127.0.0.1");在127.0.0.1 3800端监听

/_createServer是创建服务器()=>{};代表一个函数,函数里面由系统发送两个参数req,res, 当然你也可以自己起名字,这两个参数有什么用呢,下一章讲_/
