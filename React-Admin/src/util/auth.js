import Cookies from "js-cookie"
const handler = {
    get(targert, property) {
        return targert.getCookie(property)
    },
    set(targert, property,value) {
        return  targert.setCookie(property,value)
    }
}
const cookieObje = new Proxy({
    token: '',
    setCookie(key, value, expires = 7, path) {
        return Cookies.set(key, value, { expires, path })
    },
    getCookie(key) {
        return Cookies.get(key);
    },
    removeCookie(key) {
        return Cookies.remove(key);
    },
    getCookieAndTojson(key) {
        return Cookies.get(key);
    },
    getAllCookies() {
        return Cookies.get();
    }
}, handler)
// 将对象锁定
Object.preventExtensions(cookieObje)
export default cookieObje

