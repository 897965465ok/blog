# vuex  使用详解

1. vuex 是干什么的
2. 为什么需要vuex

## 为什么需要vuex

vuex 是基于vue框架而创建的一个库

vue在开发的时候碰到兄弟组件联动数据的,很难解决。

要通过靠组件传送来解决问题。

后来开发者觉得这样操作很臃肿,vuex就应需求而生

## vuex 是干什么的

vuex 总的来说就是将那些公用的数据抽离出来的载体

![](https://upload-images.jianshu.io/upload_images/7265016-ea8ea5420597448b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=VUEX%E6%B5%81%E7%A8%8B%E5%9B%BE)

## 载体 sate

所以的数据都存放在sate对象里面，

如果你想修改就必须经过一套流程

列子：
