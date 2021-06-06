# [HTML5和CSS3新特性一览](http://www.cnblogs.com/star91/p/5659134.html)

#### HTML5

###### **1.HTML5** **新元素**

HTML5提供了新的元素来创建更好的页面结构：

**标签**

**描述**



定义页面独立的内容区域。

|



|

定义页面的侧边栏内容。

|



|

允许您设置一段文本，使其脱离其父元素的文本方向设置。

|



|

定义命令按钮，比如单选按钮、复选框或按钮

|



|

用于描述文档或文档某个部分的细节

|



|

定义对话框，比如提示框

|



|

标签包含 details 元素的标题

|



|

规定独立的流内容（图像、图表、照片、代码等等）。

|



|

定义 

|



|

定义 section 或 document 的页脚。

|



|

定义了文档的头部区域

|

[](http://www.runoob.com/tags/tag-mark.html)

|

定义带有记号的文本。

|



|

定义度量衡。仅用于已知最大和最小值的度量。

|



|

定义导航链接的部分。

|



|

定义任何类型的任务的进度。

|



|

定义 ruby 注释（中文注音或字符）。

|



|

定义字符（中文注音或字符）的解释或发音。

|



|

在 ruby 注释中使用，定义不支持 ruby 元素的浏览器所显示的内容。

|



|

定义文档中的节（section、区段）。

|



|

定义日期或时间。

|



|

规定在文本中的何处适合添加换行符。

|

###### **2.HTML5 Canvas**

HTML5  元素用于图形的绘制，通过脚本 (通常是JavaScript)来完成.

 标签只是图形容器，您必须使用脚本来绘制图形。


使用 JavaScript 来绘制图像

canvas 元素本身是没有绘图能力的。所有的绘制工作必须在 JavaScript 内部完成：

拖放是一种常见的特性，即抓取对象以后拖到另一个位置。在 HTML5 中，拖放是标准的一部分，任何元素都能够拖放。

设置元素为可拖放

首先，为了使元素可拖动，把 draggable 属性设置为 true ：

---

拖动什么 - ondragstart 和 setData()

---

放到何处 - ondragover

---

进行放置 - ondrop

###### **4.HTML5** **地理定位**

HTML5 Geolocation API 用于获得用户的地理位置。

鉴于该特性可能侵犯用户的隐私，除非用户同意，否则用户位置信息是不可用的。

HTML5 规定了在网页上嵌入音频元素的标准，即使用 


HTML5 规定了一种通过 video 元素来包含视频的标准方法。


###### **6.HTML5 Input** **类型**

HTML5 拥有多个新的表单输入类型。这些新特性提供了更好的输入控制和验证。

- color、date、datetime、datetime-local、email、month、number、range、search、tel、time、url、week


###### **7.HTML5** **表单元素**

HTML5 有以下新的表单元素:

| 标签 | 描述 |

|  | 标签定义选项列表。请与 input 元素配合使用该元素，来定义 input 可能的值。 |

|  | > 标签规定用于表单的密钥对生成器字段。 |

|  |  标签定义不同类型的输出，比如脚本的输出。 |

 元素规定输入域的选项列表。

 属性规定 form 或 input 域应该拥有自动完成功能。当用户在自动完成域中开始输入时，浏览器应该在该域中显示填写的选项：

使用  元素的列表属性与  元素绑定.


###### **8.HTML5** **表单属性**

HTML5 的  和 标签添加了几个新属性.

- autocomplete、novalidate

新属性：

- autocomplete、autofocus、form、formaction、formenctype、formmethod、formnovalidate、formtarget、height and width、list、min and max、multiple、pattern (regexp)、placeholder、required、step

###### **9.HTML5** **语义元素**

HTML5提供了新的语义元素来明确一个Web页面的不同部分:

- 


- 


- 

- 

- 


- 

- 

- 



![](http://upload-images.jianshu.io/upload_images/7265016-f3280016feb344d2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240#alt=image)

###### **10.HTML5 Web** **存储**

<o:p> Web Storage DOM API 为Web应用提供了一个能够替代cookie的Javascript解决方案</o:p>

- sessionStorage—客户端数据存储，只能维持在当前会话范围内。

sessionStorage 方法针对一个 session 进行数据存储。当用户关闭浏览器窗口后，数据会被删除。

- localStorage—客户端数据存储，能维持在多个会话范围内。

localStorage 对象存储的数据没有时间限制。第二天、第二周或下一年之后，数据依然可用。

对于大量复杂数据结构，一般使用IndexDB

###### **11.HTML5** **离线Web应用（应用程序缓存）**

HTML5 引入了应用程序缓存，这意味着 web 应用可进行缓存，并可在没有因特网连接时进行访问。

应用程序缓存为应用带来三个优势：

1. 离线浏览 - 用户可在应用离线时使用它们
2. 速度 - 已缓存资源加载得更快
3. 减少服务器负载 - 浏览器将只从服务器下载更新过或更改过的资源。

HTML5 Cache Manifest 实例

下面的例子展示了带有 cache manifest 的 HTML 文档（供离线浏览）：


**Manifest 文件**

manifest 文件是简单的文本文件，它告知浏览器被缓存的内容（以及不缓存的内容）。

manifest 文件可分为三个部分：

- 
_CACHE MANIFEST_ - 在此标题下列出的文件将在首次下载后进行缓存

- 
_NETWORK_ - 在此标题下列出的文件需要与服务器的连接，且不会被缓存

- 
_FALLBACK_ - 在此标题下列出的文件规定当页面无法访问时的回退页面（比如 404 页面
1. 
`CACHE MANIFEST`

2. 
`# 2012-02-21 v1.0.0`

3. 
`/theme.css`

4. 
`/logo.gif`

5. 
`/main.js`

6. 
`NETWORK:`

7. 
`login.php`

8. 
`FALLBACK:`

9. 
`/html/ /offline.html`




###### **12.HTML5 Web Workers**

当在 HTML 页面中执行脚本时，页面的状态是不可响应的，直到脚本已完成。

web worker 是运行在后台的 JavaScript，独立于其他脚本，不会影响页面的性能。您可以继续做任何愿意做的事情：点击、选取内容等等，而此时 web worker 在后台运行。（相当于实现多线程并发）

###### **13.HTML5 SSE**

Server-Sent 事件指的是网页自动获取来自服务器的更新。

以前也可能做到这一点，前提是网页不得不询问是否有可用的更新。通过服务器发送事件，更新能够自动到达。

例子：Facebook/Twitter 更新、估价更新、新的博文、赛事结果等。

EventSource 对象用于接收服务器发送事件通知：



为了让上面的例子可以运行，您还需要能够发送数据更新的服务器（比如 PHP 和 ASP）。


###### **14.HTML5 WebSocket**

WebSocket是HTML5开始提供的一种在单个 TCP 连接上进行全双工通讯的协议。在WebSocket API中，浏览器和服务器只需要做一个握手的动作，然后，浏览器和服务器之间就形成了一条快速通道。两者之间就直接可以数据互相传送。浏览器通过 JavaScript 向服务器发出建立 WebSocket 连接的请求，连接建立以后，客户端和服务器端就可以通过 TCP 连接直接交换数据。当你获取 Web Socket 连接后，你可以通过 **send()** 方法来向服务器发送数据，并通过 **onmessage** 事件来接收服务器返回的数据。以下 API 用于创建 WebSocket 对象。

#### CSS3

###### **CSS3选择器**

| 选择器 | 示例 | 示例说明 | CSS |

| [._class_](http://www.runoob.com/cssref/sel-class.html) | .intro | 选择所有class="intro"的元素 | 1 |

| [#_id_](http://www.runoob.com/cssref/sel-id.html) | #firstname | 选择所有id="firstname"的元素 | 1 |

| [*](http://www.runoob.com/cssref/sel-all.html) | * | 选择所有元素 | 2 |

| _[element](http://www.runoob.com/cssref/sel-element.html)_ | p | 选择所有
元素 | 1 |

| _[element,element](http://www.runoob.com/cssref/sel-element-comma.html)_ | div,p | 选择所有
元素 | 1 |

| [_element_ _element_](http://www.runoob.com/cssref/sel-element-element.html) | div p | 选择
元素 | 1 |

| [_element_>_element_](http://www.runoob.com/cssref/sel-element-gt.html) | div>p | 选择所有父级是 
 元素 | 2 |

| [_element_+_element_](http://www.runoob.com/cssref/sel-element-pluss.html) | div+p | 选择所有紧接着
元素 | 2 |

| [[_attribute_]](http://www.runoob.com/cssref/sel-attribute.html) | [target] | 选择所有带有target属性元素 | 2 |

| [[_attribute_=_value_]](http://www.runoob.com/cssref/sel-attribute-value.html) | [target=-blank] | 选择所有使用target="-blank"的元素 | 2 |

| [[_attribute_~=_value_]](http://www.runoob.com/cssref/sel-attribute-value-contains.html) | [title~=flower] | 选择标题属性包含单词"flower"的所有元素 | 2 |

| [[_attribute_|=_language_]](http://www.runoob.com/cssref/sel-attribute-value-lang.html) | [lang|=en] | 选择一个lang属性的起始值="EN"的所有元素 | 2 |

| [:link](http://www.runoob.com/cssref/sel-link.html) | a:link | 选择所有未访问链接 | 1 |

| [:visited](http://www.runoob.com/cssref/sel-visited.html) | a:visited | 选择所有访问过的链接 | 1 |

| [:active](http://www.runoob.com/cssref/sel-active.html) | a:active | 选择活动链接 | 1 |

| [:hover](http://www.runoob.com/cssref/sel-hover.html) | a:hover | 选择鼠标在链接上面时 | 1 |

| [:focus](http://www.runoob.com/cssref/sel-focus.html) | input:focus | 选择具有焦点的输入元素 | 2 |

| [:first-letter](http://www.runoob.com/cssref/sel-firstletter.html) | p:first-letter | 选择每一个
元素的第一个字母 | 1 |

| [:first-line](http://www.runoob.com/cssref/sel-firstline.html) | p:first-line | 选择每一个
元素的第一行 | 1 |

| [:first-child](http://www.runoob.com/cssref/sel-firstchild.html) | p:first-child | 指定只有当
元素是其父级的第一个子级的样式。 | 2 |

| [:before](http://www.runoob.com/cssref/sel-before.html) | p:before | 在每个
元素之前插入内容 | 2 |

| [:after](http://www.runoob.com/cssref/sel-after.html) | p:after | 在每个
元素之后插入内容 | 2 |

| [:lang(_language_)](http://www.runoob.com/cssref/sel-lang.html) | p:lang(it) | 选择一个lang属性的起始值="it"的所有
元素 | 2 |

| [_element1_~_element2_](http://www.runoob.com/cssref/sel-gen-sibling.html) | p~ul | 选择p元素之后的每一个ul元素 | 3 |

| [[_attribute_^=_value_]](http://www.runoob.com/cssref/sel-attr-begin.html) | a[src^="https"] | 选择每一个src属性的值以"https"开头的元素 | 3 |

| [[_attribute_![](https://g.yuque.com/gr/latex?%3D*value*%5D%5D(http%3A%2F%2Fwww.runoob.com%2Fcssref%2Fsel-attr-end.html)%20%7C%20a%5Bsrc#card=math&code=%3D%2Avalue%2A%5D%5D%28http%3A%2F%2Fwww.runoob.com%2Fcssref%2Fsel-attr-end.html%29%20%7C%20a%5Bsrc)=".pdf"] | 选择每一个src属性的值以".pdf"结尾的元素 | 3 |

| [[*attribute**=**value*]](http://www.runoob.com/cssref/sel-attr-contain.html) | a[src*="44lan"] | 选择每一个src属性的值包含子字符串"44lan"的元素 | 3 |

| [:first-of-type](http://www.runoob.com/cssref/sel-first-of-type.html) | p:first-of-type | 选择每个p元素是其父级的第一个p元素 | 3 |

| [:last-of-type](http://www.runoob.com/cssref/sel-last-of-type.html) | p:last-of-type | 选择每个p元素是其父级的最后一个p元素 | 3 |

| [:only-of-type](http://www.runoob.com/cssref/sel-only-of-type.html) | p:only-of-type | 选择每个p元素是其父级的唯一p元素 | 3 |

| [:only-child](http://www.runoob.com/cssref/sel-only-child.html) | p:only-child | 选择每个p元素是其父级的唯一子元素 | 3 |

| [:nth-child(_n_)](http://www.runoob.com/cssref/sel-nth-child.html) | p:nth-child(2) | 选择每个p元素是其父级的第二个子元素 | 3 |

| [:nth-last-child(_n_)](http://www.runoob.com/cssref/sel-nth-last-child.html) | p:nth-last-child(2) | 选择每个p元素的是其父级的第二个子元素，从最后一个子项计数 | 3 |

| [:nth-of-type(_n_)](http://www.runoob.com/cssref/sel-nth-of-type.html) | p:nth-of-type(2) | 选择每个p元素是其父级的第二个p元素 | 3 |

| [:nth-last-of-type(_n_)](http://www.runoob.com/cssref/sel-nth-last-of-type.html) | p:nth-last-of-type(2) | 选择每个p元素的是其父级的第二个p元素，从最后一个子项计数 | 3 |

| [:last-child](http://www.runoob.com/cssref/sel-last-child.html) | p:last-child | 选择每个p元素是其父级的最后一个子级。 | 3 |

| [:root](http://www.runoob.com/cssref/sel-root.html) | :root | 选择文档的根元素 | 3 |

| [:empty](http://www.runoob.com/cssref/sel-empty.html) | p:empty | 选择每个没有任何子级的p元素（包括文本节点） | 3 |

| [:target](http://www.runoob.com/cssref/sel-target.html) | #news:target | 选择当前活动的#news元素（包含该锚名称的点击的URL） | 3 |

| [:enabled](http://www.runoob.com/cssref/sel-enabled.html) | input:enabled | 选择每一个已启用的输入元素 | 3 |

| [:disabled](http://www.runoob.com/cssref/sel-disabled.html) | input:disabled | 选择每一个禁用的输入元素 | 3 |

| [:checked](http://www.runoob.com/cssref/sel-checked.html) | input:checked | 选择每个选中的输入元素 | 3 |

| [:not(_selector_)](http://www.runoob.com/cssref/sel-not.html) | :not(p) | 选择每个并非p元素的元素 | 3 |

| [::selection](http://www.runoob.com/cssref/sel-selection.html) | ::selection | 匹配元素中被用户选中或处于高亮状态的部分 | 3 |

| [:out-of-range](http://www.runoob.com/cssref/sel-out-of-range.html) | :out-of-range | 匹配值在指定区间之外的input元素 | 3 |

| [:in-range](http://www.runoob.com/cssref/sel-in-range.html) | :in-range | 匹配值在指定区间之内的input元素 | 3 |

| [:read-write](http://www.runoob.com/cssref/sel-read-write.html) | :read-write | 用于匹配可读及可写的元素 | 3 |

| [:read-only](http://www.runoob.com/cssref/sel-read-only.html) | :read-only | 用于匹配设置 "readonly"（只读） 属性的元素 | 3 |

| [:optional](http://www.runoob.com/cssref/sel-optional.html) | :optional | 用于匹配可选的输入元素 | 3 |

| [:required](http://www.runoob.com/cssref/sel-required.html) | :required | 用于匹配设置了 "required" 属性的元素 | 3 |

| [:valid](http://www.runoob.com/cssref/sel-valid.html) | :valid | 用于匹配输入值为合法的元素 | 3 |

| [:invalid](http://www.runoob.com/cssref/sel-invalid.html) | :invalid | 用于匹配输入值为非法的元素 |
###### **CSS3** **边框（Borders）**
用CSS3，你可以创建圆角边框，添加阴影框，并作为边界的形象而不使用设计程序
| 属性 | 说明 | CSS |

| [border-image](http://www.runoob.com/cssref/css3-pr-border-image.html) | 设置所有边框图像的速记属性。 | 3 |

| [border-radius](http://www.runoob.com/cssref/css3-pr-border-radius.html) | 一个用于设置所有四个边框- *-半径属性的速记属性 | 3 |

| [box-shadow](http://www.runoob.com/cssref/css3-pr-box-shadow.html) | 附加一个或多个下拉框的阴影 | 3 |
###### 
###### **CSS3** **背景**
CSS3中包含几个新的背景属性，提供更大背景元素控制。
| 顺序 | 描述 | CSS |

| [background-clip](http://www.runoob.com/cssref/css3-pr-background-clip.html) | 规定背景的绘制区域。 | 3 |

| [background-origin](http://www.runoob.com/cssref/css3-pr-background-origin.html) | 规定背景图片的定位区域。 | 3 |

| [background-size](http://www.runoob.com/cssref/css3-pr-background-size.html) | 规定背景图片的尺寸。 | 3 |
###### 
多背景
###### **CSS3** **渐变**
CSS3 定义了两种类型的渐变（gradients）：

- 
**线性渐变（Linear Gradients）- 向下/向上/向左/向右/对角方向**

- 
**

   1. `background: linear-gradient(direction, color-stop1, color-stop2, ...);`

- 
**径向渐变（Radial Gradients）- 由它们的中心定义**

- 
**

   1. `background: radial-gradient(center, shape size, start-color, ..., last-color);`

###### **CSS3** **文本效果**
| 属性 | 描述 | CSS |

| [hanging-punctuation](http://www.runoob.com/cssref/css3-pr-hanging-punctuation.html) | 规定标点字符是否位于线框之外。 | 3 |

| [punctuation-trim](http://www.runoob.com/cssref/css3-pr-punctuation-trim.html) | 规定是否对标点字符进行修剪。 | 3 |

| text-align-last | 设置如何对齐最后一行或紧挨着强制换行符之前的行。 | 3 |

| text-emphasis | 向元素的文本应用重点标记以及重点标记的前景色。 | 3 |

| [text-justify](http://www.runoob.com/cssref/css3-pr-text-justify.html) | 规定当 text-align 设置为 "justify" 时所使用的对齐方法。 | 3 |

| [text-outline](http://www.runoob.com/cssref/css3-pr-text-outline.html) | 规定文本的轮廓。 | 3 |

| [text-overflow](http://www.runoob.com/cssref/css3-pr-text-overflow.html) | 规定当文本溢出包含元素时发生的事情。 | 3 |

| [text-shadow](http://www.runoob.com/cssref/css3-pr-text-shadow.html) | 向文本添加阴影。 | 3 |

| [text-wrap](http://www.runoob.com/cssref/css3-pr-text-wrap.html) | 规定文本的换行规则。 | 3 |

| [word-break](http://www.runoob.com/cssref/css3-pr-word-break.html) | 规定非中日韩文本的换行规则。 | 3 |

| [word-wrap](http://www.runoob.com/cssref/css3-pr-word-wrap.html) | 允许对长的不可分割的单词进行分割并换行到下一行。 | 3 |
###### **CSS3** **字体**
以前CSS3的版本，网页设计师不得不使用用户计算机上已经安装的字体。使用CSS3，网页设计师可以使用他/她喜欢的任何字体。当你发现您要使用的字体文件时，只需简单的将字体文件包含在网站中，它会自动下载给需要的用户。您所选择的字体在新的CSS3版本有关于@font-face规则描述。您"自己的"的字体是在 CSS3 [@font-face ](/font-face ) 规则中定义的。 
2D新转换属性
以下列出了所有的转换属性:
| Property | 描述 | CSS |

| [transform](http://www.runoob.com/cssref/css3-pr-transform.html) | 适用于2D或3D转换的元素 | 3 |

| [transform-origin](http://www.runoob.com/cssref/css3-pr-transform-origin.html) | 允许您更改转化元素位置 |
2D 转换方法
| 函数 | 描述 |

| matrix(_n_,_n_,_n_,_n_,_n_,_n_) | 定义 2D 转换，使用六个值的矩阵。 |

| translate(_x_,_y_) | 定义 2D 转换，沿着 X 和 Y 轴移动元素。 |

| translateX(_n_) | 定义 2D 转换，沿着 X 轴移动元素。 |

| translateY(_n_) | 定义 2D 转换，沿着 Y 轴移动元素。 |

| scale(_x_,_y_) | 定义 2D 缩放转换，改变元素的宽度和高度。 |

| scaleX(_n_) | 定义 2D 缩放转换，改变元素的宽度。 |

| scaleY(_n_) | 定义 2D 缩放转换，改变元素的高度。 |

| rotate(_angle_) | 定义 2D 旋转，在参数中规定角度。 |

| skew(_x-angle_,_y-angle_) | 定义 2D 倾斜转换，沿着 X 和 Y 轴。 |

| skewX(_angle_) | 定义 2D 倾斜转换，沿着 X 轴。 |

| skewY(_angle_) | 定义 2D 倾斜转换，沿着 Y 轴。 |
3D转换属性
下表列出了所有的转换属性：
| 属性 | 描述 | CSS |

| [transform](http://www.runoob.com/cssref/css3-pr-transform.html) | 向元素应用 2D 或 3D 转换。 | 3 |

| [transform-origin](http://www.runoob.com/cssref/css3-pr-transform-origin.html) | 允许你改变被转换元素的位置。 | 3 |

| [transform-style](http://www.runoob.com/cssref/css3-pr-transform-style.html) | 规定被嵌套元素如何在 3D 空间中显示。 | 3 |

| [perspective](http://www.runoob.com/cssref/css3-pr-perspective.html) | 规定 3D 元素的透视效果。 | 3 |

| [perspective-origin](http://www.runoob.com/cssref/css3-pr-perspective-origin.html) | 规定 3D 元素的底部位置。 | 3 |

| [backface-visibility](http://www.runoob.com/cssref/css3-pr-backface-visibility.html) | 定义元素在不面对屏幕时是否可见。 | 3 |
3D 转换方法
| 函数 | 描述 |

| matrix3d(_n_,_n_,_n_,_n_,_n_,_n_,

_n_,_n_,_n_,_n_,_n_,_n_,_n_,_n_,_n_,_n_) | 定义 3D 转换，使用 16 个值的 4x4 矩阵。 |

| translate3d(_x_,_y_,_z_) | 定义 3D 转化。 |

| translateX(_x_) | 定义 3D 转化，仅使用用于 X 轴的值。 |

| translateY(_y_) | 定义 3D 转化，仅使用用于 Y 轴的值。 |

| translateZ(_z_) | 定义 3D 转化，仅使用用于 Z 轴的值。 |

| scale3d(_x_,_y_,_z_) | 定义 3D 缩放转换。 |

| scaleX(_x_) | 定义 3D 缩放转换，通过给定一个 X 轴的值。 |

| scaleY(_y_) | 定义 3D 缩放转换，通过给定一个 Y 轴的值。 |

| scaleZ(_z_) | 定义 3D 缩放转换，通过给定一个 Z 轴的值。 |

| rotate3d(_x_,_y_,_z_,_angle_) | 定义 3D 旋转。 |

| rotateX(_angle_) | 定义沿 X 轴的 3D 旋转。 |

| rotateY(_angle_) | 定义沿 Y 轴的 3D 旋转。 |

| rotateZ(_angle_) | 定义沿 Z 轴的 3D 旋转。 |

| perspective(_n_) | 定义 3D 转换元素的透视视图。 |
###### **CSS3** **过渡**
过渡属性
下表列出了所有的过渡属性:
| 属性 | 描述 | CSS |

| [transition](http://www.runoob.com/cssref/css3-pr-transition.html) | 简写属性，用于在一个属性中设置四个过渡属性。 | 3 |

| [transition-property](http://www.runoob.com/cssref/css3-pr-transition-property.html) | 规定应用过渡的 CSS 属性的名称。 | 3 |

| [transition-duration](http://www.runoob.com/cssref/css3-pr-transition-duration.html) | 定义过渡效果花费的时间。默认是 0。 | 3 |

| [transition-timing-function](http://www.runoob.com/cssref/css3-pr-transition-timing-function.html) | 规定过渡效果的时间曲线。默认是 "ease"。 | 3 |

| [transition-delay](http://www.runoob.com/cssref/css3-pr-transition-delay.html) | 规定过渡效果何时开始。默认是 0。 | 3 |
###### **CSS3** **动画**
要创建CSS3动画，你需要了解@keyframes规则。@keyframes规则是创建动画。 @keyframes规则内指定一个CSS样式和动画将逐步从目前的样式更改为新的样式。
## 实例
当动画为 25% 及 50% 时改变背景色，然后当动画 100% 完成时再次改变：
下面的表格列出了 [@keyframes ](/keyframes ) 规则和所有动画属性： 
| 属性 | 描述 | CSS |

| [@keyframes](http://www.runoob.com/cssref/css3-pr-animation-keyframes.html) | 规定动画。 | 3 |

| [animation](http://www.runoob.com/cssref/css3-pr-animation.html) | 所有动画属性的简写属性，除了 animation-play-state 属性。 | 3 |

| [animation-name](http://www.runoob.com/cssref/css3-pr-animation-name.html) | 规定 [@keyframes ](/keyframes ) 动画的名称。 | 3 | 

| [animation-duration](http://www.runoob.com/cssref/css3-pr-animation-duration.html) | 规定动画完成一个周期所花费的秒或毫秒。默认是 0。 | 3 |

| [animation-timing-function](http://www.runoob.com/cssref/css3-pr-animation-timing-function.html) | 规定动画的速度曲线。默认是 "ease"。 | 3 |

| [animation-delay](http://www.runoob.com/cssref/css3-pr-animation-delay.html) | 规定动画何时开始。默认是 0。 | 3 |

| [animation-iteration-count](http://www.runoob.com/cssref/css3-pr-animation-iteration-count.html) | 规定动画被播放的次数。默认是 1。 | 3 |

| [animation-direction](http://www.runoob.com/cssref/css3-pr-animation-direction.html) | 规定动画是否在下一周期逆向地播放。默认是 "normal"。 | 3 |

| [animation-play-state](http://www.runoob.com/cssref/css3-pr-animation-play-state.html) | 规定动画是否正在运行或暂停。默认是 "running"。 | 3 |
###### **CSS3** **多列**
下表列出了所有 CSS3 的多列属性：
| 属性 | 描述 |

| [column-count](http://www.runoob.com/cssref/css3-pr-column-count.html) | 指定元素应该被分割的列数。 |

| [column-fill](http://www.runoob.com/cssref/css3-pr-column-fill.html) | 指定如何填充列 |

| [column-gap](http://www.runoob.com/cssref/css3-pr-column-gap.html) | 指定列与列之间的间隙 |

| [column-rule](http://www.runoob.com/cssref/css3-pr-column-rule.html) | 所有 column-rule-* 属性的简写 |

| [column-rule-color](http://www.runoob.com/cssref/css3-pr-column-rule-color.html) | 指定两列间边框的颜色 |

| [column-rule-style](http://www.runoob.com/cssref/css3-pr-column-rule-style.html) | 指定两列间边框的样式 |

| [column-rule-width](http://www.runoob.com/cssref/css3-pr-column-rule-width.html) | 指定两列间边框的厚度 |

| [column-span](http://www.runoob.com/cssref/css3-pr-column-span.html) | 指定元素要跨越多少列 |

| [column-width](http://www.runoob.com/cssref/css3-pr-column-width.html) | 指定列的宽度 |

| [columns](http://www.runoob.com/cssref/css3-pr-columns.html) | 设置 column-width 和 column-count 的简写 |
###### **CSS3** **盒模型**
在 CSS3 中, 增加了一些新的用户界面特性来调整元素尺寸，框尺寸和外边框，主要包括以下用户界面属性：

- resize：none | both | horizontal | vertical | inherit
- box-sizing: content-box | border-box | inherit
- outline:outline-color outline-style outline-width outine-offset

resize属性指定一个元素是否应该由用户去调整大小。
box-sizing 属性允许您以确切的方式定义适应某个区域的具体内容。
outline-offset 属性对轮廓进行偏移，并在超出边框边缘的位置绘制轮廓。
###### **CSS3伸缩布局盒模型(弹性盒)**
CSS3 弹性盒（ Flexible Box 或 flexbox），是一种当页面需要适应不同的屏幕大小以及设备类型时确保元素拥有恰当的行为的布局方式。
引入弹性盒布局模型的目的是提供一种更加有效的方式来对一个容器中的子元素进行排列、对齐和分配空白空间。
下表列出了在弹性盒子中常用到的属性:
| 属性 | 描述 |

| [display](http://www.runoob.com/cssref/pr-class-display.html) | 指定 HTML 元素盒子类型。 |

| [flex-direction](http://www.runoob.com/cssref/css3-pr-flex-direction.html) | 指定了弹性容器中子元素的排列方式 |

| [justify-content](http://www.runoob.com/cssref/css3-pr-justify-content.html) | 设置弹性盒子元素在主轴（横轴）方向上的对齐方式。 |

| [align-items](http://www.runoob.com/cssref/css3-pr-align-items.html) | 设置弹性盒子元素在侧轴（纵轴）方向上的对齐方式。 |

| [flex-wrap](http://www.runoob.com/cssref/css3-pr-flex-wrap.html) | 设置弹性盒子的子元素超出父容器时是否换行。 |

| [align-content](http://www.runoob.com/cssref/css3-pr-align-content.html) | 修改 flex-wrap 属性的行为，类似 align-items, 但不是设置子元素对齐，而是设置行对齐 |

| [flex-flow](http://www.runoob.com/cssref/css3-pr-flex-flow.html) | flex-direction 和 flex-wrap 的简写 |

| [order](http://www.runoob.com/cssref/css3-pr-order.html) | 设置弹性盒子的子元素排列顺序。 |

| [align-self](http://www.runoob.com/cssref/css3-pr-align-self.html) | 在弹性子元素上使用。覆盖容器的 align-items 属性。 |

| [flex](http://www.runoob.com/cssref/css3-pr-flex.html) | 设置弹性盒子的子元素如何分配空间。 |
###### **CSS3** **多媒体查询**
从 CSS 版本 2 开始，就可以通过媒体类型在 CSS 中获得媒体支持。如果您曾经使用过打印样式表，那么您可能已经使用过媒体类型。清单 1 展示了一个示例。
清单 1. 使用媒体类型
清单 2. 媒体查询规则

- 
`@media all` 是媒体类型，也就是说，将此 CSS 应用于所有媒体类型。

- 
`(min-width:800px)` 是包含媒体查询的表达式，如果浏览器的最小宽度为 800 像素，则会告诉浏览器只运用下列 CSS。


清单 3. `and` 条件
清单 4. `or` 关键词
清单 5. 使用 `not`

- 参考：[菜鸟教程](http://www.runoob.com/)




