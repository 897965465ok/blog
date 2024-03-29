'use strict'
// Template version: 1.3.1
// see http://vuejs-templates.github.io/webpack for documentation.

const path = require('path')

module.exports = {
  dev: {
    // Paths
    assetsSubDirectory: 'static',
    assetsPublicPath: '/',
    proxyTable: {
      '/v1': {
        target: 'http://localhost:3800',  //目标接口域名
        changeOrigin: true,  //是否跨域
        // secure: true, // 如果是 https ,需要开启这个选项
        pathRewrite: {
          '^/v1': "/v1"                  //如果地址不想带api就这样替换掉
        }
      },
   
      // 'https://pixabay.com/api': {
      //   target: 'https://pixabay.com/api',  //目标接口域名
      //   changeOrigin: true,  //是否跨域
      //   secure: false, // 如果是 https ,需要开启这个选项
      //   pathRewrite: {
      //     '^/https://pixabay.com/api': ""                  //如果地址不想带api就这样替换掉
      //   }
      // }
    '/wallhaven':{
      target: 'https://wallhaven.cc/api/v1/search',  //目标接口域名
        changeOrigin: true,  //是否跨域
        secure: true, // 如果是 https ,需要开启这个选项
        pathRewrite: {
          '^/wallhaven': ""                  //如果地址不想带api就这样替换掉
        }
      },
      '/ws': {
        target: 'ws://localhost:3800/ws',  //目标接口域名
        changeOrigin: true,//是否允许跨域
        pathRewrite: {
          '^/ws': ""                  //如果地址不想带api就这样替换掉
        },
        ws:true,
      },
      '/Oauth': {
        target: 'https://gitee.com/oauth/authorize',  //目标接口域名
        changeOrigin: true,//是否允许跨域
        pathRewrite: {
          '^/Oauth': ""                  //如果地址不想带api就这样替换掉
        },
        secure: true, // 如果是 https ,需要开启这个选项
      },
    },
    // Various Dev Server settings
    host: 'localhost', // can be overwritten by process.env.HOST
    port: 8080, // can be overwritten by process.env.PORT, if port is in use, a free one will be determined
    autoOpenBrowser: false,
    errorOverlay: true,
    notifyOnErrors: true,
    poll: false, // https://webpack.js.org/configuration/dev-server/#devserver-watchoptions-


    /**
     * Source Maps
     */

    // https://webpack.js.org/configuration/devtool/#development
    devtool: 'cheap-module-eval-source-map',

    // If you have problems debugging vue-files in devtools,
    // set this to false - it *may* help
    // https://vue-loader.vuejs.org/en/options.html#cachebusting
    cacheBusting: true,

    cssSourceMap: true
  },

  //编译后路径存放的配置
  build: {
    // Template for index.html
    index: path.resolve(__dirname, '../../Server/view/index.html'),
    // Paths
    assetsRoot: path.resolve(__dirname, '../../Server/view'),
    assetsSubDirectory: 'static',
    assetsPublicPath: '/',

    /**
     * Source Maps
     */

    productionSourceMap: false,
    // https://webpack.js.org/configuration/devtool/#production
    devtool: '#source-map',

    // Gzip off by default as many popular static hosts such as
    // Surge or Netlify already gzip all static assets for you.
    // Before setting to `true`, make sure to:
    // npm install --save-dev compression-webpack-plugin
    productionGzip: true,
    productionGzipExtensions: ['js', 'css'],

    // Run the build command with an extra argument to
    // View the bundle analyzer report after build finishes:
    // `npm run build --report`
    // Set to `true` or `false` to always turn it on or off
    bundleAnalyzerReport: process.env.npm_config_report
  }
}
