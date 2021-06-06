# DOM-节点操作

# DOM节点操作

js对每个document文档都是抽象化成了一棵树，看下面

![](https://upload-images.jianshu.io/upload_images/7265016-0953237530fd782a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image.png)

### 获取节点的方法有很多种下面列出常用几种。

**通过id获取**。

`document.getElementById("idName")`

**通过标签获取，返回的是数组**

`document.getElementsByTagName("body")`

**通过class获取, 返回的是数组**

`document.getElementsByCalssName(" className ")`

**通过标签获取body下面的所有子节点**

`document.getElementsByTagName("body")[0].childNodes`

## 节点的类型

![](https://upload-images.jianshu.io/upload_images/7265016-c9d5eb814fce6afd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image.png)

#### 节点类型有12种我们常用到的有3种

**1.元素节点**

**2.属性节点**

**3.文本节点**

**使用nodeType方法判断这个节点是什么类型**

`Element.nodeType`

**1.元素节点返回的是1**

**2.属性节点返回的是2**

**3.文本节点返回的是3** nodeTextValue只对文本节点有用

### 获取属性值的方法和设置方法

**获取属性值方法1**

```
var tmp = document.getElementById("idName");
var a_url = tmp.getAttribute("src");
```

**获取属性值方法2**

```
var tmp = document.getElementById("idName");
var a_url = tmp.src;
```

**设置属性方法1**

```
var tmp = document.getElementById("idName");
 tmp.setAttribute("src","www.baidu.com");
```

**设置属性方法2**

```
var tmp = document.getElementById("idName");
 tmp.src=www.baidu.com;
```

###获取子节点和兄弟或者父节点

**检查是否有子节点hasChildNode 如果有返回true否则返回false**

`Element.hasChildNode`

**获取子节点childNodes以数组的形式返回**

`document.getElementsByTagName("body").childNodes`

**获取第一个子节点firstChild**

`document.getElementsByTagName("body").firstChild`

**获取最后一个子节点flastChild**

`document.getElementsByTagName("body").flastChild`

**获取父节点parentNode**

`document.getElementsByTagName("li").parentNode`

**获取下一个兄弟节点nextSbiling**

`document.getElementsByTagName("li").nextSbiling`

**获取上一个兄弟节点previousSbiling**

`document.getElementsByTagName("li").previousSbiling`

###动态创建HTML内容

**创建元素creatElement**

`var newTag = document.creatElement("p")`

**创建文本标签creatTextNode**

`var newText = document.creatTextNode("我是渣渣辉")`

**删除节点removeChild()**

`newText.removeChild (p)`

**替换节点replaceChild()**

`Element.replaceChild(新节点,需要被删除的节点)`

**插入节点方法1 appendChild()**

`newTag.appendChild(newText)`

**插入节点方法2 insertBefore()**

`newTag.insertBefore(newText, 插入的位置比如newTag.length-1)`

**复制节点 true代表克隆整个节点和子孙节点,false代表只克隆自己**

`p_count.cloneNode(true)`

###文本写入

`document.write("<p>哈哈哈哈</p>") ,Element.innerHTML`
