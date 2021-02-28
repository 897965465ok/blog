# myajax

```

ajax.readyState
0(未初始化)还没有调用open()方法
1(载入) 已经调用send()方法,正在发送请求
2(载入完成)send方法完成, 已经收到全部响应内容
3(解析）正在解析响应内容
4(完成) 响应内容解析完成，可以在客户端调用了
  function myAjax(url){
      let ajax = null;
      if(window.XMLHttpRequest){  //兼容主流浏览器
       ajax = new XMLHttpRequest();
      }else{
          ajax = new ActiveXObject('Microsoft.XMLHTTP'); // ie6-7
      }
      ajax.open('GET',url,true);
      ajax.send();
      ajax.onreadystatechange=function(){
         if(ajax.readyState == 4){ //读完完成不代表成功
            if(ajax.status == 200){ //得判断这个才行
                alert(ajax.responseText); //内容
            }else{
                alert(ajax.responseText);
            }
         }
      }
  }   
  myAjax('app.js')
```
