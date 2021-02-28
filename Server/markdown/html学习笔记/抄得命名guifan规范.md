# 抄得命名guifan规范


作者：知乎用户

链接：[https://www.zhihu.com/question/19586885/answer/48933504](https://www.zhihu.com/question/19586885/answer/48933504)

来源：知乎

著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

**CSS书写顺序**

1.位置属性(position, top, right, z-index, display, float等)

2.大小(width, height, padding, margin)

3.文字系列(font, line-height, letter-spacing, color- text-align等)

4.背景(background, border等)

5.其他(animation, transition等)

![](https://pic4.zhimg.com/50/5a67fa4dab92c018abfd97e4ab286ac9_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-a64ad085226e5085..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

**CSS书写规范使用CSS缩写属性**

CSS有些属性是可以缩写的，比如padding,margin,font等等，这样精简代码同时又能提高用户的阅读体验。

![](https://pic1.zhimg.com/50/2166c1fb168908a5858998eab2d309aa_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-4dbe5c15b4090684..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

去掉小数点前的“0”

![](https://pic2.zhimg.com/50/ea65b1c1c00c792866a670e9d66161a8_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-19ae17aefa8ed1ac..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

简写命名

很多用户都喜欢简写类名，但前提是要让人看懂你的命名才能简写哦!

![](https://pic1.zhimg.com/50/53ddaefb44a6a8ce5b54de9f15466fca_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-87317aa86ed6de63..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

16进制颜色代码缩写

有些颜色代码是可以缩写的，我们就尽量缩写吧，提高用户体验为主。

![](https://pic1.zhimg.com/50/46bf23e08a7195ec7091217bb5c667d6_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-c9fc028b571b89c1..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

连字符CSS选择器命名规范

1.长名称或词组可以使用中横线来为选择器命名。

2.不建议使用“_”下划线来命名CSS选择器，为什么呢?

输入的时候少按一个shift键;

浏览器兼容问题 (比如使用_tips的选择器命名，在IE6是无效的)

能良好区分JavaScript变量命名(JS变量命名是用“_”)

![](https://pic1.zhimg.com/50/45dddce1b5295a7c09f2380900dce8d1_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-e12ed2ec1a4716f6..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

不要随意使用Id

id在JS是唯一的，不能多次使用，而使用class类选择器却可以重复使用，另外id的优先级优先与class，所以id应该按需使用，而不能滥用。

![](https://pic1.zhimg.com/50/951bee68367f559a1fd37db3d242e59f_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-297beae467227382..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

为选择器添加状态前缀

有时候可以给选择器添加一个表示状态的前缀，让语义更明了，比如下图是添加了“.is-”前缀。

![](https://pic4.zhimg.com/50/d81807b01ec5d710debe3f0b4191ae2f_hd.jpg#width=500)
![](http://upload-images.jianshu.io/upload_images/7265016-387517dbde1c6e44..jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

CSS命名规范(规则)常用的CSS命名规则

头：header

内容：content/container

尾：footer

导航：nav

侧栏：sidebar

栏目：column

页面外围控制整体佈局宽度：wrapper

左右中：left right center

登录条：loginbar

标志：logo

广告：banner

页面主体：main

热点：hot

新闻：news

下载：download

子导航：subnav

菜单：menu

子菜单：submenu

搜索：search

友情链接：friendlink

页脚：footer

版权：copyright

滚动：scroll

内容：content

标签：tags

文章列表：list

提示信息：msg

小技巧：tips

栏目标题：title

加入：joinus

指南：guide

服务：service

注册：regsiter

状态：status

投票：vote

合作伙伴：partner

**注释的写法:**

/_ Header _/

内容区

/_ End Header _/

**Id的命名:**

1)页面结构

容器: container

页头：header

内容：content/container

页面主体：main

页尾：footer

导航：nav

侧栏：sidebar

栏目：column

页面外围控制整体佈局宽度：wrapper

左右中：left right center

**(2)导航**

导航：nav

主导航：mainnav

子导航：subnav

顶导航：topnav

边导航：sidebar

左导航：leftsidebar

右导航：rightsidebar

菜单：menu

子菜单：submenu

标题: title

摘要: summary

**(3)功能**

标志：logo

广告：banner

登陆：login

登录条：loginbar

注册：register

搜索：search

功能区：shop

标题：title

加入：joinus

状态：status

按钮：btn

滚动：scroll

标籤页：tab

文章列表：list

提示信息：msg

当前的: current

小技巧：tips

图标: icon

注释：note

指南：guild

服务：service

热点：hot

新闻：news

下载：download

投票：vote

合作伙伴：partner

友情链接：link

版权：copyright

**注意事项::**

1.一律小写;

2.尽量用英文;

3.不加中槓和下划线;

4.尽量不缩写，除非一看就明白的单词。

**CSS样式表文件命名**

主要的 master.css

模块 module.css

基本共用 base.css

布局、版面 layout.css

主题 themes.css

专栏 columns.css

文字 font.css

表单 forms.css

补丁 mend.css

打印 print.css

[发布于 2015-05-25](/question/19586885/answer/48933504)
