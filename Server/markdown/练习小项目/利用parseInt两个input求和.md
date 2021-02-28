# 利用parseInt两个input求和

```
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>parseint解析字母</title>
</head>

<body>
    <input type="text"> + <input type="text"> = ? <br>
    <button>求和</button>

</body>
<script>

    let inputs: any = document.querySelectorAll('input');
    let button: any = document.querySelector('button');
    button.onclick = () => {
        let one = inputs[0].value.match(/\d/g)
        let two = inputs[1].value.match(/\d/g)
        alert(parseInt(one.join('')) + parseInt(two.join('')))
    }

</script>

</html>
```
