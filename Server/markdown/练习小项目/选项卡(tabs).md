# 选项卡(tabs)

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>选项卡</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      list-style: none;
      color: white;
    }

    .action {
      background: wheat;
      color: red;
    }

    .container {
      width: 700px;
      height: 500px;
      margin: 0 auto;
      border: black solid 0.5px;
    }

    .tabs-wrapper {
      width: 100%;
      display: flex;
      text-align: center;
      height: 35px;
      line-height: 35px;
      background: black;
    }

    .tabs-wrapper>p {
      padding: 0 5px;
    }

    .tabs-content {
      width: 100%;
      display: flex;
      height: 465px;
    }

    .tabs-content>div {
      display: none;
      background: salmon;
      width: 100%;
      text-align: center;
    }

    .tabs-content>div:nth-child(1) {
      display: block;
    }

    .tabs-content li {
      margin: 5px 0;
    }
  </style>
</head>

<body>
  <section class="container">
    <div class="tabs-wrapper">
      <p class="action">第一课</p>
      <p>第二课</p>
      <p>第三课</p>
    </div>
    <div class="tabs-content">
      <div>
        <ul>
          <li>网页特效原理分析</li>
          <li>响应用户操作</li>
          <li>提示框效果</li>
          <li>事件驱动</li>
          <li>元素属性操作</li>
          <li>动手编写第一个JS特效</li>
          <li>引入函数</li>
          <li>网页换肤效果</li>
          <li>展开/收缩播放列表效果</li>

        </ul>
      </div>
      <div>
        <ul>
          <li>改变网页背景颜色</li>
          <li>函数传参</li>
          <li>高重用性函数的编写</li>
          <li>126邮箱全选效果</li>
          <li>循环及遍历操作</li>
          <li>调试器的简单使用</li>
          <li>典型循环的构成</li>
          <li>for循环配合if判断</li>
          <li>className的使用</li>
          <li>innerHTML的使用</li>
          <li>戛纳印象效果</li>
          <li>数组</li>
          <li>字符串连接</li>

        </ul>
      </div>
      <div>
        <ul>
          <li>JavaScript组成：ECMAScript、DOM、BOM，JavaScript兼容性来源</li>
          <li>JavaScript出现的位置、优缺点</li>
          <li>变量、类型、typeof、数据类型转换、变量作用域</li>
          <li>闭包：什么是闭包、简单应用、闭包缺点</li>
          <li>运算符：算术、赋值、关系、逻辑、其他运算符</li>
          <li>程序流程控制：判断、循环、跳出</li>
          <li>命名规范：命名规范及必要性、匈牙利命名法</li>
          <li>函数详解：函数构成、调用、事件、传参数、可变参、返回值</li>
          <li>定时器的使用：setInterval、setTimeout</li>
          <li>定时器应用：站长站导航效果</li>
          <li>定时器应用：自动播放的选项卡</li>
          <li>定时器应用：数码时钟</li>
          <li>程序调试方法</li>
        </ul>
      </div>
    </div>
  </section>
</body>
<script>
  let p: any = document.querySelectorAll('.tabs-wrapper p');
  let div: any = document.querySelectorAll('.tabs-content div')
  p.forEach((item, index) => {
    item.onmouseover = () => {
      p.forEach(item => item.removeAttribute('class'));
      item.setAttribute('class', 'action');
      div.forEach(item => item.style.display = 'none');
      div[index].style.display = 'block';
    }
  })


  console.log(div[0].parentNode.children)
  console.log(div)//NodeList   

//HTMLCollection 和nodelist是两个东西别搞混了

</script>

</html>
```
