# agruments应用一求出函数参数的总合

```
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>agruments应用一求出函数参数的总合</title>
</head>

<body>
    <pre style="
    color: green;
    padding: 10px 15px;
    background: #f0f0f0;
    border: 1px dotted #333;
    font: 12px/1.5 Courier New;
">
    function test (){
        let arr = Array.prototype.map.call(arguments,item=>item);
         alert(eval(arr.join('+')))
       }
       
       test(1,2,3,5,4,6,7,8,9,7,4,5,4);
 </pre>
</body>
<script>
    function test() {
        let arr = Array.prototype.map.call(arguments, item => item);
        alert(eval(arr.join('+')))
    }
    test(1, 2, 3, 5, 4, 6, 7, 8, 9, 7, 4, 5, 4);

</script>

</html>
```
