# v-on-methods 事件

#vue 事件

```js
<div id="App">
        <p>{{mygame}}</p>
      //  mygame 是属性值
        <button v-on:click="clickme('do you like game ?')" >no</button>
 // v-on:click="clickme('do you like game')" v-on 监听事件,click是点击事件类型
        <button @click="clickme('yes i am like game')" > yes</button>
//@click事件监听简写@click="clickme('yes iam like game')"
</div>
```

```js
   var App = new Vue({
        el: "#App",
        data: { 数据区//data
            mygame: "超级马里奥",
        },
        methods: { //函数区 methods
            clickme: function (value) {
                this.mygame = value;
            },
        },
    });
```
