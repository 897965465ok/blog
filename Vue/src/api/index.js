import axios from 'axios'
let api = axios.create({
    // baseURL:"http://146.56.206.160:3800",
})
api.interceptors.response.use((config) => {
    let { status, data } = config
    switch (status) {
        case 200: {
            if (data.code == 200) {
                return config
            }
        }
        case 400: {
            // Message.error("服务器出错")
            return config
        }
        default: {
            return config
        }
    }
    return config
},
    (error) => {

        return new Promise.reject(error)
    })

api.interceptors.request.use((config) => {
    let token = window.localStorage.getItem("token")
    if (token) {
        config.headers.authorization = token;    //将token放到请求头发送给服务器
        return config;
    } else {
        return config
    }
}, (error) => {
    // Do something with request error
    return Promise.reject(error);
});

export { api }
