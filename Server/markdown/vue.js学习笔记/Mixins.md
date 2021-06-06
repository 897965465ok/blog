# Mixins

```js
混入 语法
import mixins from './Mixins'
Vue.components( {
   name:'test2',
    created(){
     this.printf()
    },
    data(){
      return {
          a :1
      }
    },
   mixins:[mixins]  // 这样使用
})
```

```js
被导入的模块
import Vue from 'vue'
Vue.mixin({  //全局的混入 这样不好
    methods:{
      printf(){
          console.log('我是全局的mixin')
      }
    }
})
export default {
    data () {
        return {
            a:0
        }
    },
    methods:{
        printf(){
            console.log(this.aa)
        }
    }
}
```

```js
执行顺序
全局》混入》原生
如果属性名有冲突优先使用本地属性名 原生 》混入 》全局
 delimiters:['${','}'] 修改 默认{{}}插值 改成  ${变量}方式
  extends 执行顺序 钩子的话有是扩展 > 原生  方法的话是原生》扩展
```
