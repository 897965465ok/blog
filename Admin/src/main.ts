import { createApp, createVNode,provide} from 'vue'
import App from './App.vue'
import { router } from "./router"
import { store } from "./store"
import './index.css'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as Icons from '@element-plus/icons-vue'

// 创建Icon组件
const Icon = (props: { icon: string }) => {
    const { icon } = props
    // icon as (keyof typeof Icons)
    // keyof 约束只能值只能为对象的key
    return createVNode(Icons[icon as (keyof typeof Icons)])
}
createApp(App)
    .use(ElementPlus)
    .use(router)
    .use(store)
    .component('Icon', Icon)
    .mount('#app')

  // 注册Icon组件
