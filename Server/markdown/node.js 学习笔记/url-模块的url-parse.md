# url-模块的url-parse

前面说的http的req,res今天+url模块继续复习

```node.js
var http=require("http");//引用http
var url=require("url"); //引用url
http.createServer((req,res)=>{
var jinashu=url.parse("http://www.jianshu.com/writer#/notebooks/15144949/notes/15524193/preview",true);//默认为false是返回字符串,true是返回对象,
console.log(jinashu);//在控制台输出
}).listen(3800,"127.0.0.1");
```

//url模块的parse是把一个完整的网址拆分成12块详情请看图。

![](http://upload-images.jianshu.io/upload_images/7265016-ae2a0aed0906ba0e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=QQ%E5%9B%BE%E7%89%8720170807231439.png)

//常用host,port,pathname,path,query;

这12个块代表什么呢，相信英语差的同学一定不知道吧,

根据群里某位大神得来的结果请看下,

1 protocol 协议 http

2 slashes 斜杠 true代表有

3 auth 授权 null 路径里是否包含密码

4 host 主机 www.jianshu.com

5 port 端口比如127.0.0.1/3800这个3800就是端口

6 hostname 主机名www.jianshu.com

7 hash 哈希 #以及后面的

8 search 搜索  search它和 query 类似区别在于 search 会把 ? 以及后面的内容完整地用字符串显示

9 query 查询？后面的比如?name="gay"&age="28"

10 pathname 路径名

11 path 路径 路径名

12href 链接 完整链接
