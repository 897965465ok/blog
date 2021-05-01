# ---------formidable表单提交模块

今天来说一下   formidable, 这个模块的主要功能上拿来提交表单和，文件

接下来看怎么用。
```
if (req.url=="/favicon.ioc"){return ;};

if(req.url=="/dopost"&&req.method.toLowerCase()=="post"){

	var form = new formidable.IncomingForm();      //先引用对象

 form.uploadDir = __dirname+"/opop"; //设置路径

 form.keepExtensions = true;  //是否保存文件扩展名是
```

//执行里面的回调函数的时候，表达已经全部接收完毕了. 第2个是表单第三个是文件

form.parse(req, function(err, fields, files){

var extname=path.extname(files.tupian.name);// 获取文件扩展名

fs.rename(files.tupian.path,__dirname+"/opop/897965465"+extname,(err)=>{

res.writeHead(200, { 'Content-Type': 'text/plain;charset=UTF-8' });

if (err) {console.log("改名失败")};

});

[//console.log](//console.log)(files);

res.end();

});

return ;

}

}).listen(3800,"127.0.0.1");

