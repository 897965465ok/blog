# 点击显示innerHTML

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>点击获取内部内容</title>
</head>
<style>
  .container {
    border: pink solid 1px;
    text-align: center;
    width: 500px;
    margin: 0 auto;
  }
</style>

<body>
  <div class="container">
    新华网北京5月29日电(新华社记者)从三聚氰胺到瘦肉精，
    从染色馒头到毒豆芽，
    在食品中添加非食用物质和滥用食品添加剂已成为百姓餐桌上突出的"不安全因素"。
    近期以来，从北京到广东，从浙江到宁夏，一场打击食品非法添加的"风暴"席卷全国。
  </div>
</body>
<script>
  let container: any = document.querySelector('.container');

  container.onclick = () => {
    alert(container.innerHTML)
  }

</script>

</html>
```
