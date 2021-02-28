# 用户按下键盘，显示keyCode

```
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>用户按下键盘，显示keyCode</title>
</head>
<style>
    pre {
        color: green;
        padding: 10px 15px;
        background: #f0f0f0;
        border: 1px dotted #333;
        font: 12px/1.5 Courier New;
        text-align: left;
    }

    body {
        text-align: center;
    }
</style>

<body>
    
<pre>
        &lt;script type="text/javascript"&gt;
        window.onload = function ()
        {
            var oP = document.getElementsByTagName("p")[0];	
            document.onkeydown = function (event)
            {
                var event = event || window.event;
                oP.innerHTML = event.keyCode;
                return false
            }
        }
        &lt;/script&gt;
        </pre>
<h2></h2>
</body>
<script>
    let H2: HTMLDivElement = document.querySelector('H2');
    window.onkeydown = (event: KeyboardEvent) => {
        H2.innerHTML = `${event.keyCode}`
    }
</script>

</html>
```
