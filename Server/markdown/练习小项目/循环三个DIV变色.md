# 循环三个DIV变色

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>循环三个DIV将其变色</title>
  <style>
    div {
      width: 300px;
      height: 300px;
      float: left;
      border: red solid 0.5px;
      margin: 1px;

    }
  </style>

</head>

<body>
  <div></div>
  <div></div>
  <div></div>
  <button>点击变色</button>
</body>
<script>
let div: any = document.querySelectorAll('div');
let button: any = document.querySelector('button');
button.onclick = (event) => {

    div.forEach(item => {
    item.style = 'background:blue;'
    })

}



</script>

</html>
```
