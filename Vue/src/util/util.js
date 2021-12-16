
import qs from 'qs'
// 如果前端配置了这个withCredentials=true，后段设置Access-Control-Allow-Origin不能为 " * ",必须是你的源地址啊
// axios.defaults.withCredentials=true
export default {
    install: function (Vue, options) {
        // 1. 添加全局方法或属性
        // Vue.myGlobalMethod = function () {
        //     // 逻辑...
        //     console.log('myGlobalMethod')
        // }

        // // 2. 添加全局资源
        // Vue.directive('my-directive', {
        //     bind(el, binding, vnode, oldVnode) {
        //         // 逻辑...
        //     }
        // })

        // // 3. 注入组件
        // Vue.mixin({
        //     created: function () {
        //         // 逻辑...
        //     }
        // })

        // 4. 添加实例方法
        https://api.ixiaowai.cn/ylapi/index.php/?code=js
        Vue.prototype.$getComments = async (articleId) => {
            return await Vue.prototype.$api.get('v1/comment', {
                params: {
                    articleId
                }
            })

        }
        Vue.prototype.$comment = async (articleId, content,replyArticle,userName) => {
            console.log(articleId, content,replyArticle,userName)
            return await Vue.prototype.$api.post('v1/comment',
                qs.stringify({
                    articleId,
                    content,
                    replyArticle,
                    userName
                }),
               )
        }

        Vue.prototype.$shouldJson = async (articleId, content) => {
            return await Vue.prototype.$api.post('v1/should',
                qs.stringify({
                    articleId,
                    content
                }))
        }

        Vue.prototype.$generateJSON = async () => {
            try {
                let data = (await Vue.prototype.$api.get('v1/returnjson')).data
                return JSON.parse(data.json)
            } catch (e) {
                return false
            }
        }

        Vue.prototype.$wallhaven = async (params) => {
            try {
                let { data } = await Vue.prototype.$api.get('v1/wallhaven', {
                    params: {
                        q: params.q,
                    }
                })
                return data.result
            } catch (e) {
                return false
            }
        }

        Vue.prototype.$random = (min, max) => {

            return Math.floor(Math.random() * (max - min)) + min;

        }
        Vue.prototype.$GetUrl = async () => {
            try {
                return (await axios.get('https://pixabay.com/api', {
                    params: {
                        key: "21226858-57a14a3bedc89005c85e668cc",
                        per_page: 200,
                        //    q:"",
                        category: "nature",

                        safesearch: true
                    }
                }))
            } catch (e) {
                return false
            }
        }
        Vue.prototype.$skip = (link) => {
            window.open(link, 'newwindow')
        }

        // 当持续触发事件时，一定时间段内没有再触发事件，事件处理函数才会执行一次，
        // 如果设定的时间到来之前，又一次触发了事件，就重新开始延时
        Vue.prototype.$debounce = function (fun, t, tagerNow) {
            let timer = null
            let debounce = (...args) => {
                if (timer) {
                    //如果定速器存在说明下面执行过了
                    clearTimeout(timer)
                }
                // 是否第一次执行
                if (tagerNow) {
                    let exec = !timer
                    // 延迟
                    timer = setTimeout(() => {
                        timer = null
                    }, t)
                    if (exec) {
                        fun.apply(this, args)
                    }
                } else {
                    timer = setTimeout(() => {
                        fun.apply(this, args)
                    }, t)
                }

            }

            debounce.remover = () => {
                clearTimeout(timer)
                timer = null
            }
            return debounce
        }
    }
}