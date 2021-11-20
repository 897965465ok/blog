const {
    createProxyMiddleware
} = require('http-proxy-middleware')
module.exports = app => {
    app.use(
        createProxyMiddleware('/v1', {
            "target": "http://127.0.0.1:3800",
            "changeOrigin": true,
        })
    );
    // app.use(
    //     createProxyMiddleware('/developer', {
    //         "target": "https://gate.snssdk.com",
    //         "changeOrigin": true,
    //     })
    // )
}