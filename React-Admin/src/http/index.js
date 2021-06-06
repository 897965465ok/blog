import http from '../util/request'
import qs from "axios"
export default {
    async login(account, password) {
        return await http.post('v1/login', {
            data: {
                account,
                password
            }
        })
    },
    async getArticles() {
        return await http.get('v1/articles')
    }
}


