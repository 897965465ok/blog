# vue-js-初体验

### v-model 数据绑定

```js
<div id="myApp">
    <p>喜欢的游戏是:{{ooxx}}</p>
    <p>最喜欢的女孩是:{{girl}}</p>
    <input type="text" v-model="ooxx"><br>
    <input type="text" v-model="girl">
    //因为是双向绑定所以改这个data也会随着改变
 </div>
```

通过v-model绑定 input

```js
var myApp = new Vue({
    el: "#myApp",// 绑定那个div  这里是#myApp
    data:{     
        ooxx:"尾行", 
        girl:"波多野结衣",
    },
 });
```
