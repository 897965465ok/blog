# jquery选择器

```
   $('#one').css('background','#bbffaa');
    $('.mini').css('background','#bbffaa');
    $('div').css('background','#bbffaa');
    $('*').css('background','#bbffaa');
    $('span,$two').css('background','#bbffaa');
    $('.noe+div') //选取class 为one的下一个div同辈元素
    $('#two~div')//选取ID为two的元素后面的所有元素 
    $('.noe+div')//可用next代替
    $('#two~div')//可用nextAll代替
    //siblings()
    $('#div:not(class)');//不是class的元素
    $('#div:even');//偶数
    $('#div:odd')//奇数
    $('#div:eq(1)')//索引1
    $('#div:gt(3)')//大于
    $('#div:lt(3)')//小于
    $('nav:header')//选取h1~h6的标签
    $('div:anumated')//正在动画的元素
    $('div:focus')//焦点可用可用用focus()代替
    $('div:contains("is")')//查找含有is 的div
    $('div:empty')//选择空的div
    $('div:has(p)')//选择有p标签的div
    $('div:parent')//选择带有子元素的div
    $('div:hidden')//选择所有隐藏的div
    $('div:visible')//选择可见的div  
    $('div[id]')//选择带有id属性的div
    $('div[id=box]')//选择属性id等于box值的div
    $('div[id!=box]')//选择属性id不等于box值的div
    $('div[id^=box]')//选择属性id 以box值开头的div
    $('div[id$=box]')//选择属性id 以box值结尾的div
    $('div[id*=box]')//选择属性id 包含box值的div
    $('div[title|="en"]')//选择属性tetle为en或以en为前缀的div(tetle=en or tetle=en-ok )    
    $('div[title~="uk"]')//目前不知道等会查
    $('div[id][title$="test"]')//选择属性为id 和属性title属性为test结尾的div
    $('div:nht-child(1)')//选择div下面的第一个元素 nht()值有 even偶数 odd奇数 nth-child(3n)倍数 3 6 9
    $('ui li:first-child')//选择ul 下面的第一个il元素
    $('ul li :last-child')//选择ul 下面的最后一个il元素
    $('ul li :only-child')//选取唯一的li 元素如果含有其他元素这个将无效
    //***********************表单选择器*********************
    $('#form1:enabled')//选取id 为form1d 的表单内的所有可用元素
    $('#form2:disabled')//选择id 为form2的表单内所有不可用元素
    $('input:checked')//选取所有被选中的input元素
    $('input:selected')//在下啦框option中选中预定的
    $(":input")//选取所有<input> <textarea> <select> <button>
    $(':text')//选取所有单行文本框
    $(':password')//选取所有密码框
    $(':radio')//选取所有单选框
    $(':checkbox')//选取所有的复选框
    $(':submit')//选取所有提交按钮
    $(':image')//选取所有的图形按钮
    $(':reset')//选取所有重置按钮
    $(':button')//选取所有提交按钮
    $(':file')//选取所有上传域
```
