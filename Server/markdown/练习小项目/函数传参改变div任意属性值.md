# 函数传参改变div任意属性值

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>函数传参改变div任意属性的值</title>
</head>
<style>
  .container {
    width: 500px;
    height: 500px;
    position: relative;
    text-align: center;
    margin: 0 auto;
  }

  .container .test {
    width: 120px;
    height: 120px;
    border: red solid 1px;
    position: absolute;
    left: 50%;
    top: 50%;
    margin-left: -60px;
    margin-top: -60px;
    background: black;
  }
</style>

<body>
  <div class="container">
    <table>
      属性名:<input type="text" value="background"><br>
      属性值:<input type="text" value="red"><br>
      <button>确定</button> <button>重置</button>
    </table>
    <div class="test">

    </div>
  </div>
</body>
<script>
  let buttons: any = document.querySelectorAll('button')
  let test: any = document.querySelector('.test')
  let input: any = document.querySelectorAll('input')

  buttons[0].onclick = (event: any): void => {
    let arr = Array.prototype.map.call(input, item => item.value)
    test.style.setProperty(arr[0], arr[1]);
  }
  buttons[1].onclick = (event: any): void => {
    test.removeAttribute('style');
  }


</script>

</html>
```
