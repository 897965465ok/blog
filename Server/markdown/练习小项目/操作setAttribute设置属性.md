# 操作setAttribute设置属性

```
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>切换背景</title>
</head>
<style>
nav,ul,li{
    list-style: none;
    display: inline-block;
    width: 50px;
    height: 50px;
    float: left;
}

.container{
    width:250px;
    height:250px;
    background:green;
}
.container-one{
    width:250px;
    height:250px;
    background:red;
}
.container-two{
    width:250px;
    height:250px;
    background:black;
}
</style>
<body>
    <header>
        <nav>
            <ul>
               <li><button>变色 1</button></li>
               <li><button>变色 2</button></li>
            </ul>
        </nav>
    </header>
    <section class="container">
     <img src="123OOXX" alt="">
    </section>
    <footer>

    </footer>
</body>
<script>
let but: any = document.querySelectorAll('header nav ul li button');
let div: any = document.querySelector('.container ')
let arr:Array<string> = ['container-one','container-two'] 
but.forEach((item,index) => {
    item.addEventListener('click', event => {
        div.setAttribute('class',arr[index])
    }, false)
});

</script>
</html>
```
