import axios from "axios"
import cookie from "./auth"
cookie.token = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYyMDA3NTY4NCwiaWF0IjoxNjIwMDU0MDg0LCJpc3MiOiJKaWFuZyIsInN1YiI6InVzZXIgdG9rZW4ifQ.34LwYzN11BasQnGovZhQfm4YYmruSnY8ZoKPnwBVm6o"
const http = axios.create({
    timeout: 5000
})
http.interceptors.request.use(config => {
    if (cookie.token) {
        config.headers.authorization = cookie.token
    } else {
        console.log(cookie.token, "不存在提示用户登录")
    }
    return config
})

// let veriffer[] 

// 50008:非法令牌; 50012:其他客户已登录 ;50014:令牌过期;
http.interceptors.response.use(config => {
    const { status, data } = config
    const code = data.code
    switch (status) {
        case 200: {

            console.log("响应成功")
            break
        }
        case 500: {
            console.log("错误")
            break
        }
        case 404: {
            console.log("资源不存在")
            break
        }
        case 50008: {
            console.log("非法令牌")
            break
        }
        case 50012: {
            console.log("其他客户已登录")
            break
        }
        case 50014: {
            console.log("令牌过期")
            break
        }

    }
    return config
})

export default http