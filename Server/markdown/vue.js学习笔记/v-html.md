# v-html

### V-HTML

由于{{}}只能输出变量 ，有时候需要插入标签怎么办，

可以使用v-html

```js
 <div id="myapp">
 {{ooxx}}
 <div v-html="message"></div> 
//把message值 也就是H3标签插入当前div
 </div>
```

```js
var myapp = new Vue({
  el: "#myapp",
  data:{
  ooxx:'hello vue.js ',
  message:"<h3>i am not gay</h3>",
  }
})
```
