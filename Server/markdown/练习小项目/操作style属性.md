# 操作style属性

```
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <style>
        .body {
            height: 200px;
            width: 200px;
            background: rgb(223, 207, 207);
            margin: 0 auto;
            
        }
        .body ul li{
          list-style-type: none;
        }
    </style>
</head>

<body>
    <header>
        <nav>

        </nav>
    </header>
    <section class="container">
        <div class="body">
            <ul>
                <li><button>变宽</button></li>
                <li><button>变高</button></li>
                <li><button>变色</button></li>
                <li><button>隐藏</button> </li>
                <li> <button>重置</button> </li>
            </ul>
        </div>
    </section>
    <footer>

    </footer>
</body>
<script>
let div: any = document.querySelector('.body');
let button: any = div.querySelectorAll('ul li button');
let arr: Array<string> = [
    'width:500px;',
    'height:500px;',
    'background:red;',
    'display:none;',
    'width:200px;,background: rgb(223, 207, 207);,height:200px;'
];
button.forEach((item, index) => {
    item.addEventListener('click', (event: any): void => {
        console.log(event.target) //获取到点击的对象
        div.style.cssText = arr[index]
    }, false)
})

// js改变css样式的三种方法
// div.cssText='width:200px;,background: rgb(223, 207, 207);,height:200px;'
// div.style.setProperty('border','1px solid blue');
// div.style.width = "800px";
// getComputedStyle(element,null).getPropertyValue('height')//获取非行间样式属性名和属性值 ，然后再获取指定属性值

</script>
</html>
```
