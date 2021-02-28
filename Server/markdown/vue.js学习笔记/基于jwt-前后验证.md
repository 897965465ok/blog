# 基于jwt-前后验证

```
技术栈  axios   node 
第一步 
cnpm install  jsonwebtoken --save
cnpm install koa2-jwt --save



 axios({
            method: "post",
            url: url.registerUser,
            data: {
              userName: this.userInfo.user,
              password: this.userInfo.password
            }
          })
            .then(resolve => {
              if (resolve.data.code ==200 && resolve.status == 200) {
                console.log(resolve.data)
                localStorage.userToken =  JSON.stringify(resolve.data.userToken)  //将后台返回的token 存放 存哪里随意
                this.$Message.success("注册成功");
                this.openLoading = false;
              } else {
                console.log(resolve)
                this.$Message.error("注册失败");
                this.openLoading = false;
              }
            })
            .catch(err => {
              console.log(err)
              this.$Message.error("注册失败");
              this.openLoading = false;
            });


axios.interceptors.request.use(
    config => {
        let Token = localStorage.userToken ? JSON.parse(localStorage.userToken) : ''
        if (Token) {
            config.headers.authorization = `Bearer ${Token['userToken']}`   这个Bearer  后面有个空格一定要加上
        }
        return config
    }
)

验证
const KoaJwt = require('koa-jwt');
app.use(KoaJwt({
    secret,
    //initSchemas:true
}).unless({
    path: [/^\/user\/register/,/^\/user\/login/], //数组中的路径不需要通过jwt验证
     method: 'GET' 不需要验证的方法  
}))

//jwt.js   
/*
token 默认使用hs256对称算法
我们可以使用非对称算法 RS256
签名时候使用 pricate key// 私钥 
根据私钥生成验证的时候使用 public key//公钥
签名的时候使用私钥 
认证的时候使用公钥 
*/
const fs = require('fs')
const path = require('path')
const jwt = require('jsonwebtoken')
const promisify = require('util').promisify

function verify(token, publicKey) {
    return new Promise((resolve, reject) => {
        jwt.verify(token, publicKey, (err, decoded) => {
            if (err) {
                reject(err.message)
            } else {
                resolve(decoded)
            }
        })
    })
}
function sign(playload, privateKey, options) {
    options = {
        algorithm: 'RS256',
        expiresIn: '24h'
    }
    return new Promise((resolve, reject) => {
        jwt.sign(playload, privateKey, options, (err, newtoken) => {
            if (err) {
                console.log('生成token失败了')
                reject(err.message)
            } else {
                resolve(newtoken)
            }
        })
    })
}
const jwtToken = {
    async generateToken(playload) { // 生成 密匙
        let privateKey = await promisify(fs.readFile)(path.resolve(__dirname + '/private.key'), 'utf8')
        if (privateKey) {
            return await sign(playload, privateKey)
        } else {
            return '读取密匙失败'
        }
    },
    async compareToken(token) {
        let publicKey = await promisify(fs.readFile)(path.resolve(__dirname + '/public.key'), 'utf8')
        if (publicKey) {
            return await verify(token.split(" ")[1], publicKey) // 验证用公钥 
        } else {
            return ' 读取密匙错误'
        }
    }
}
module.exports = jwtToken;
```
