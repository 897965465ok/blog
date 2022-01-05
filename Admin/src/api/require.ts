import axios, { AxiosRequestHeaders, AxiosRequestConfig } from 'axios'
const api = axios.create({
    timeout: 10000, // 设置统一的超时时长
})

// api.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
api.interceptors.response.use((config) => {
    let { status, data } = config
    switch (status) {
        case 200: {
                return config
    
        }
        case 400: {

            return config
        }
        case 500: {
            return config
        }
        case 302: {
            return config
        }
        default: {
            return config
        }
    }
}, (error) => {
    return Promise.reject(error)
})

api.interceptors.request.use((config) => {
    let token = window.localStorage.getItem("token")
    if (token) {
        (config.headers as any).authorization = token;
    }
    return config
}, (error) => {
    return Promise.reject(error);
});
export { api }